package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewMovieServer(NewStubMovieStore(Movies{
		{"Home Alone", 8.5},
		{"Home Alone 2", 7},
	}))

	t.Run("get all movies", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/movies", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var movies Movies
		err := json.NewDecoder(response.Body).Decode(&movies)
		if err != nil {
			t.Errorf("Decoding JSON failed, %v", err)
		}

		want := Movies{
			{"Home Alone", 8.5},
			{"Home Alone 2", 7},
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertDeepEqual(t, movies, want)
	})
}
