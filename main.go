package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"

	"github.com/alexyou8021/sleeper-wrapper.git/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default()
	router.Use(CORSMiddleware())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "success")
	})
	router.GET("/refreshSleeperPlayersTable", func(ctx *gin.Context) {
		handlers.CreateSleeperPlayersTable(ctx)
	})
	router.GET("/refreshESPNPlayersTable", func(ctx *gin.Context) {
		handlers.CreateESPNPlayersTable(ctx)
	})
	router.GET("/refreshStatsTable", func(ctx *gin.Context) {
		handlers.CreateStatsTable(ctx)
	})
	router.GET("/user/sleeper/:name", func(ctx *gin.Context) {
		handlers.Handler(ctx)
	})
	router.GET("/user/espn/:id", func(ctx *gin.Context) {
		handlers.GetESPNTransactions(ctx)
	})

	router.Run(":" + port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		fmt.Println(c.Request.Method)

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
