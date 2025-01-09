# filedetails

`filedetails` is a Go package that provides structures and methods for
interacting with the Steam Workshop's Published File Service API.  
It allows retrieval of detailed information about files
(mods, guides, images, etc.) published on the Steam Workshop,
including metadata, tags, votes, playtime statistics, and more.

API Documentation:

* [Steam API GetDetails v1][]
* [Steam API Reference by XPaw][]

To obtain an API key, visit: [Steam Dev API Key][].

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

  "github.com/woozymasta/steam/filedetails"
)

func main() {
  // Retrieve the Steam API key from environment variables.
  key, exists := os.LookupEnv("STEAM_API_KEY")
  if !exists {
    log.Fatal("STEAM_API_KEY environment variable is not set")
  }

  // Steam files ID
  fileIDs := []uint64{
    1559212036, // CF (mod)
    2545327648, // Dabs Framework (mod)
    2874283306, // DayZ launcher Linux (guide)
    3213284654, // DayZ (screenshots)
  }

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

[Steam API GetDetails v1]: https://api.steampowered.com/IPublishedFileService/GetDetails/v1/
[Steam API Reference by XPaw]: https://steamapi.xpaw.me/#IPublishedFileService/GetDetails
[Steam Dev API Key]: https://steamcommunity.com/dev/apikey
