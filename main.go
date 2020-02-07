package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/ping", func(c *gin.Context) {
		c.JSON(200, nil)
	})

	return r
}

func main() {
	r := setupRouter()
	log.Fatal(r.Run(":9000"))
}
