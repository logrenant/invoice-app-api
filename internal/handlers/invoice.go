package handlers

import (
	"net/http"
	"strconv"

	"invoice-api/internal/models"
	"invoice-api/internal/services"

	"github.com/gin-gonic/gin"
)

type InvoiceHandler struct {
	service *services.InvoiceService
}

func NewInvoiceHandler(service *services.InvoiceService) *InvoiceHandler {
	return &InvoiceHandler{service: service}
}

// GET /invoices
func (h *InvoiceHandler) GetInvoices(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	search := c.Query("search")

	invoices, totalCount, err := h.service.GetInvoices(page, pageSize, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"invoices":    invoices,
		"total_count": totalCount,
	})
}

// GET /invoices/:id
func (h *InvoiceHandler) GetInvoice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	invoice, err := h.service.GetInvoiceByID(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else if invoice == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
	} else {
		c.JSON(http.StatusOK, invoice)
	}
}

// POST /invoices
func (h *InvoiceHandler) CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}

// PUT /invoices/:id
func (h *InvoiceHandler) UpdateInvoice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var invoice models.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invoice.ID = uint(id)
	if err := h.service.UpdateInvoice(&invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

// DELETE /invoices/:id
func (h *InvoiceHandler) DeleteInvoice(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.service.DeleteInvoice(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
