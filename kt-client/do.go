package main

import (
	"log"

	"github.com/remexre/kill-trigger"
)

func do(b byte) {
	switch b {
	case kt.HelloWorld.ID:
		log.Println("Hello, world!")
	case kt.KillJavaw.ID:
		killJavaw()
	case kt.Ping.ID:
		ping()
	case kt.Pong.ID:
		log.Println("PONG")
	default:
		log.Printf("Unknown command: %d", b)
	}
}
