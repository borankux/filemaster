package db

import (
	"github.com/borankux/filemaster/server/models"
	"log"
)

func SetupData(){
	err := db.AutoMigrate(&models.Request{})
	if err != nil {
		log.Printf("Failed to auto migrate: %v", err)
	}
}
