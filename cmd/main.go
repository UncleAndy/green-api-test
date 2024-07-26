package main

import (
	"github.com/UncleAndy/green-api-test/internal/config"
	"github.com/UncleAndy/green-api-test/internal/greenapi"
	"github.com/UncleAndy/green-api-test/internal/server"
)

func main() {
	cfg := config.GetConfig()

	client := greenapi.NewClient()

	srv, err := server.New(&cfg.Server, client)
	if err != nil {
		panic(err)
	}

	srv.Run()
}
