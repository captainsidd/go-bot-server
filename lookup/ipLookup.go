package lookup

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gobotserver/models"
)

func LookupIPAddress(ipAddress string) (countryCode string, city string) {
	url := fmt.Sprintf("http://ip-api.com/json/%s", ipAddress)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Failed to get location data:", err)
		return "", ""
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read response body:", err)
		return "", ""
	}

	var locationInfo models.Location
	err = json.Unmarshal(body, &locationInfo)
	if err != nil {
		log.Println("Failed to unmarshal location data:", err)
		return "", ""
	}
	return locationInfo.Country, locationInfo.City
}

// Lookup geolocation

// Print request details to console
// fmt.Printf("IP: %s\nTime: %s\nUser-Agent: %s\nURL Path: %s\nGeolocation: %+v\n\n", ip, timestamp, userAgent, urlPath)
// }

// // GeolocationData represents the geolocation data structure.

// func getGeolocation(ip string) GeolocationData {
// 	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Println("Failed to get geolocation data:", err)
// 		return GeolocationData{}
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Println("Failed to read response body:", err)
// 		return GeolocationData{}
// 	}

// 	var geoData GeolocationData
// 	err = json.Unmarshal(body, &geoData)
// 	if err != nil {
// 		log.Println("Failed to unmarshal geolocation data:", err)
// 		return GeolocationData{}
// 	}
// 	return geoData
// }
