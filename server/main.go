package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log/slog"
	"os"
)

var port *int

func init() {
	port = flag.Int("port", -1, "port to listen on")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// https://github.com/gin-gonic/examples/blob/master/websocket/server/server.go
// https://pkg.go.dev/github.com/gorilla/websocket
func main() {
	flag.Parse()
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("ws", func(c *gin.Context) {
		w, r := c.Writer, c.Request
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.ErrorContext(c, "websocket.upgrader.upgrade", "error", err)
			return
		}
		defer conn.Close()

		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.ErrorContext(c, "read message", "error", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			logger.ErrorContext(c, "write message", "error", err)
			return
		}
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
