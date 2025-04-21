package server

import (
	"fmt"
	"net/http"
	"strings"
)

func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, GetPlayersScore(player))
}

func GetPlayersScore(player string) string {
	switch player {
	case "Pepper":
		return "20"
	case "Floyd":
		return "10"
	}
	return ""
}
