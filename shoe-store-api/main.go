// main.go
package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"shoe-store-api/config"
	"shoe-store-api/database"
	"shoe-store-api/handlers"
	"time"
)

func main() {
	// Инициализация базы данных
	database.InitDB()
	database.RunMigrations()

	// Создание роутера Gin
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Публичные маршруты
	api := r.Group("/api")
	{
		api.GET("/products", handlers.GetProducts)
		api.POST("/orders", handlers.CreateOrder)
		api.POST("/login", handlers.Login)
	}

	// Защищенные маршруты
	admin := r.Group("/admin")
	admin.Use(handlers.AuthMiddleware())
	{
		admin.POST("/products", handlers.CreateProduct)
		admin.PUT("/products/:id", handlers.UpdateProduct)
		admin.DELETE("/products/:id", handlers.DeleteProduct)
	}

	// Запуск сервера
	port := ":" + config.LoadConfig().Port
	log.Printf("Server running on port %s", port)
	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
