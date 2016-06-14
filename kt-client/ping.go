package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping() error {
	log.Println(">>> got ping")

	res, err := http.Post(fmt.Sprintf("%s/api/255/send", origin), "", nil)
	if err != nil {
		return err
	} else if err = res.Body.Close(); err != nil {
		return err
	}
	log.Println(">>> sent pong")
	return nil
}
