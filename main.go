package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var usageMessage = `Usage: uhttpd [-l] [-a] [path]
If no path is specified, it picks automatically the one
in which the command was invoked.
Flags:
 -l=Listen port
 -a=Listen IP address
 `

var (
	port = flag.String("l", "8080", "Listen port")
	addr = flag.String("a", "localhost", "Listen IP address")
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

	if *addr == "" || *port == "" {
		flag.Usage()
	}

	startServer(*addr, *port, rootPath)
}

func deriveWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
