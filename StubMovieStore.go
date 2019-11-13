package main

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