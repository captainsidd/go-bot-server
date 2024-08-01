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
