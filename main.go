package main

import (
	"fmt"
	"log"
	"net/http"

	"gobotserver/handlers"

	"github.com/joho/godotenv"
)

// This list of ports comes from SecLists. Using only some at the moment because of local dev.
// https://github.com/danielmiessler/SecLists/blob/master/Discovery/Infrastructure/common-http-ports.txt
var commonPorts = []int{
	66,
	80,
	81,
	443,
	445,
	457,
	1080,
	1241,
	1352,
	1433}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV vars, unable to start.")
	}

	http.HandleFunc("/", handlers.DefaultHandler)
	for _, port := range commonPorts {
		addr := fmt.Sprintf(":%d", port)
		go startServer(port, addr)
	}

	go func() {
		http.HandleFunc("/snapshot", handlers.SnapshotHandler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start /snapshot on port 8080: %v\n", err)
		}
		log.Println("DB Snapshots can be found at :8080/snapshot")
	}()
	// Prevent the main process from exiting immediately
	select {}
}

func startServer(port int, addr string) {
	log.Printf("Starting server on port %d\n", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start server on port %d: %v\n", port, err)
	}
}
