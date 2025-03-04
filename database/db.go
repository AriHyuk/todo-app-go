package database

import (
	"log"
	"go-todo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Ganti dengan informasi database MySQL yang sesuai
	dsn := "root:@tcp(127.0.0.1:3306)/todos?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrasi database untuk memastikan tabel `todos` dibuat
	DB.AutoMigrate(&models.Todo{})
}
