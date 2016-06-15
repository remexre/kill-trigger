package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func ping(m map[string]string) error {
	log.Println(">>> got ping")

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(map[string]string{
		"name": "pong",
		"host": hostname,
	})

	res, err := http.Post(fmt.Sprintf("%s/api/send", origin), "application/json", buf)
	if err != nil {
		return err
	} else if err = res.Body.Close(); err != nil {
		return err
	}
	log.Println(">>> sent pong")
	return nil
}
