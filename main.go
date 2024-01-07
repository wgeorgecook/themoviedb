package main

import (
	"fmt"

	moviedb "moviedb/internal/moviedb"
)

var movieClient *moviedb.Api

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("Goodbye!")

	cfg := loadEnv()
	movieId, listId := validateInitParams(cfg.token, cfg.movieID, cfg.listId)

	fmt.Println("initializing client")
	// init the client
	movieClient = moviedb.Init(cfg.token)

	fmt.Println("adding movie to list")
	// add "The Nightmare Before Christmas" to the list
	if err := movieClient.AddMovieToList(movieId, listId); err != nil {
		panic(fmt.Sprintf("could not add to list: %v", err))
	}

	// exit!
	fmt.Println("added movie to list")
}
