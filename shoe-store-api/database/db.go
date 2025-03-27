package database

import (
	"fmt"
	"log"
	"shoe-store-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
}

func RunMigrations() {
	// Автомиграции GORM (альтернатива SQL-миграциям)
	// DB.AutoMigrate(&models.Product{})

	// Или выполнение SQL-миграций из файлов
	// Можно использовать библиотеку golang-migrate/migrate
}
