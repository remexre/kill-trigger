package main

import (
	"fmt"
	"log"
	"net/http"
)

func ping() {
	res, err := http.Post(fmt.Sprintf("%s/api/255/send", origin), "", nil)
	if err != nil {
		log.Println(err)
	} else if err = res.Body.Close(); err != nil {
		log.Println(err)
	}
}