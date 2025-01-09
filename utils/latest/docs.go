/*
Package latest provides a threshold-based version selection mechanism
particularly helpful for automatically updating Steam-based game servers
when a critical mass of official servers has moved to a newer version.

This approach is more reliable than simply picking the highest version, because:
 1. There may be test or beta versions that do not represent a stable release.
 2. Some unofficial or spoofed servers might artificially report versions higher than the real official versions.

By defining a main threshold (in percent) for how many servers must be running a certain version, you can effectively:
  - Avoid updating prematurely to a test build.
  - Filter out incorrect (or maliciously spoofed) versions.

The package also supports a fallback threshold, which is a secondary, lower threshold that can be used
in case no version meets the main threshold. If no version meets either threshold, you still have the option
to fall back to the absolute latest version in the provided data (e.g., if only a very small percentage of servers
have the new version, but it is in fact the correct official release).

# How it works:
 1. You provide a map of versions to the number of servers using them.
 2. The system calculates the total server count and checks if any versions meet or exceed your threshold (percent).
 3. If no version meets the main threshold, we check the fallback threshold.
 4. If still no version meets the fallback, we choose the highest version from the map, assuming that
    the data provided is reliable enough to identify the correct release.

This logic helps server owners safely and automatically update their servers while avoiding unstable or spoofed versions.

# Usage:

	import (
		"fmt"
		"github.com/woozymasta/steam/utils/latest"
	)

	func main() {
		versionMap := map[string]uint32{
			"1.0.0": 10,
			"1.1.0": 25,
			"2.0.0": 5,
		}

		version, err := latest.FindVersion(versionMap, 40.0)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Selected version:", version)
	}
*/
package latest
