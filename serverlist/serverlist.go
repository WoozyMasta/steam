/*
Package serverlist provides functionality to interact with the Steam Game Servers API.
It allows retrieval and filtering of game servers based on various criteria.

The package uses an undocumented API endpoint [Steam API Reference by XPaw]
and filters based on the deprecated [Master Server Query Protocol].

To obtain an API key, visit: [Steam Dev API Key]

# Example usage:

	package main

	import (
		"fmt"
		"log"
		"os"

		"github.com/woozymasta/steam/serverlist"
	)

	func main() {
		// Retrieve the Steam API key from environment variables.
		key, exists := os.LookupEnv("STEAM_API_KEY")
		if !exists {
			log.Fatal("STEAM_API_KEY environment variable is not set")
		}

		// Initialize the filter with desired criteria.
		filter := &serverlist.Filter{}
		filter.Add(serverlist.KeyAppID, "221100")             // Filter for DayZ servers.
		filter.Add(serverlist.KeyGameType, "battleye")        // Include servers with BattleEye.
		filter.AddNor(serverlist.KeyGameType, "external")     // Exclude servers with "external" game type.

		// Create a new SteamQuery instance with the API key.
		query := serverlist.New(key)

		// Optionally, set a custom limit for the number of servers to retrieve.
		query.SetLimit(300)

		// Execute the query with the specified filter.
		servers, err := query.Get(filter)
		if err != nil {
			log.Fatalf("Error retrieving servers: %v", err)
		}

		// Output the retrieved servers.
		for _, server := range servers {
			fmt.Printf("Server Name: %s, Address: %s, Players: %d/%d\n",
				server.Name, server.Addr, server.Players, server.MaxPlayers)
		}
	}

[Steam API Reference by XPaw]: https://steamapi.xpaw.me/#IGameServersService/GetServerList
[Master Server Query Protocol]: https://developer.valvesoftware.com/wiki/Master_Server_Query_Protocol
[Steam Dev API Key]: https://steamcommunity.com/dev/apikey
*/
package serverlist

const (
	// DefaultLimit defines the default maximum number of servers to retrieve.
	DefaultLimit = 10000

	// baseURL is the endpoint for the Steam Game Servers API.
	baseURL = "https://api.steampowered.com/IGameServersService/GetServerList/v1/"
)
