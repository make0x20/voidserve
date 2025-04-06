package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	acceptAny := flag.Bool("a", false, "Accept any IP - other devices will be able to connect")
	port := flag.String("p", "8080", "Port to serve on")
	flag.Parse()

	// Get current dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// Create a file server
	fs := http.FileServer(http.Dir(dir))

	// Set handler
	http.Handle("/", fs)

	// Set listen address
	addr := ":" + string(*port)
	if !*acceptAny {
		addr = "127.0.0.1" + addr
	}

	// Start the server
	log.Println("Serving files on", addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
