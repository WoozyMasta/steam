package serverlist

import (
	"fmt"
	"os"
	"testing"

	"github.com/woozymasta/steam/utils/appid"
)

// Return Counter-Strike 2 servers
func TestGetCS2Servers(t *testing.T) {
	var id uint64 = appid.CounterStrike2.Uint64()

	filter := &Filter{}
	filter.Add(KeyAppID, fmt.Sprintf("%d", id))
	filter.Add(KeyMap, "de_dust2")
	filter.AddNor(KeyGameType, "insecure")
	filter.AddNor(KeyGameType, "empty")

	if err := helperGetServers(id, filter, 50); err != nil {
		t.Error(err)
	}
}

// Return Counter-Strike Source servers
func TestGetCSSServers(t *testing.T) {
	var id uint64 = appid.CounterStrikeSource.Uint64()

	filter := &Filter{}
	filter.Add(KeyAppID, fmt.Sprintf("%d", id))
	filter.Add(KeyMap, "de_dust2")

	if err := helperGetServers(id, filter, 50); err != nil {
		t.Error(err)
	}
}

// Return official servers of DayZ
func TestGetDayZServers(t *testing.T) {
	var id uint64 = appid.DayZ.Uint64()

	filter := &Filter{}
	filter.Add(KeyAppID, fmt.Sprintf("%d", id))
	filter.Add(KeyGameType, "battleye")
	filter.AddNor(KeyGameType, "external")

	if err := helperGetServers(id, filter, 200); err != nil {
		t.Error(err)
	}
}

// Return official servers of Arma3
func TestGetArma3Servers(t *testing.T) {
	var id uint64 = appid.Arma3.Uint64()

	filter := &Filter{}
	filter.Add(KeyAppID, fmt.Sprintf("%d", id))
	filter.Add(KeyName, "\xb6 [ OFFICIAL ] Arma 3 *")
	filter.Add(KeyGameType, "bt")
	filter.Add(KeyGameType, "dt")

	if err := helperGetServers(id, filter, 150); err != nil {
		t.Error(err)
	}
}

func helperGetServers(appId uint64, filter *Filter, limit int) error {
	key, ok := os.LookupEnv("STEAM_API_KEY")
	if !ok {
		return fmt.Errorf("Steam API key must be pass in variable 'STEAM_API_KEY'")
	}

	query := New(key)
	query.SetLimit(limit)
	servers, err := query.Get(filter)

	if err != nil {
		return fmt.Errorf("Cant get files %s", err)
	}
	if len(servers) > limit {
		return fmt.Errorf("Return %d servers, but expected %d by limit", len(servers), limit)
	}
	if len(servers) == 0 {
		return fmt.Errorf("Return 0 servers, but expected up to %d by limit", limit)
	}

	for _, s := range servers {
		fmt.Printf(
			"Game: %s, Address: %s, Title: \"%s\", Map: %s, Online: [%d/%d]\n",
			s.GameDir, s.Addr, s.Name, s.Map, s.Players, s.MaxPlayers,
		)

		if s.Appid != appId {
			return fmt.Errorf("Return %d (%s) game, but expected %d", s.Appid, s.GameDir, appId)
		}

		fmt.Printf("Keywords: %s\n", s.GameType)
	}

	for version, count := range servers.GetVersionMap() {
		fmt.Printf("Version %s used on %d servers\n", version, count)
	}

	return nil
}
