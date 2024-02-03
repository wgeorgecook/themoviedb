package server

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func initRouter() {
	if router == nil {
		router = gin.Default()
	}

	router.GET("/movies", getMovies)

}

func StartServer() {
	initRouter()
	router.Run()
}
