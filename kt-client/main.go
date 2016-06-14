package main

import (
	"log"
	"strconv"

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

	for {
		var codeStr string
		err := websocket.Message.Receive(ws, &codeStr)
		if err != nil {
			log.Fatal(err)
		}

		code, err := strconv.ParseUint(codeStr, 10, 8)
		if err != nil {
			log.Println(err)
			continue
		}

		do(byte(code))
	}
}
