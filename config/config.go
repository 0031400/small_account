package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Addr string `json:"addr"`
}

var C Config

func Init() {
	b, err := os.ReadFile("config.json")
	if err != nil {
		log.Panicln(err)
	}
	err = json.Unmarshal(b, &C)
	if err != nil {
		log.Panicln(err)
	}
}
