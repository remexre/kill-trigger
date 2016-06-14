package main

import (
	"log"

	"github.com/remexre/kill-trigger"
)

func do(b byte) {
	switch b {
	case kt.KeepAlive.ID:
		// Do nothing on keepalive.
	case kt.HelloWorld.ID:
		log.Println("Hello, world!")
	case kt.KillJava.ID:
		killJava()
	case kt.Ping.ID:
		ping()
	case kt.Pong.ID:
		log.Println(">>> got pong")
	default:
		log.Printf("Unknown command: %d", b)
	}
}
