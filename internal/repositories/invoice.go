package repositories

import (
	"errors"
	"invoice-api/internal/models"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func (r *InvoiceRepository) GetDB() *gorm.DB {
	return r.db
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{db: db}
}

func (r *InvoiceRepository) GetAll(page, pageSize int, search string) ([]models.Invoice, int64, error) {
	var invoices []models.Invoice
	var totalCount int64

	query := r.db.Model(&models.Invoice{})

	if search != "" {
		query = query.Where("service_name ILIKE ? OR invoice_number::text LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	query.Count(&totalCount)

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&invoices).Error

	return invoices, totalCount, err
}

func (r *InvoiceRepository) GetByID(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	err := r.db.First(&invoice, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &invoice, err
}

func (r *InvoiceRepository) Create(invoice *models.Invoice) error {
	return r.db.Create(invoice).Error
}

func (r *InvoiceRepository) Update(invoice *models.Invoice) error {
	return r.db.Save(invoice).Error
}

func (r *InvoiceRepository) Delete(id uint) error {
	return r.db.Delete(&models.Invoice{}, id).Error
}
