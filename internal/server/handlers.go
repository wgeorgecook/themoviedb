package server

import (
	"database/sql"
	"log"
	"moviedb/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	Message string `json:"message"`
}

func getMovies(c *gin.Context) {
	movies, err := db.GetMovies(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovie(c *gin.Context) {
	var movie db.Movie
	if err := c.ShouldBind(&movie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, httpError{Message: "request could not marshal"})
		return
	}

	// query the db cache first
	found, err := db.GetMovie(c.Request.Context(), movie)
	if err != nil && err != sql.ErrNoRows {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	if found != nil {
		c.IndentedJSON(http.StatusOK, found)
		return
	}

	// search on themoviedb to get another movie
	searchResults, err := movieDbClient.SearchMovies(c.Request.Context(), movie.Name)
	if err != nil {
		c.IndentedJSON(http.StatusFailedDependency, httpError{Message: err.Error()})
		return
	}
	for _, result := range searchResults {
		if err := db.UpsertMovie(c.Request.Context(), &db.Movie{
			ID:   int64(result.ID),
			Name: result.Title,
		}); err != nil {
			log.Printf("could not upsert %s: %s", result.Title, err.Error())
		}
	}

	c.IndentedJSON(http.StatusOK, searchResults)

}
