package main

import (
	"github.com/augustoroman/promise"
	"github.com/gopherjs/gopherjs/js"
	"github.com/remexre/kill-trigger/agora"
)

func main() {
	js.Global.Set("agora", promise.Promisify(agora.Run))
}
