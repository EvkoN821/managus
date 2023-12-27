package interfaces

import "github.com/IlyaZayats/managus/internal/entity"

type InvoiceRepository interface {
	GetInvoices() ([]entity.Invoice, error)
	UpdateInvoice(invoice entity.Invoice) error
	InsertInvoice(invoice entity.Invoice) error
	DeleteInvoice(id int) error
}
