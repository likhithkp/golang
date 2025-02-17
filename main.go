package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func greet(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Welcome to Gin",
	})
}

func ginQuery(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"data": "Welcome to Gin",
		"name": name,
		"age":  age,
	})
}

func ginUrl(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"data": "Welcome to ginUrl",
		"name": name,
		"age":  age,
	})
}

func ginPost(c *gin.Context) {
	value, _ := io.ReadAll(c.Request.Body)
	c.JSON(200, gin.H{
		"data":     "Welcome to ginPost",
		"bodyData": string(value),
	})
}

func main() {
	router := gin.Default()

	file, _ := os.Create("goLogger.log")
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)

	router.GET("/", greet)

	auth := gin.BasicAuth(gin.Accounts{
		"user": "pass1",
	})

	query := router.Group("/query")
	{
		query.GET("/ginQuery", ginQuery)
	}
	url := router.Group("/url", auth)
	{
		url.GET("/ginUrl/:name/:age", ginUrl)
	}
	post := router.Group("/posy")
	{
		post.POST("/ginPost", ginPost)
	}

	server := &http.Server{
		Addr:         ":5000",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server.ListenAndServe()
}
