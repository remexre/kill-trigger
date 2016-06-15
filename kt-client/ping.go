package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func ping(m map[string]string) error {
	log.Println(">>> got ping")

	res, err := http.Post(fmt.Sprintf("%s/api/send", origin),
		"application/json",
		strings.NewReader(`{"name":"pong"}`))
	if err != nil {
		return err
	} else if err = res.Body.Close(); err != nil {
		return err
	}
	log.Println(">>> sent pong")
	return nil
}
