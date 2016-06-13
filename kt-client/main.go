package main

import (
	"log"

	"golang.org/x/net/websocket"
)

const (
	origin = "http://kill-trigger.herokuapp.com"
	url    = "ws://kill-trigger.herokuapp.com/socket"
)

func main() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan byte, 1)
	go func(ch chan byte) {
		for b := range ch {
			process(b)
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

func process(b byte) {
	switch b {
	case 0x00:
		log.Println("Hello, world!")
	default:
		log.Printf("Unknown command: %d", b)
	}
}
