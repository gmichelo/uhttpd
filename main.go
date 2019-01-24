package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

	//Check if user wants an HTTP or HTTPS server
	if *cert != "" || *prvKey != "" {
		//TODO: sanitize path to certificate and private key
		startHTTPSServer(*addr, *port, rootPath, *cert, *prvKey)
	} else {
		startHTTPServer(*addr, *port, rootPath)
	}
	//TODO: graceful shutdown, any resource cleanup?

}

func deriveWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
