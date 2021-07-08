package db

import (
	"fmt"
	"github.com/borankux/filemaster/server/config"
	"github.com/borankux/filemaster/server/models"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"
var db *gorm.DB

func getDSN () string {
	conf := config.GetConf()
	user := conf.GetString("db.user")
	pass := conf.GetString("db.pass")
	host := conf.GetString("db.host")
	port := conf.GetInt32("db.port")
	name := conf.GetString("db.name")

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host , port, name)
}

func Init()  {
	log.Println(getDSN())
	db, _ = gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	err := db.AutoMigrate(&models.Request{})

	if err != nil {
		log.Printf("Failed to auto migrate: %v", err)
	}
}

func GetDB() *gorm.DB{
	return db
}