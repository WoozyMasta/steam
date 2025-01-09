package serverlist

import (
	"fmt"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
)

// SteamQuery provides an interface for interacting with the Steam API.
type SteamQuery struct {
	client *http.Client
	key    string
	limit  int
}

// New creates a new instance of SteamQuery with the provided API key.
// It initializes the HTTP client and sets the default limit for server retrieval.
func New(apiKey string) *SteamQuery {
	return &SteamQuery{
		key:    apiKey,
		client: &http.Client{},
		limit:  DefaultLimit,
	}
}

// SetKey sets the API key for the SteamQuery instance.
// This allows changing the API key after initialization if needed.
func (sq *SteamQuery) SetKey(key string) {
	sq.key = key
}

// SetLimit sets the maximum number of servers to retrieve in a single API request.
// This overrides the default limit defined by DefaultLimit.
func (sq *SteamQuery) SetLimit(limit int) {
	sq.limit = limit
}

// Get performs a request to the Steam API with the provided filter and returns a list of servers.
// It constructs the filter string, sends the HTTP GET request, and decodes the JSON response.
// Returns an error if the request fails, the response status is not OK, or the response cannot be decoded.
func (sq *SteamQuery) Get(filter *Filter) (Servers, error) {
	filterString, err := filter.String()
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("key", sq.key)
	params.Set("filter", filterString)
	params.Set("format", "json")
	params.Set("limit", fmt.Sprintf("%d", sq.limit))

	resp, err := sq.client.Get(baseURL + "?" + params.Encode())
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	var result struct {
		Response struct {
			Servers Servers `json:"servers"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Response.Servers, nil
}
