package main

import (
	"fmt"
)

// StubMovieStore to implement the functionality of the real store
type StubMovieStore struct {
	movies Movies
}

// NewStubMovieStore initializes a new StubMovieStore
func NewStubMovieStore(movies Movies) *StubMovieStore {
	return &StubMovieStore{
		movies: movies,
	}
}

// GetMovies returns all the movies in the store
func (s *StubMovieStore) GetMovies() Movies {
	return s.movies
}

// GetMovie returns a single movie by ID
func (s *StubMovieStore) GetMovie(id string) (Movie, error) {
	var result Movie
	for _, movie := range s.movies {
		if movie.Id == id {
			result = movie
		}
	}

	if result.Id != id {
		return result, fmt.Errorf("Could not find movie with id %s", id)
	}

	return result, nil
}

// AddMovie takes a JSON object movie and appends to the movies store
func (s *StubMovieStore) AddMovie(movie Movie) {
	s.movies = append(s.movies, movie)
}
