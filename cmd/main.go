package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("webhook", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "webhook spun here"})
	})

	server.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "heher spun here"})
	})

	server.Run(":8080")

}
