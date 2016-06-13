package main

import (
	"log"

	"golang.org/x/net/websocket"
)

const (
	origin = "http://kill-trigger.herokuapp.com"
	wsURL  = "ws://kill-trigger.herokuapp.com/socket"
)

func main() {
	ws, err := websocket.Dial(wsURL, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan byte, 1)
	go func(ch chan byte) {
		for b := range ch {
			do(b)
		}
	}(ch)
	for {
		buf := make([]byte, 1)
		_, err = ws.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		ch <- buf[0]
	}
}
