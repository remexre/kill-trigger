package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/remexre/kill-trigger"

	"golang.org/x/net/websocket"
)

var chm = NewChanMux()

func main() {
	r := gin.Default()

	html := template.Must(template.New("index.html").Parse(index))
	r.SetHTMLTemplate(html)

	r.GET("/", indexHandler)
	r.GET("/api/numUsers", numUsersHandler)
	r.POST("/api/:code/send", sendHandler)
	r.Any("/socket", gin.WrapH(websocket.Handler(handler)))

	if err := r.Run(":" + os.Getenv("PORT")); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"commands": kt.Commands,
	})
}

func numUsersHandler(c *gin.Context) {
	c.String(200, "%d", chm.Len())
}

func sendHandler(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.ParseUint(codeStr, 10, 8)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	chm.Send(byte(code))
}

func handler(ws *websocket.Conn) {
	ch := chm.NewChan()
	ticker := time.NewTicker(5 * time.Second)
	stop := false
	for !stop {
		select {
		case b := <-ch:
			_, err := fmt.Fprint(ws, b)
			if err != nil {
				log.Println(err)
				stop = true
			}
		case <-ticker.C:
			ch <- kt.KeepAlive.ID
		}
	}
	ticker.Stop()
	chm.Delete(ch)
	ws.Close()
}
