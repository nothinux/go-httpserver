package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	directory := flag.String("dir", ".", "directory location want to serve")
	port := flag.Int("port", 8080, "which port go-httpserver will be running")
	flag.Parse()

	if isExist(directory) {
		log.Fatal("directory is not exist")
	}

	runServe(*directory, *port)
}

func runServe(dir string, port int) {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      logRequest(mux),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fs := http.FileServer(http.Dir(dir))
	mux.Handle("/", http.StripPrefix("/", fs))

	log.Printf("server is running on port %d", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
