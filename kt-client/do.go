package main

import (
	"log"

	"github.com/remexre/kill-trigger"
)

func doNothing() error { return nil }

var commandActions = map[byte]func() error{
	kt.KeepAlive.ID:  doNothing,
	kt.HelloWorld.ID: helloWorld,
	kt.KillJava.ID:   killJava,
	kt.Ping.ID:       ping,
	kt.Pong.ID:       doNothing,
}

func do(b byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Caught panic in do:", r)
			if e, ok := r.(error); ok {
				err = e
			}
		}
	}()

	f, ok := commandActions[b]
	if !ok {
		log.Panicf("Unknown command: %d", b)
	}
	return f()
}
