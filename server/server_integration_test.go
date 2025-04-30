package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	interations "github.com/carlosetorresm/tdd_go_web_server/domain/interactions"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := interations.NewInMemoryPlayerStore()
	serv := server.NewPlayerServer(store)
	player := "Pepper"

	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		serv.ServeHTTP(response, newPlayersRequest(http.MethodGet, player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		serv.ServeHTTP(response, newLeagueRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := []league.Player{
			{Name: "Pepper", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}
