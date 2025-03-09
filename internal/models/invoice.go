package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ServiceName   string    `gorm:"not null"`
	InvoiceNumber int       `gorm:"unique;not null"`
	Date          time.Time `gorm:"type:date;not null"`
	Amount        float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"check:status IN ('Paid', 'Pending', 'Unpaid')"`
}
