package main

import (
	"log"
)

var commandActions = map[string]func(map[string]string) error{
	"keepalive": nil,
	"agora":     agoraHandler,
	"kill":      kill,
	"ping":      ping,
	"pong":      nil,
}

func do(m map[string]string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Caught panic in do:", r)
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()

	name, ok := m["name"]
	if !ok {
		log.Panic("Invalid command", m)
	}

	f, ok := commandActions[name]
	if !ok {
		log.Panic("Unknown command", name)
	}

	if f == nil {
		return nil
	}
	return f(m)
}
