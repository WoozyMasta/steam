package filedetails

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"
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

func TestGetManyFiles(t *testing.T) {
	key, ok := os.LookupEnv("STEAM_API_KEY")
	if !ok {
		t.Error("Steam API key must be pass in variable 'STEAM_API_KEY'")
	}

	count := 700
	ids := randomIDs(count, 1500000000, 3500000000)
	query := New(ids, key)

	files, err := query.Get()
	if err != nil {
		t.Errorf("Cant get files %v", err)
	}

	if len(files) < count {
		t.Errorf("Return %d files, but expected %d", len(files), count)
	}
}

func TestGetManyFilesConcurrent(t *testing.T) {
	key, ok := os.LookupEnv("STEAM_API_KEY")
	if !ok {
		t.Error("Steam API key must be pass in variable 'STEAM_API_KEY'")
	}

	count := 5000
	ids := randomIDs(count, 1500000000, 3500000000)
	query := New(ids, key)
	query.SetConcurrency(100)

	files, err := query.GetConcurrent()
	if err != nil {
		t.Errorf("Cant get files %v", err)
	}

	if len(files) < count {
		t.Errorf("Return %d files, but expected %d", len(files), count)
	}
}

func randomIDs(count int, min, max uint64) []uint64 {
	if min > max {
		return []uint64{min, max}
	}

	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	numbers := make([]uint64, count)

	for i := 0; i < count; i++ {
		numbers[i] = rnd.Uint64()%(max-min+1) + min
	}

	return numbers
}
