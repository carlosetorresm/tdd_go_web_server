package server_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	league "github.com/carlosetorresm/tdd_go_web_server/infraestructure"
	"github.com/carlosetorresm/tdd_go_web_server/server"
	test "github.com/carlosetorresm/tdd_go_web_server/testing"
	"github.com/gorilla/websocket"
)

var dummyGame = &test.GameSpy{}

func newPlayersRequest(method string, name string) *http.Request {
	request, _ := http.NewRequest(method, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func TestGETPlayers(t *testing.T) {
	store := test.NewStubPlayerStore(map[string]int{
		"Pepper": 20,
		"Floyd":  10,
	}, nil, nil)
	server := mustMakePlayerServer(t, store, dummyGame)

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
	server := mustMakePlayerServer(t, store, dummyGame)

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
		serv := mustMakePlayerServer(t, store, dummyGame)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		serv.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		test.AssertStatus(t, response.Code, http.StatusOK)
		test.AssertLeague(t, got, wantedLeague)
		assertionContentType(t, response, "application/json")
	})
}

func newGameRequest() *http.Request {
	request, _ := http.NewRequest(http.MethodGet, "/game", nil)
	return request
}

func TestGame(t *testing.T) {
	t.Run("GET /game return 200", func(t *testing.T) {
		server := mustMakePlayerServer(t, &test.StubPlayerStore{}, dummyGame)

		request := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		test.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("when we get a message over a websocket it is a winner of a game", func(t *testing.T) {
		store := &test.StubPlayerStore{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, store, dummyGame))
		defer server.Close()

		wsUrl := "ws" + strings.TrimPrefix(server.URL, "http") + "/ws"

		ws := mustDialWS(t, wsUrl)
		defer ws.Close()

		writeWSMessage(t, ws, winner)
		time.Sleep(10 * time.Millisecond)
		test.AssertPlayerWin(t, store, winner)
	})

	t.Run("start game with 3 players and finish game with 'Ruth' as winner", func(t *testing.T) {
		game := &test.GameSpy{}
		dummyStore := &test.StubPlayerStore{}
		winner := "Ruth"
		server := httptest.NewServer(mustMakePlayerServer(t, dummyStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")

		defer server.Close()
		defer ws.Close()

		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		time.Sleep(10 * time.Millisecond)
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
	})
}

func writeWSMessage(t *testing.T, ws *websocket.Conn, message string) {
	t.Helper()
	if err := ws.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func mustDialWS(t *testing.T, wsUrl string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", wsUrl, err)
	}
	return ws
}

func mustMakePlayerServer(t *testing.T, store server.PlayerStore, game server.Game) *server.PlayerServer {
	server, err := server.NewPlayerServer(store, game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
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

func assertGameStartedWith(t *testing.T, game *test.GameSpy, numberOfPlayers int) {
	t.Helper()
	if game.StartedWith != numberOfPlayers {
		t.Errorf("wanted Start called with %d but got %d", numberOfPlayers, game.StartedWith)
	}
}

func assertFinishCalledWith(t *testing.T, game *test.GameSpy, winner string) {
	t.Helper()
	if game.FinishedWith != winner {
		t.Errorf("expected finish called with %q but got %q", winner, game.FinishedWith)
	}
}
