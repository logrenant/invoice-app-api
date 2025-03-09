package main

import (
	"log"
	"math/rand"
	"time"

	"invoice-api/internal/config"
	"invoice-api/internal/db"
	"invoice-api/internal/models"
	"invoice-api/internal/repositories"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
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

	repo := repositories.NewInvoiceRepository(database)

	if err := clearExistingData(database); err != nil {
		log.Fatal("Veriler temizlenemedi:", err)
	}

	if err := seedInvoices(repo, 100); err != nil {
		log.Fatal("Seed işlemi başarısız:", err)
	}

	log.Println("100 adet test verisi başarıyla oluşturuldu!")
}

func clearExistingData(db *gorm.DB) error {
	return db.Exec("DELETE FROM invoices").Error
}

func seedInvoices(repo *repositories.InvoiceRepository, count int) error {
	rand.Seed(time.Now().UnixNano())

	services := []string{"DMP", "SSP", "Ad Server", "Analytics"}
	statuses := []string{"Paid", "Pending", "Unpaid"}

	tx := repo.GetDB().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for i := 0; i < count; i++ {
		invoice := models.Invoice{
			ServiceName:   services[rand.Intn(len(services))],
			InvoiceNumber: 1000 + i + 1,
			Date:          time.Now().AddDate(0, 0, -rand.Intn(365)),
			Amount:        rand.Float64() * 10000,
			Status:        statuses[rand.Intn(len(statuses))],
		}

		if err := tx.Create(&invoice).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
