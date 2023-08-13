package main

import (
	"bookz/api/server"
	"bookz/config"
)

func main() {
	config.LoadDotEnv()

	s := server.Server{}

	s.Create()
}
