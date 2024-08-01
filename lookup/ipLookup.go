package lookup

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"gobotserver/models"
)

// We want to lookup the geolocation of an IP address, specifically the country code and city
// This isn't something that we need, but it's good to know if there's an associated location for
// a request.
// ipAddress must be IPv4
// Returns two strings for country code and city. Strings are empty string if no geolocation was found.
func LookupIPAddress(ipAddress string) (countryCode string, city string) {
	// Using the ip-api API. It's free and simple
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
	// Only interested in the city and the country, but marshalling into the Location struct for
	// future extendability.
	var locationInfo models.Location
	err = json.Unmarshal(body, &locationInfo)
	if err != nil {
		log.Println("Failed to unmarshal location data:", err)
		return "", ""
	}
	return locationInfo.Country, locationInfo.City
}
