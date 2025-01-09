# latest

A lightweight Go package to help you automatically update game servers
(especially Steam-based) by analyzing version usage across a set of servers.
Rather than blindly picking the newest version—which could be a test release
or a spoofed version—latest lets you specify thresholds (in percentages) to
identify when a certain version is widely adopted enough to be considered
safe for an automatic update. Why Use latest?

* **Prevent Early Adoptions**: Avoid picking a brand-new, potentially unstable
  build just because it has a higher version number.
* **Ignore Spoofed Versions**: Some servers might falsely advertise a higher
  version. If they don’t meet a certain adoption threshold, they get
  filtered out.
* **Fallback Mechanism**: If no version meets your main threshold, the package
  lets you define a fallback threshold. Only if no version meets the
  fallback either, it defaults to the absolute latest version.

## Installation

Install the package using `go get`:

```bash
go get github.com/woozymasta/steam
```

## Usage

In this example, you want at least 40% of servers to have switched to the
new version. If no version meets 40%, FindVersion automatically tries a
fallback threshold of 24% (60% of 40). If still no version qualifies, it
returns the highest version.

```go
import (
    "fmt"
    "github.com/woozymasta/steam/utils/latest"
)

func main() {
    versionMap := map[string]uint32{
        "1.0.0":  10,
        "1.1.0":  25,
        "2.0.0":  5,
    }

    version, err := latest.FindVersion(versionMap, 40.0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Selected version:", version)
}
```

You will find even more examples of explanations of work in different
situations in the [file with tests](latest_test.go)

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
