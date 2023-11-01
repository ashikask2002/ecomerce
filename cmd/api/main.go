package main

import (
	"log"

	"github.com/ashikask2002/ecomerce.git/pkg/config"
	"github.com/ashikask2002/ecomerce.git/pkg/di"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	server, err := di.InitializeAPI(config)
	if err != nil {
		log.Fatal(err)
	} else {
		server.Start()
	}
}
