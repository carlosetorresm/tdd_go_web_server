package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	inmemoryserver "github.com/carlosetorresm/tdd_go_web_server/domain/interactions"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := inmemoryserver.NewInMemoryPlayerStore()
	server := server.PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRequest(http.MethodGet, player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
