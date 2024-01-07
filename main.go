package main

import (
	"fmt"

	moviedb "moviedb/internal/moviedb"
	utils "moviedb/internal/utils"
)

var movieClient *moviedb.Api

func main() {
	fmt.Println("Hello!")
	defer fmt.Println("Goodbye!")

	cfg := utils.LoadEnv()
	movieId, listId := utils.ValidateInitParams(cfg.Token, cfg.MovieID, cfg.ListId)

	fmt.Println("initializing client")
	// init the client
	movieClient = moviedb.Init(cfg.Token)

	fmt.Println("adding movie to list")
	// add povided movie to the provided list
	if err := movieClient.AddMovieToList(movieId, listId); err != nil {
		panic(fmt.Sprintf("could not add to list: %v", err))
	}

	// exit!
	fmt.Println("added movie to list")
}
