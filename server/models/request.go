package models

import (
	DB "github.com/borankux/filemaster/server/db"
	"log"
)

type Request struct {
	Id int `gorm:"primaryKey"`
	Ip string
}

func(r Request) Save() *Request {
	db := DB.GetDB()

	result := db.Create(&r)

	if result.Error != nil {
		log.Fatalf("Failed to craete user: %v", result.Error)
	}

	if result.RowsAffected == 1 {
		log.Printf("Created user: %v", result)
	}

	return &r
}