package filedetails

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	json "github.com/json-iterator/go"
)

// Structure describing the parameters of a request to IPublishedFileService/GetDetails/v1/
type Query struct {
	key string // Access API key

	Language                  string   `json:"language,omitempty"`                  // Specifies the localized text to return. Defaults to English. //* ELanguage
	DesiredRevision           string   `json:"desired_revision,omitempty"`          // Return the data for the specified revision. //* EPublishedFileRevision
	PublishedFileIDs          []uint64 `json:"publishedfileids"`                    // Set of published file Ids to retrieve details for.
	AppID                     uint64   `json:"appid,omitempty"`                     // Application ID
	ReturnPlaytimeStats       uint32   `json:"return_playtime_stats,omitempty"`     // Return playtime stats for the specified number of days before today.
	IncludeTags               bool     `json:"includetags,omitempty"`               // If true, return tag information in the returned details.
	IncludeAdditionalPreviews bool     `json:"includeadditionalpreviews,omitempty"` // If true, return preview information in the returned details.
	IncludeChildren           bool     `json:"includechildren,omitempty"`           // If true, return children in the returned details.
	IncludeKVTags             bool     `json:"includekvtags,omitempty"`             // If true, return key value tags in the returned details.
	IncludeVotes              bool     `json:"includevotes,omitempty"`              // If true, return vote data in the returned details.
	ShortDescription          bool     `json:"short_description,omitempty"`         // If true, return a short description instead of the full description.
	IncludeForSaleData        bool     `json:"includeforsaledata,omitempty"`        // If true, return pricing data, if applicable.
	IncludeMetadata           bool     `json:"includemetadata,omitempty"`           // If true, populate the metadata field.
	StripDescriptionBBCode    bool     `json:"strip_description_bbcode,omitempty"`  // Strips BBCode from descriptions.
	IncludeReactions          bool     `json:"includereactions,omitempty"`          // If true, then reactions to items will be returned.
	AdminQuery                bool     `json:"admin_query,omitempty"`               // Admin tool is doing a query, return hidden items
}

/*
New creates a new Query instance with the required fields.
It initializes the Query struct with the provided file IDs and API key.
It also sets default values for ShortDescription, StripDescriptionBBCode, and IncludeKVTags.

Parameters:
  - fileIDs: A slice of published file IDs to retrieve details for.
  - key: The API key for accessing the Steam API.

Returns:
  - A pointer to a Query instance if key is non-empty and fileIDs is not empty.
  - nil otherwise.
*/
func New(fileIDs []uint64, key string) *Query {
	if key == "" || len(fileIDs) == 0 {
		return nil
	}

	return &Query{
		key:                    key,
		PublishedFileIDs:       fileIDs,
		ShortDescription:       true,
		StripDescriptionBBCode: true,
		IncludeKVTags:          true,
	}
}

/*
SetKey sets the API key for the GetDetails request.

Parameters:
  - key: The API key string.
*/
func (q *Query) SetKey(key string) {
	q.key = key
}

/*
SetFileIDs sets the list of published file IDs for the GetDetails request.

Parameters:
  - ids: A slice of uint64 representing the published file IDs.
*/
func (q *Query) SetFileIDs(ids []uint64) {
	q.PublishedFileIDs = ids
}

/*
SetAppID sets the Application ID for the GetDetails request.

Parameters:
  - id: The application ID as a uint64.
*/
func (q *Query) SetAppID(id uint64) {
	q.AppID = id
}

/*
Get sends the GetDetails request to the Steam API and retrieves file details.

It constructs an HTTP request to the Steam API's GetDetails endpoint using the parameters
defined in the Query struct. The method handles authentication using the provided API key,
sends the request, parses the JSON response, and returns a slice of FileDetail or an error if the
request fails.

Returns:
  - A slice of FileDetail containing the details of the requested files.
  - An error if the request or parsing fails.

Example:

	params := modinfo.New([]uint64{123456, 789012}, "your_api_key")
	details, err := params.Get()
	if err != nil {
		// handle error
	}
	// use details
*/
func (q *Query) Get() ([]FileDetail, error) {
	if q == nil {
		return nil, fmt.Errorf("Query request parameters not set")
	}
	if len(q.key) != 32 {
		return nil, fmt.Errorf("Steam API key is empty or does not match")
	}

	// build URL with query
	query := url.Values{}
	query.Set("key", q.key)
	for i, id := range q.PublishedFileIDs {
		query.Set("publishedfileids["+strconv.Itoa(i)+"]", strconv.FormatUint(id, 10))
	}

	v := reflect.ValueOf(*q)
	t := reflect.TypeOf(*q)

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" || tag == "publishedfileids" {
			continue
		}
		key := strings.Split(tag, ",")[0]

		switch value.Kind() {
		case reflect.String:
			if value.String() != "" {
				query.Set(key, value.String())
			}
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			if value.Uint() != 0 {
				query.Set(key, strconv.FormatUint(value.Uint(), 10))
			}
		case reflect.Bool:
			if value.Bool() {
				query.Set(key, "true")
			}
		}
	}
	url := baseURL + "?" + query.Encode()

	// get and decode response
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error close response body: %v", err)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	var result struct {
		Response struct {
			Details []FileDetail `json:"publishedfiledetails"`
		} `json:"response"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// Set file details URL
	for i, f := range result.Response.Details {
		if f.URL == "" {
			result.Response.Details[i].URL = fmt.Sprintf("%s%d", baseFileURL, f.PublishedFileID)
		}
	}

	// Return struct if AppID not set
	if q.AppID == 0 {
		return result.Response.Details, nil
	}

	// Validate AppID == ConsumerAppid
	for _, f := range result.Response.Details {
		if q.AppID != f.ConsumerAppID {
			return result.Response.Details, fmt.Errorf(
				"not match in response CreatorAppid %d and ConsumerAppid %d for item %d",
				f.CreatorAppID, f.ConsumerAppID, f.PublishedFileID,
			)
		}
	}

	return result.Response.Details, nil
}
