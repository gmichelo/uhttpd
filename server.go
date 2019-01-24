package main

import (
	"log"
	"net/http"
	"time"
)

func createCustomServer(address, port string) *http.Server {
	//Create server with custom values
	s := http.Server{
		Addr:              address + ":" + port,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}
	//TODO: create a custom http.Transport with customizable values
	//especially MaxConnections
	return &s
}

// registerHandles set the necessary handles into
// DefaultServeMux
func registerHandles(path string) {
	//Register handle to serve files
	http.Handle("/", http.FileServer(http.Dir(path)))
}

func startHTTPServer(address, port, path string) {
	//Create server with custom values
	s := createCustomServer(address, port)
	//Register FileServer handle
	registerHandles(path)
	//Start the HTTP server
	log.Println("uHTTP (unsecured) server listening on address:", address+":"+port, "with root path:", path)
	log.Fatal(s.ListenAndServe())
}

func startHTTPSServer(address, port, path, certFile, keyFile string) {
	//Create server with custom values
	s := createCustomServer(address, port)
	//Register FileServer handle
	registerHandles(path)
	//Start the HTTP server
	log.Println("uHTTP (secure) server listening on address:", address+":"+port, "with root path:", path)
	//TODO: find a better way to pass certificate and key
	//TODO: is there a way to force the browser to use https? It should exist, find it.
	log.Fatal(s.ListenAndServeTLS(certFile, keyFile))
}
