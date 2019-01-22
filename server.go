package main

import (
	"log"
	"net/http"
	"time"
)

func startServer(address, port, path string) {
	//Create server with custom values
	s := &http.Server{
		Addr:              address + ":" + port,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	//Register handle to serve files
	http.Handle("/", http.FileServer(http.Dir(path)))
	//Start the HTTP server
	log.Println("uHTTP server listening on address:", address+":"+port, "with root path:", path)
	log.Fatal(s.ListenAndServe())

	//TODO: add support for HTTPS
	//TODO: create a custom http.Transport with customizable values
	//especially MaxConnections
}
