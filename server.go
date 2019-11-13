package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const jsonContentType = "application/json"

// MovieStore is an interface for a movie store
type MovieStore interface {
	GetMovies() Movies
	GetMovie(id string) Movie
}

// MovieServer takes a store of movies
type MovieServer struct {
	store MovieStore
	http.Handler
}

// NewMovieServer initializes a new MovieServer
func NewMovieServer(store MovieStore) *MovieServer {
	s := new(MovieServer)
	s.store = store

	router := http.NewServeMux()
	router.Handle("/movies", http.HandlerFunc(s.moviesHandler))
	router.Handle("/movies/", http.HandlerFunc(s.movieHandler))

	s.Handler = router

	return s
}

func (s *MovieServer) moviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)

	err := json.NewEncoder(w).Encode(s.store.GetMovies())
	if err != nil {
		log.Fatal("Could not encode into JSON", err)
	}
}

func (s *MovieServer) movieHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/movies/"):]

	w.Header().Set("content-type", jsonContentType)

	err := json.NewEncoder(w).Encode(s.store.GetMovie(id))
	if err != nil {
		log.Fatal("Could not encode into JSON", err)
	}
}
