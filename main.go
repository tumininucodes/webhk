package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("webhook", func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{"success": "webhook spun here"})
	})

	port := os.Getenv("PORT")
	if port != "" {
		server.Run(port)
	} else {
		server.Run(":5050")
	}
	

}
