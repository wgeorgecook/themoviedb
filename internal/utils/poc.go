package utils

import (
	"context"
	"fmt"

	"moviedb/internal/db"
	moviedb "moviedb/internal/moviedb"
)

func TestModifyList(cfg InitConfig) {
	movieId, listId := validateModifyListParams(cfg.Token, cfg.MovieID, cfg.ListId)

	fmt.Println("initializing moviedb client")
	movieClient := moviedb.Init(cfg.Token)

	fmt.Println("adding movie to list")
	if err := movieClient.AddMovieToList(movieId, listId); err != nil {
		panic(fmt.Sprintf("could not add to list: %v", err))
	}
}

func TestCreateMovie(cfg InitConfig) {
	movieId, movieName := validateCreateMovieParams(cfg.MovieID, cfg.MovieName)
	fmt.Println("inserting movie")
	if err := db.CreateMovie(context.TODO(), &db.Movie{ID: int64(movieId), Name: movieName}); err != nil {
		panic(fmt.Sprintf("could not insert movie: %v", err))
	}
	fmt.Println("insert successful, checking created movie")
	m, err := db.GetMovie(context.TODO(), int64(movieId))
	if err != nil {
		panic(fmt.Sprintf("could not get movie: %v", err))
	}
	fmt.Printf("got movie: %+v", m)
}
