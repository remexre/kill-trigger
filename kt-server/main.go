package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"golang.org/x/net/websocket"
)

var chm = NewChanMux()

func main() {
	r := gin.Default()

	r.LoadHTMLFiles("assets/index.html")
	r.StaticFile("agora.js", "assets/agora.js")
	r.StaticFile("agora.js.map", "assets/agora.js.map")
	r.StaticFile("main.js", "assets/main.js")

	r.GET("/", indexHandler)
	r.GET("/api/numUsers", numUsersHandler)
	r.POST("/api/send", sendHandler)
	r.Any("/socket", gin.WrapH(websocket.Handler(handler)))

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func numUsersHandler(c *gin.Context) {
	c.String(200, "%d", chm.Len())
}

func sendHandler(c *gin.Context) {
	var data map[string]string
	c.BindJSON(&data)
	chm.Send(data)
	c.JSON(200, data)
}

func handler(ws *websocket.Conn) {
	enc := json.NewEncoder(ws)

	ch := chm.NewChan()
	ticker := time.NewTicker(5 * time.Second)
	stop := false
	for !stop {
		select {
		case val := <-ch:
			if err := enc.Encode(val); err != nil {
				log.Println(err)
				stop = true
			}
		case <-ticker.C:
			ch <- map[string]string{"name": "keepalive"}
		}
	}
	ticker.Stop()
	chm.Delete(ch)
	ws.Close()
}
