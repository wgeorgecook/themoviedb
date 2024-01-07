package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type initConfig struct {
	token   string
	movieID string
	listId  string
}

func loadEnv() initConfig {
	// Load environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("could not load .env file: %v\n", err)
	}

	bearerToken := os.Getenv("BEARER_TOKEN")
	movieIdStr := os.Getenv("MOVIE_ID")
	listIDStr := os.Getenv("LIST_ID")

	return initConfig{
		token:   bearerToken,
		movieID: movieIdStr,
		listId:  listIDStr,
	}
}

func validateInitParams(bearerToken, movieIdStr, listIDStr string) (int, int) {
	if movieIdStr == "" {
		panic("must provide MOVIE_ID")
	}

	if bearerToken == "" {
		panic("must provide BEARER_TOKEN")
	}

	if listIDStr == "" {
		panic("must provide LIST_ID")
	}

	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		panic("MOVIE_ID must be type int")
	}

	listId, err := strconv.Atoi(listIDStr)
	if err != nil {
		panic("LIST_ID must be type int")
	}

	return movieId, listId
}
