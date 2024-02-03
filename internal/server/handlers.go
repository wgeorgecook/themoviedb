package server

import (
	"moviedb/internal/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMovies(c *gin.Context) {
	movies, err := db.GetMovies(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}
	c.IndentedJSON(http.StatusOK, movies)
}
