package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/deFarro/fastpoke_backend/app/config"
	"github.com/deFarro/fastpoke_backend/app/middleware"
	"github.com/deFarro/fastpoke_backend/app/router"
)

func main() {
	config, err := config.GetConfig("config.yml")
	if err != nil {
		log.Fatalf("error while reading config file: %s\n", err)
	}

	router, err := router.NewRouter(config)
	if err != nil {
		log.Fatalf("error while creating router: %s\n", err)
	}

	fmt.Printf("Up and running on localhost:%s\n", config.AppPort)

	http.Handle("/version", middleware.Adapt(
		http.HandlerFunc(router.HandleVersion),
		middleware.WithHeaders,
	))

	http.ListenAndServe(":"+config.AppPort, nil)
}
