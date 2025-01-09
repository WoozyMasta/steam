# serverlist

`serverlist` is a Go package that provides functionality to interact with
the Steam Game Servers API. It allows developers to retrieve and filter game
servers based on various criteria, making it easier to integrate Steam
server data into your applications.

The package uses an undocumented API endpoint [Steam API Reference by XPaw][]
and filters based on the deprecated [Master Server Query Protocol][].

To obtain an API key, visit: [Steam Dev API Key][].

## Features

* **Retrieve Server Lists:**
  Fetch a list of game servers using Steam's Game Servers API.
* **Advanced Filtering:**
  Apply complex filters to narrow down servers based on game type, map,
  player count, and more.
* **Customizable Limits:**
  Specify the number of servers to retrieve per request.
* **Easy Integration:**
  Simple API design for seamless integration into Go projects.

## Installation

Install the package using `go get`:

```bash
go get github.com/woozymasta/steam
```

## Usage

Below is an example of how to use the serverlist package to retrieve a list
of official DayZ servers (AppID: `221100`) using the provided
API key and filters.

```go
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
```

## Support me 💖

If you enjoy my projects and want to support further development,
feel free to donate! Every contribution helps to keep the work going.
Thank you!

<!-- omit in toc -->
### Crypto Donations

<!-- cSpell:disable -->
* **BTC**: `1Jb6vZAMVLQ9wwkyZfx2XgL5cjPfJ8UU3c`
* **USDT (TRC20)**: `TN99xawQTZKraRyvPAwMT4UfoS57hdH8Kz`
* **TON**: `UQBB5D7cL5EW3rHM_44rur9RDMz_fvg222R4dFiCAzBO_ptH`
<!-- cSpell:enable -->

Your support is greatly appreciated!

<!-- Links-->

[Steam API Reference by XPaw]: https://steamapi.xpaw.me/#IGameServersService/GetServerList
[Master Server Query Protocol]: https://developer.valvesoftware.com/wiki/Master_Server_Query_Protocol
[Steam Dev API Key]: https://steamcommunity.com/dev/apikey
