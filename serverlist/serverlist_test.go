package serverlist

import (
	"fmt"
	"os"
	"testing"
)

func TestGetServers(t *testing.T) {
	key, ok := os.LookupEnv("STEAM_API_KEY")
	if !ok {
		t.Error("Steam API key must be pass in variable 'STEAM_API_KEY'")
	}

	limit := 1000

	query := New(key)
	query.SetLimit(limit)

	filter := &Filter{}
	filter.Add(KeyAppID, "221100")
	filter.Add(KeyGameType, "battleye")
	filter.AddNor(KeyGameType, "external")

	servers, err := query.Get(filter)

	if err != nil {
		t.Errorf("Cant get files %v", err)
	}
	if len(servers) > limit {
		t.Errorf("Return %d servers, but expected %d by limit", len(servers), limit)
	}

	for _, s := range servers {
		fmt.Printf(
			"Game: %s, Address: %s, Title: \"%s\", Map: %s, Online: [%d/%d]\n",
			s.GameDir, s.Addr, s.Name, s.Map, s.Players, s.MaxPlayers,
		)

		if s.Appid != 221100 {
			t.Errorf("Return %d (%s) game, but expected %d", s.Appid, s.GameDir, 221100)
		}
	}

	for version, count := range servers.GetVersionMap() {
		fmt.Printf("Version %s used on %d servers\n", version, count)
	}
}
