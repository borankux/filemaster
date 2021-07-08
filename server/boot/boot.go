package boot

import (
	DB "github.com/borankux/filemaster/server/db"
	"github.com/borankux/filemaster/server/models"
	"log"
)

func Migrate(){
	db := DB.GetDB()
	err := db.AutoMigrate(&models.Request{})
	if err != nil {
		log.Printf("Failed to auto migrate: %v", err)
	}
}
