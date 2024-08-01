package handlers

import (
	"fmt"
	"gobotserver/db"
	"gobotserver/lookup"
	"gobotserver/models"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
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

	// Return a resp to the client
	fmt.Fprintln(w, "HTTP 200/OK")

	// Lookup geography of IP address and add to DB
	/* For local testing, spoof the calling IP address */
	spoofedIP := fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
	log.Printf("Spoofed IP: %s\n", spoofedIP)
	requestInfo.IPAddress = spoofedIP
	dbRecord := requestInfo.ConvertToDBRecord()
	dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(spoofedIP)
	/* What it should be: */
	// dbRecord.CountryCode, dbRecord.City = lookup.LookupIPAddress(requestInfo.IPAddress)

	client := db.NewClient()
	_, err = client.StoreRequest(dbRecord)
	if err != nil {
		log.Printf("Error inserting DB records: %v\n", err)
	}
}
