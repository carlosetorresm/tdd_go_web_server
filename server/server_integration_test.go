package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	interations "github.com/carlosetorresm/tdd_go_web_server/domain/interactions"
	"github.com/carlosetorresm/tdd_go_web_server/server"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := interations.NewInMemoryPlayerStore()
	server := server.NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))
	server.ServeHTTP(httptest.NewRecorder(), newRequest(http.MethodPost, player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newRequest(http.MethodGet, player))

	assertStatus(t, response.Code, http.StatusOK)
	assertResponseBody(t, response.Body.String(), "3")
}
