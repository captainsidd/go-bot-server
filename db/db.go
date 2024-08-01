package db

import (
	"fmt"
	"gobotserver/models"
	"os"
	"time"

	supa "github.com/nedpals/supabase-go"
)

type DB struct {
	Client *supa.Client
}

func NewClient() (db *DB) {
	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)
	return &DB{
		Client: supabase,
	}
}

func (db *DB) StoreRequest(request *models.DBRecord) (*[]models.DBRecord, error) {
	request.CreatedAt = time.Now().UTC()
	var results []models.DBRecord
	err := db.Client.DB.From("requests").Insert(request).Execute(&results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &results, nil // inserted rows
}

func (db *DB) GetLastRecords(count int) (*[]models.DBRecord, error) {
	var results []models.DBRecord
	err := db.Client.DB.From("requests").Select("*").Limit(count).Execute(&results)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &results, nil // fetched rows
}
