package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/websocket"
)

var chm = NewChanMux()

func main() {
	http.Handle("/socket", websocket.Handler(handler))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
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
