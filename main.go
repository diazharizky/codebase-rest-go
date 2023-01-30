package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/diazharizky/codebase-rest-go/config"
	"github.com/gin-gonic/gin"
)

func init() {
	config, err := config.LoadEnv(".")
	if err != nil {
		log.Fatal("unable to load config:", err)
	}

	fmt.Println("server host", config.AppHost)
	fmt.Println("server port", config.AppPort)
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
