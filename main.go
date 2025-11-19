package main

import (
	"log"
	"net/http"
	"small_account/config"
	"small_account/logger"
	"small_account/model"
	"small_account/router"
)

func main() {
	logger.Init()
	config.Init()
	model.Init()
	r := router.SetupRouter()
	log.Printf("The server is listening on %s", config.C.Addr)
	err := http.ListenAndServe(config.C.Addr, r)
	if err != nil {
		log.Panicln(err)
	}
}
