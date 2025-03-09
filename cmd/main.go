package main

import (
	"log"

	"invoice-api/internal/config"
	"invoice-api/internal/db"
	"invoice-api/internal/handlers"
	"invoice-api/internal/models"
	"invoice-api/internal/repositories"
	"invoice-api/internal/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Config yüklenemedi:", err)
	}

	database, err := db.Connect(cfg)
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}

	if err := database.AutoMigrate(&models.Invoice{}); err != nil {
		log.Fatal("Migration hatası:", err)
	}

	invoiceRepo := repositories.NewInvoiceRepository(database)
	invoiceService := services.NewInvoiceService(invoiceRepo)
	invoiceHandler := handlers.NewInvoiceHandler(invoiceService)

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/invoices", invoiceHandler.GetInvoices)
	router.GET("/invoices/:id", invoiceHandler.GetInvoice)
	router.POST("/invoices", invoiceHandler.CreateInvoice)
	router.PUT("/invoices/:id", invoiceHandler.UpdateInvoice)
	router.DELETE("/invoices/:id", invoiceHandler.DeleteInvoice)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Sunucu başlatılamadı:", err)
	}
}
