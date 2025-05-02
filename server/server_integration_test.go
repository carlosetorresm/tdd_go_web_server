package server_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	filesystem "github.com/carlosetorresm/tdd_go_web_server/domain/file_system"
	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]")
	defer cleanDatabase()

	store, err := filesystem.NewFileSystemPlayerStore(database)
	assertNoError(t, err)
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
		want := league.League{
			{Name: "Pepper", Wins: 3},
		}
		assertLeague(t, got, want)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn' expect an erro but got one, %v", err)
	}
}

func createTempFile(t testing.TB, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpFile.Write([]byte(initialData))
	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}
	return tmpFile, removeFile
}
