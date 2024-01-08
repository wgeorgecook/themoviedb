package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type initConfig struct {
	Token       string
	MovieID     string
	ListId      string
	MovieName   string
	PostgresURI string
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
	movieName := os.Getenv("MOVIE_NAME")
	postgresUri := os.Getenv("POSTGRES_URI")

	return initConfig{
		Token:       bearerToken,
		MovieID:     movieIdStr,
		ListId:      listIDStr,
		MovieName:   movieName,
		PostgresURI: postgresUri,
	}
}

func validateModifyListParams(bearerToken, movieIdStr, listIDStr string) (int, int) {
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

func validateCreateMovieParams(movieIdStr, movieName string) (int, string) {
	if movieIdStr == "" {
		panic("must provide MOVIE_ID")
	}

	if movieName == "" {
		panic("must provide MOVIE_NAME")
	}

	movieId, err := strconv.Atoi(movieIdStr)
	if err != nil {
		panic("MOVIE_ID must be type int")
	}

	return movieId, movieName
}
