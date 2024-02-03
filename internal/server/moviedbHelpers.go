package server

import "moviedb/internal/moviedb"

var movieDbClient *moviedb.Api

func initMovieDbClient(token string) {
	movieDbClient = moviedb.Init(token)
}
