package db

import (
	"fmt"
	"gobotserver/models"
	"time"

	supa "github.com/nedpals/supabase-go"
)

type DB struct {
	Client *supa.Client
}

func NewClient() (db *DB) {
	supabaseUrl := "https://svwwtzplqnbkpfzmambv.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InN2d3d0enBscW5ia3Bmem1hbWJ2Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MjI0NDA1NzcsImV4cCI6MjAzODAxNjU3N30.cTuKDqisefzdJaipIv4DOJN_-UgZXVx6kP-4xiApDCo"
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return &DB{
		Client: supabase,
	}
}

func (db *DB) StoreRequest(request *models.DBRecord) {
	request.CreatedAt = time.Now().UTC()
	var results []models.DBRecord
	err := db.Client.DB.From("requests").Insert(request).Execute(&results)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(results) // Inserted rows
}
