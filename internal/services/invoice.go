package services

import (
	"invoice-api/internal/models"
	"invoice-api/internal/repositories"
)

type InvoiceService struct {
	repo *repositories.InvoiceRepository
}

func NewInvoiceService(repo *repositories.InvoiceRepository) *InvoiceService {
	return &InvoiceService{repo: repo}
}

func (s *InvoiceService) GetInvoices(page int, pageSize int, search string) ([]models.Invoice, int64, error) {
	return s.repo.GetAll(page, pageSize, search)
}

func (s *InvoiceService) GetInvoiceByID(id uint) (*models.Invoice, error) {
	return s.repo.GetByID(id)
}

func (s *InvoiceService) CreateInvoice(invoice *models.Invoice) error {
	return s.repo.Create(invoice)
}

func (s *InvoiceService) UpdateInvoice(invoice *models.Invoice) error {
	return s.repo.Update(invoice)
}

func (s *InvoiceService) DeleteInvoice(id uint) error {
	return s.repo.Delete(id)
}
