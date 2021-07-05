package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/ping", hello)
	v1.GET("/cards", getCards)

	r.Run(":8080")
}
