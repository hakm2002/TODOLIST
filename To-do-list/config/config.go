package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/hakm2002/TODOLIST/models"
)

var JwtSecret = []byte("8Eb3yGaQzCz4E38Hw4QvwasViZKlnnO6d8oMM+kNMqQ=") //secret key
var db *gorm.DB

func InitDB() {
	// username pass 
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "")
	dbHost := getEnv("DB_HOST", "127.0.0.1")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "todolist")
	dsn := "gormuser:gormpass@tcp(192.168.68.75:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}
	if err := db.AutoMigrate(&models.Memo{}); err != nil {
		log.Fatal("Gagal koneksi database:", err)
	}
	log.Println("db tekoneksi via .env!")
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Gagal koneksi database")
	}
	return db
}
