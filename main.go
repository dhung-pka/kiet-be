package main

import (
	"log"
	"net/http"
	"service/config"
	"service/router"
	"time"
)

// @title           Service API
// @version         1.0
// @description     Service API

// @host            localhost:18888
// @BasePath        /api/v1
func main() {

	server := http.Server{
		Addr:         ":" + config.GetAppPort(),
		ReadTimeout:  6000 * time.Second,
		WriteTimeout: 6000 * time.Second,
		Handler:      router.Router(),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println("Error run server: ", err)
	}
}
