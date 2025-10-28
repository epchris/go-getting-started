package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// Emit a log line to stdout every minute
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	heartbeatLogger := log.New(os.Stdout, "", log.LstdFlags)
	go func() {
		for range ticker.C {
			heartbeatLogger.Println("minute tick")
		}
	}()

	router.Run(":" + port)
}
