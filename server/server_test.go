package server_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	"github.com/carlosetorresm/tdd_go_web_server/server"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
)

func newPlayersRequest(method string, name string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestGETPlayers(t *testing.T) {
	store := test.NewStubPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}, nil, nil)
	server := server.NewPlayerServer(store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newPlayersRequest(http.MethodGet, "Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newPlayersRequest(http.MethodGet, "Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing Players", func(t *testing.T) {
		request := newPlayersRequest(http.MethodGet, "Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusNotFound

		test.AssertStatus(t, got, want)
	})
}

func TestScoreWins(t *testing.T) {
	store := test.NewStubPlayerStore(map[string]int{}, nil, nil)
	server := server.NewPlayerServer(store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		winner := "Pepper"
		request := newPlayersRequest(http.MethodPost, winner)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		test.AssertStatus(t, response.Code, http.StatusAccepted)
		test.AssertPlayerWin(t, store, winner)
	})
}

func newLeagueRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return request
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := league.League{
			*league.NewPlayer("Cleo", 32),
			*league.NewPlayer("Chris", 20),
			*league.NewPlayer("Tiest", 14),
		}

		store := test.NewStubPlayerStore(nil, nil, wantedLeague)
		serv := server.NewPlayerServer(store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		serv.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertLeague(t, got, wantedLeague)
		assertionContentType(t, response, "application/json")
	})
}

func assertionContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.Result().Header)
	}
}

func getLeagueFromResponse(t testing.TB, body io.Reader) (league league.League) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}
	return
}
