package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

// SimpleServer represents one instance of
// HTTP/HTTPS file server
type SimpleServer struct {
	srv *http.Server
}

// NewSimpleServer creates the handle for a
// new simple server listeting on <address> and <port>
// serving the directory pointed by <path>
func NewSimpleServer(address, port, path string) SimpleServer {
	//Create server with custom values
	s := http.Server{
		Addr:              fmt.Sprintf("%v:%v", address, port),
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	//Register handle to serve files
	http.Handle("/", http.FileServer(http.Dir(path)))

	//TODO: create a custom http.Transport with customizable values
	//especially MaxConnections
	log.Println("Created server on path", path, "under", s.Addr, "address")
	return SimpleServer{&s}
}

// Shutdown terminates the server instance by calling
// http.Shutdown. In case of errors, it immediately closes
// the server through http.Close
func (s SimpleServer) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() //Don't forget to call it anyway. It releases the context's resources
	err := s.srv.Shutdown(ctx)
	if err != nil {
		log.Println("Shutdown failed, trying to close the server...", err)
		err := s.srv.Close()
		log.Fatalln("Close failed, just exiting...", err)
	}
	log.Println("Server shutdown")
}

// StartHTTPServer starts the HTTP server.
// It blocks until Shutdown is called or
// an error is returned
func (s SimpleServer) StartHTTPServer() {
	//Start the HTTP server
	log.Println("uHTTP (unsecured) server listening")
	err := s.srv.ListenAndServe()
	if err != http.ErrServerClosed {
		log.Fatal("Server unexpectedly crashed:", err)
	}
}

// StartHTTPSServer starts the HTTPS server
// using certificate and private key.
// It blocks until Shutdown is called or
// an error is returned
func (s SimpleServer) StartHTTPSServer(certFile, keyFile string) {
	//Start the HTTPS server
	log.Println("uHTTP (secure) server listening")
	//TODO: find a better way to pass certificate and key
	//TODO: Write custom handle that forces redirect/redial with TLS: 301 Moved Permanently
	err := s.srv.ListenAndServeTLS(certFile, keyFile)
	if err != http.ErrServerClosed {
		log.Fatal("Server unexpectedly crashed:", err)
	}
}
