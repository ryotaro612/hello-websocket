package main

import (
	"flag"
	"github.com/gin-gonic/gin"
)

var port *int

func init() {
	port = flag.Int("port", -1, "port to listen on")
}

func main() {
	flag.Parse()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
