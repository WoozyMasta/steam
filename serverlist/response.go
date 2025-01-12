package serverlist

import (
	"strings"

	json "github.com/json-iterator/go"
)

// GameType represents a slice of game type strings, parsed from a comma-separated string.
type GameType []string

// UnmarshalJSON implements the json.Unmarshaler interface for the GameType type.
// It splits a comma-separated string into a slice of trimmed strings.
func (gt *GameType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parts := strings.Split(s, ",")
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	*gt = parts
	return nil
}

// Servers represents the structure of list a game servers as returned by the Steam Master Server Query Protocol.
type Servers []Server

// Server represents the structure of a game server as returned by the Steam Master Server Query Protocol.
type Server struct {
	Addr       string   `json:"addr"`             // Server address in the format IP:Port
	GameDir    string   `json:"gamedir"`          // Game directory (e.g., "cstrike" for Counter-Strike)
	Map        string   `json:"map"`              // Current map on the server
	Name       string   `json:"name"`             // Server name
	Product    string   `json:"product"`          // Product or game associated with the server
	Version    string   `json:"version"`          // Game version running on the server
	OS         string   `json:"os"`               // Server operating system ("l" for Linux, "w" for Windows)
	GameType   GameType `json:"gametype"`         // Game type or server tags (from sv_tags or "battleye,lqs0,etm2.300000")
	SteamID    uint64   `json:"steamid,string"`   // Server SteamID
	Appid      uint64   `json:"appid"`            // Game ID (e.g., 221100 for DayZ)
	Bots       uint16   `json:"bots,omitempty"`   // Number of bots on the server
	GamePort   uint16   `json:"gameport"`         // Game port (for client connections)
	MaxPlayers uint16   `json:"max_players"`      // Maximum number of players on the server
	Players    uint16   `json:"players"`          // Current number of players on the server
	Region     int      `json:"region,omitempty"` // Server region code
	Dedicated  bool     `json:"dedicated"`        // Indicates if the server is dedicated
	Secure     bool     `json:"secure"`           // Indicates if protection is enabled (e.g., VAC or BattlEye)
}

// GetVersionMap populates the version map from the server list
// Returns map of version -> count of occurrences
func (s *Servers) GetVersionMap() map[string]uint32 {
	versionMap := make(map[string]uint32)
	for _, srv := range *s {
		versionMap[srv.Version]++
	}

	return versionMap
}
