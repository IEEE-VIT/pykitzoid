package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/IEEE-VIT/pykitzoid/algorithms/Linear_Regression/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server := &http.Server{
		Addr: ":" + port, Handler: api.NewHandler(),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("Pykitzoid API listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
