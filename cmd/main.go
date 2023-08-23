package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.POST("webhook", func(ctx *gin.Context) {
		var webhookData map[string]interface{}
		ctx.ShouldBindJSON(&webhookData)
		fmt.Println("webhook data -> ", webhookData)
	})

	server.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"success": "the world is yours for the taking"})
	})

	server.Run(":8080")

}
