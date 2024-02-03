package server

import (
	"database/sql"
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
	found, err := db.GetMovie(c.Request.Context(), movie)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			c.IndentedJSON(http.StatusNotFound, httpError{Message: "movie not found"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	c.IndentedJSON(http.StatusOK, found)

}
