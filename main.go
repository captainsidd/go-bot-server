package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"gobotserver/db"
	"gobotserver/lookup"
	"gobotserver/models"

	"github.com/joho/godotenv"
)

var commonPorts = []int{80, 443, 21, 22, 25, 110, 143, 993, 995, 587}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading ENV vars, unable to start.")
	}

	http.HandleFunc("/", handler)
	for _, port := range commonPorts {
		addr := fmt.Sprintf(":%d", port)
		go startServer(port, addr)
		// go listenOnPort(port)
	}

	go func() {
		http.HandleFunc("/snapshot", snapshotHandler)
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

func handler(w http.ResponseWriter, r *http.Request) {
	// Get the port this handler is handling by digging into the guts of the request
	// Ugly but can't find a way around it
	serverAddr := r.Context().Value(http.LocalAddrContextKey).(net.Addr)
	_, port, _ := net.SplitHostPort(serverAddr.String())

	// Fetch the things from the thing
	timestamp := time.Now().UTC()
	remoteAddr := r.RemoteAddr
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		log.Printf("Error parsing remote address %s: %v\n", remoteAddr, err)
		return
	}
	log.Printf("Connection from IP %s on Port %s at %s\n", ip, port, timestamp)

	requestInfo := models.RequestInfo{
		IPAddress: ip,
		Timestamp: timestamp,
		Port:      port,
		Path:      r.URL.Path,
	}

	fmt.Fprintln(w, "HTTP 200/OK")

	/* For local testing, spoof the calling IP address */
	spoofedIP := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	log.Printf("Spoofed IP: %s\n", spoofedIP)
	requestInfo.IPAddress = spoofedIP
	dbRecord := requestInfo.ConvertToDBRecord()
	dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(spoofedIP)

	/* What it should be: */
	// Lookup geography of IP address and add to DB
	// dbRecord := requestInfo.ConvertToDBRecord()
	// dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(requestInfo.IPAddress)

	addRequestToDB(dbRecord)
}

func snapshotHandler(w http.ResponseWriter, r *http.Request) {
	records, err := db.NewClient().GetLastRecords(25)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Error fetching DB records")
	}
	jsonRecords, err := json.Marshal(records)
	if err != nil {
		log.Println(err)
		fmt.Fprintln(w, "Error converting records to JSON")
	}
	fmt.Fprintln(w, string(jsonRecords))
}

func addRequestToDB(request *models.DBRecord) {
	client := db.NewClient()
	_, err := client.StoreRequest(request)
	if err != nil {
		log.Printf("Error inserting DB records: %v\n", err)
	}
}
