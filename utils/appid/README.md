# appid

Package AppID provides a collection of constants representing Steam
application IDs.

These constants encompass a variety of games across different engines
and categories, including those built on the GoldSource and Source
1/2 engines, as well as other popular titles.

References:

* <https://api.steampowered.com/ISteamApps/GetAppList/v2/>
* <https://api.steampowered.com/IStoreService/GetAppList/v1/>
* <https://steamdb.info/>

## Get more ID's

Use bash script [./get.sh](./get.sh) for get all application ID only for
games in Steam

```bash
./get.sh $api_key apps.json
# or
STEAM_KEY=abc OUT_FILE=apps.json ./get.sh
```

## Installation

Install the package using `go get`:

```bash
go get github.com/woozymasta/steam
```

## Usage

```go
import (
  "fmt"
  "github.com/woozymasta/steam/utils/appid"
)

func main() {
  fmt.Println("Dota 2 application ID: %d\n", appid.Dota2)
  fmt.Println("Application name for const Dota2: %s\n", appid.Dota2.String())
  fmt.Println("Application name for id 570: %s\n", appid.AppID(570).String())
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
