package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	"github.com/carlosetorresm/tdd_go_web_server/server"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := test.CreateTempFile(t, "[]")
	defer cleanDatabase()

	store, err := filesystem.NewFileSystemPlayerStore(database)
	test.AssertNoError(t, err)
	serv, _ := server.NewPlayerServer(store, dummyGame)
	player := "Pepper"

	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))
	serv.ServeHTTP(httptest.NewRecorder(), newPlayersRequest(http.MethodPost, player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		serv.ServeHTTP(response, newPlayersRequest(http.MethodGet, player))

		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		serv.ServeHTTP(response, newLeagueRequest())
		test.AssertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response.Body)
		want := league.League{
			*league.NewPlayer(player, 3),
		}
		test.AssertLeague(t, got, want)
	})
}
