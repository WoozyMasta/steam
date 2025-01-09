package serverlist

import (
	"fmt"
	"strings"
)

// FilterKey represents the keys used for filtering server lists in API requests.
type FilterKey string

const (
	// KeyDedicated filters for dedicated servers.
	//  - 1: Dedicated servers
	//  - 0: Non-dedicated servers
	KeyDedicated FilterKey = "dedicated"

	// KeySecure filters for servers using anti-cheat technologies (e.g., VAC, BattlEye).
	//  - 1: Secure servers
	//  - 0: Insecure servers
	KeySecure FilterKey = "secure"

	// KeyGameDir filters for servers running a specific game modification or directory.
	//  - str: Game modification or directory (e.g., "tf", "cstrike", "dayz")
	KeyGameDir FilterKey = "gamedir"

	// KeyMap filters for servers running a specific map.
	//  - str: Specific map name (e.g., "ctf_2fort")
	KeyMap FilterKey = "map"

	// KeyLinux filters for servers running on Linux.
	//  - 1: Linux servers
	//  - 0: Non-Linux servers
	KeyLinux FilterKey = "linux"

	// KeyPassword filters for servers based on password protection.
	//  - 1: Servers protected by a password
	//  - 0: Servers without password protection
	KeyPassword FilterKey = "password"

	// KeyEmpty filters for servers based on their occupancy.
	//  - 1: Servers that are not empty
	//  - 0: Servers that are empty
	KeyEmpty FilterKey = "empty"

	// KeyFull filters for servers based on their player capacity.
	//  - 1: Servers that are not full
	//  - 0: Servers that are full
	KeyFull FilterKey = "full"

	// KeyProxy filters for servers acting as proxy servers.
	//  - 1: Proxy servers
	//  - 0: Non-proxy servers
	KeyProxy FilterKey = "proxy"

	// KeyAppID filters for servers running a specific game ID.
	//  - id: Specific game ID (e.g., "221100" for DayZ)
	KeyAppID FilterKey = "appid"

	// KeyNotAppID filters for servers not running a specific game ID.
	//  - id: Game ID to exclude (e.g., "221100" to exclude DayZ servers)
	KeyNotAppID FilterKey = "napp"

	// KeyNoPlayers filters for servers based on player count.
	//  - 1: Servers with no players (empty)
	//  - 0: Servers with one or more players
	KeyNoPlayers FilterKey = "noplayers"

	// KeyWhite filters for servers using a whitelist.
	//  - 1: Servers with a whitelist enabled
	//  - 0: Servers without a whitelist
	KeyWhite FilterKey = "white"

	// KeyGameType filters for servers with a specific game type tag.
	//  - str: Game type tag defined in sv_tags (e.g., "battleye")
	KeyGameType FilterKey = "gametype"

	// KeyGameData filters for servers with specific hidden game data.
	//  - str: Hidden game data tag (e.g., "tag1") use in L2D2
	KeyGameData FilterKey = "gamedata"

	// KeyName filters for servers with matching names, supporting wildcards.
	//  - str: Server name with optional wildcards (e.g., "*server*")
	KeyName FilterKey = "name_match"

	// KeyVersion filters for servers with a specific version, supporting wildcards.
	//  - str: Server version with optional wildcards (e.g., "1.3.*")
	KeyVersion FilterKey = "version_match"

	// KeySingleAddr filters to include only one server per unique IP address.
	//  - 1: Include only one server per unique IP
	//  - 0: Include multiple servers per IP
	KeySingleAddr FilterKey = "collapse_addr_hash"

	// KeyGameAddr filters for servers with a specific IP address (port is optional).
	//  - ip: Specific IP address (e.g., "192.168.1.1") with optional port (e.g., "192.168.1.1:27015")
	KeyGameAddr FilterKey = "gameaddr"
)

// Filter is used to build filter conditions for API requests.
// It supports standard, NOR, and NAND filter conditions.
type Filter struct {
	conditions     []string
	norConditions  []string
	nandConditions []string
}

// Add adds a standard filter condition to the Filter.
// It appends the condition in the format "key\value".
func (f *Filter) Add(key FilterKey, value string) {
	f.conditions = append(f.conditions, fmt.Sprintf("%s\\%s", key, value))
}

// AddNor adds a NOR filter condition to the Filter.
// It appends the condition in the format "key\value" to NOR conditions.
func (f *Filter) AddNor(key FilterKey, value string) {
	f.norConditions = append(f.norConditions, fmt.Sprintf("%s\\%s", key, value))
}

// AddNand adds a NAND filter condition to the Filter.
// It appends the condition in the format "key\value" to NAND conditions.
func (f *Filter) AddNand(key FilterKey, value string) {
	f.nandConditions = append(f.nandConditions, fmt.Sprintf("%s\\%s", key, value))
}

// Remove removes a standard filter condition from the Filter.
// It searches for the condition "key\value" and removes it if found.
func (f *Filter) Remove(key FilterKey, value string) {
	toRemove := fmt.Sprintf("%s\\%s", key, value)
	for i, condition := range f.conditions {
		if condition == toRemove {
			f.conditions = append(f.conditions[:i], f.conditions[i+1:]...)
			break
		}
	}
}

// RemoveNor removes a NOR filter condition from the Filter.
// It searches for the condition "key\value" in NOR conditions and removes it if found.
func (f *Filter) RemoveNor(key FilterKey, value string) {
	toRemove := fmt.Sprintf("%s\\%s", key, value)
	for i, condition := range f.norConditions {
		if condition == toRemove {
			f.norConditions = append(f.norConditions[:i], f.norConditions[i+1:]...)
			break
		}
	}
}

// RemoveNand removes a NAND filter condition from the Filter.
// It searches for the condition "key\value" in NAND conditions and removes it if found.
func (f *Filter) RemoveNand(key FilterKey, value string) {
	toRemove := fmt.Sprintf("%s\\%s", key, value)
	for i, condition := range f.nandConditions {
		if condition == toRemove {
			f.nandConditions = append(f.nandConditions[:i], f.nandConditions[i+1:]...)
			break
		}
	}
}

// String constructs the filter string from the Filter's conditions.
// It ensures that both NOR and NAND conditions are not used simultaneously.
// Returns the constructed filter string or an error if validation fails.
func (f *Filter) String() (string, error) {
	if len(f.norConditions) > 0 && len(f.nandConditions) > 0 {
		return "", fmt.Errorf("cannot use both nor and nand conditions in the same filter")
	}

	var parts []string
	if len(f.conditions) > 0 {
		parts = append(parts, strings.Join(f.conditions, "\\"))
	}
	if len(f.norConditions) > 0 {
		parts = append(parts, fmt.Sprintf("nor\\%d\\%s", len(f.norConditions), strings.Join(f.norConditions, "\\")))
	}
	if len(f.nandConditions) > 0 {
		parts = append(parts, fmt.Sprintf("nand\\%d\\%s", len(f.nandConditions), strings.Join(f.nandConditions, "\\")))
	}

	return strings.Join(parts, "\\"), nil
}
