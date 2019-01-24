package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

var usageMessage = `Usage: uhttpd [flags] [path]
If no path is specified, it picks automatically the one
in which the command was invoked (working directory).
When both certificate and private key files are specified
uhttpd will start HTTPS server instead of (unsecured) HTTP.
Flags:
 -l=Listen port
 -a=Listen IP address
 -c=Path to certificate file
 -p=Path to private key file
 `

var (
	port   = flag.String("l", "8080", "Listen port")
	addr   = flag.String("a", "localhost", "Listen IP address")
	cert   = flag.String("c", "", "Certificate file")
	prvKey = flag.String("p", "", "Private key file")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usageMessage)
		os.Exit(2)
	}
	flag.Parse()

	var rootPath string
	switch flag.NArg() {
	case 1:
		rootPath = flag.Arg(0)
	default:
		rootPath = deriveWorkingDir()
	}

	//Create new server
	s := NewSimpleServer(*addr, *port, rootPath)

	//Register function that calls s.Shutdown once
	//the program receives interrupt signal
	wg := registerCleanupFunction(s)

	//Check if user wants an HTTP or HTTPS server and start
	if *cert != "" || *prvKey != "" {
		//TODO: sanitize path to certificate and private key
		s.StartHTTPSServer(*cert, *prvKey)
	} else {
		s.StartHTTPServer()
	}
	//Wait for cleanup function to finish
	wg.Wait()
	log.Println("Exited")
}

func registerCleanupFunction(s SimpleServer) *sync.WaitGroup {
	wg := new(sync.WaitGroup)
	go func() {
		wg.Add(1)
		defer wg.Done()
		//Register handler for interrupt signal
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		//Wait for interr signal
		<-sigint

		// We received an interrupt signal, shut down.
		s.Shutdown()
	}()
	return wg
}

func deriveWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
