package main

import (
	"github.com/ormushq/ormus/config"
	"github.com/ormushq/ormus/source/delivery/httpserver"
)

func main() {
	cfg := config.C()

	server := httpserver.New(cfg.Source)

	server.Serve()
}