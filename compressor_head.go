package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/image/", imageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	server := &http.Server{
		Addr:              net.JoinHostPort("", port),
		ReadHeaderTimeout: time.Second * 10,
		ReadTimeout:       time.Second * 10,
		WriteTimeout:      time.Second * 10,
		IdleTimeout:       time.Second * 10,
	}

	log.Printf("Open http://localhost:%s in the browser", port)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("Error starting server - %v", err)
	}
}
