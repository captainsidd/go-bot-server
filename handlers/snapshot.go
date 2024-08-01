package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gobotserver/db"
)

func SnapshotHandler(w http.ResponseWriter, r *http.Request) {
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
