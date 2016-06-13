package main

import (
	"html/template"
	"log"
	"os"
	"strconv"

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
	for b := range ch {
		_, err := ws.Write([]byte{b})
		if err != nil {
			log.Println(err)
			break
		}
	}
	chm.Delete(ch)
	ws.Close()
}
