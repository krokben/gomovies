package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	server := NewMovieServer(NewStubMovieStore(Movies{
		{"id1", "Home Alone", 8.5},
		{"id2", "Home Alone 2", 7},
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
			{"id1", "Home Alone", 8.5},
			{"id2", "Home Alone 2", 7},
		}

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertDeepEqual(t, movies, want)
	})

	t.Run("get one movie", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/movies/id1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var movie Movie
		err := json.NewDecoder(response.Body).Decode(&movie)
		if err != nil {
			t.Errorf("Decoding JSON failed, %v", err)
		}

		want := Movie{"id1", "Home Alone", 8.5}

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertDeepEqual(t, movie, want)
	})

	t.Run("404 on missing movie", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/movies/id3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})

	t.Run("post a movie", func(t *testing.T) {
		movie, _ := json.Marshal(Movie{"id3", "Home Alone 3", 5})
		request, _ := http.NewRequest(http.MethodPost, "/movies", bytes.NewReader(movie))
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
	})
}
