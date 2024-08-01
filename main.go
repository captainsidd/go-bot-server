package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"gobotserver/db"
	"gobotserver/lookup"
	"gobotserver/models"
)

var commonPorts = []int{80, 443, 21, 22, 25, 110, 143, 993, 995, 587}

func main() {
	http.HandleFunc("/", handler)
	for _, port := range commonPorts {
		addr := fmt.Sprintf(":%d", port)
		go startServer(port, addr)
		// go listenOnPort(port)
	}
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
		fmt.Printf("Error parsing remote address %s: %v\n", remoteAddr, err)
		return
	}
	fmt.Printf("Connection from IP %s on Port %s at %s\n", ip, port, timestamp)

	requestInfo := models.RequestInfo{
		IPAddress: ip,
		Timestamp: timestamp,
		Port:      port,
		Path:      r.URL.Path,
	}

	/* For local testing, spoof the calling IP address */
	spoofedIP := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	fmt.Printf("Spoofed IP: %s\n", spoofedIP)
	requestInfo.IPAddress = spoofedIP
	fmt.Println(requestInfo)
	dbRecord := requestInfo.ConvertToDBRecord()
	dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(spoofedIP)

	/* What it should be: */
	// Lookup geography of IP address and add to DB
	// dbRecord := requestInfo.ConvertToDBRecord()
	// dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(requestInfo.IPAddress)

	addRequestToDB(dbRecord)
}

func listenOnPort(port int) {
	address := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("Error listening on port %d: %v\n", port, err)
		return
	}
	defer listener.Close()
	fmt.Printf("Listening on port %d\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection on port %d: %v\n", port, err)
			continue
		}
		go handleConnection(conn, port)
	}
}

func handleConnection(conn net.Conn, port int) {
	// Find the IP & Port of the request
	timestamp := time.Now().UTC()
	remoteAddr := conn.RemoteAddr().String()
	ip, _, err := net.SplitHostPort(remoteAddr)
	if err != nil {
		fmt.Printf("Error parsing remote address %s: %v\n", remoteAddr, err)
		return
	}
	fmt.Printf("Connection from IP %s on Port %d at %s\n", ip, port, timestamp)

	requestInfo := models.RequestInfo{
		IPAddress: ip,
		Timestamp: timestamp,
		Port:      strconv.Itoa(port),
	}
	// Close the connection
	conn.Close()

	// Lookup geography of IP address and add to DB
	dbRecord := requestInfo.ConvertToDBRecord()
	dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(requestInfo.IPAddress)
	addRequestToDB(dbRecord)
}

func addRequestToDB(request *models.DBRecord) {
	client := db.NewClient()
	client.StoreRequest(request)
}
