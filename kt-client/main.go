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
	for {
		log.Println(connect())
	}
}

func connect() error {
	ws, err := websocket.Dial(wsURL, "", origin)
	if err != nil {
		return err
	}

	for {
		var codeStr string
		err := websocket.Message.Receive(ws, &codeStr)
		if err != nil {
			return err
		}

		code, err := strconv.ParseUint(codeStr, 10, 8)
		if err != nil {
			log.Println(err)
			continue
		}

		err = do(byte(code))
		if err != nil {
			return err
		}
	}
}
