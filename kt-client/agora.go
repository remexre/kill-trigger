package main

import (
	"log"

	"github.com/remexre/kill-trigger/agora"
)

func agoraHandler(m map[string]string) error {
	ret, err := agora.Run(m["code"])
	if err != nil {
		return err
	}

	log.Println("agora returned", ret)
	return nil
}
