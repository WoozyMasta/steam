package filedetails

import (
	"fmt"
	"os"
	"testing"
)

func TestGetMods(t *testing.T) {
	key, ok := os.LookupEnv("STEAM_API_KEY")
	if !ok {
		t.Error("Steam API key must be pass in variable 'STEAM_API_KEY'")
	}

	ids := []uint64{
		1559212036, // CF
		2545327648, // Dabs Framework
		2874283306, // DayZ launcher Linux
		3213284654, // DayZ (screenshots)
	}

	query := New(ids, key)
	query.SetAppID(221100)

	files, err := query.Get()
	if err != nil {
		t.Errorf("Cant get files %v", err)
	}

	for i, f := range files {
		fmt.Printf("ID: %d, URL: %s, Title: %s, Updated: %s\n", f.PublishedFileID, f.URL, f.Title, f.TimeUpdated)
		if f.PublishedFileID != ids[i] {
			t.Errorf("Return mismatched ID %d, expected: %d", f.PublishedFileID, ids[i])
		}
	}
}
