package main

import (
	"encoding/json"
	"log"

	"golang.org/x/net/websocket"
)

const (
	origin = "https://kill-trigger.herokuapp.com"
	wsURL  = "wss://kill-trigger.herokuapp.com/socket"
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
	defer ws.Close()
	dec := json.NewDecoder(ws)

	for {
		var m map[string]string
		err := dec.Decode(&m)
		if err != nil {
			return err
		}

		err = do(m)
		if err != nil {
			return err
		}
	}
}
