/*
Package filedetails provides structures and methods for interacting with the Steam Workshop's Published File Service API.
It allows retrieval of detailed information about files (mods, guides, images, etc.)
published on the Steam Workshop, including metadata, tags, votes, playtime statistics, and more.

The package includes types for representing file details, custom types for handling specific data formats,
and methods for constructing and sending API requests to the Steam API's GetDetails endpoint.

API Documentation:
  - [Steam API GetDetails v1]
  - [Steam API Reference by XPaw]

To obtain an API key, visit: [Steam Dev API Key]

# Example usage:

	package main

	import (
		"fmt"
		"log"

		"github.com/woozymasta/steam/filedetails"
	)

	func main() {
		// Retrieve the Steam API key from environment variables.
		key, exists := os.LookupEnv("STEAM_API_KEY")
		if !exists {
			log.Fatal("STEAM_API_KEY environment variable is not set")
		}

		// Steam files ID
		fileIDs := []uint64{123456, 789012}

		// Create a new SteamQuery instance with the API key.
		query := filedetails.New(fileIDs, apiKey)
		if query == nil {
			log.Fatal("Invalid API key or file IDs")
		}

		// Execute the query.
		files, err := query.Get()
		if err != nil {
			log.Fatalf("Failed to get file files: %v", err)
		}

		// Output the retrieved files info.
		for _, f := range files {
			fmt.Printf("Title: %s, Description: %s\n", f.Title, f.FileDescription)
		}
	}

[Steam API GetDetails v1]: https://api.steampowered.com/IPublishedFileService/GetDetails/v1/
[Steam API Reference by XPaw]: https://steamapi.xpaw.me/#IPublishedFileService/GetDetails
[Steam Dev API Key]: https://steamcommunity.com/dev/apikey
*/
package filedetails

const (
	baseURL         string = "https://api.steampowered.com/IPublishedFileService/GetDetails/v1/"
	baseFileURL     string = "https://steamcommunity.com/sharedfiles/filedetails/?id="
	defaultChunkMax        = 220
	defaultConns           = 10
)
