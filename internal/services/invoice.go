package services

import (
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"strconv"
)

type InvoiceService struct {
	repo interfaces.InvoiceRepository
}

func NewInvoiceService(repo interfaces.InvoiceRepository) (*InvoiceService, error) {
	return &InvoiceService{
		repo: repo,
	}, nil
}

func (s *InvoiceService) GetInvoices() ([]map[string]string, error) {
	invoices, err := s.repo.GetInvoices()
	if err != nil {
		return nil, err
	}
	invoicesSlice := []map[string]string{}
	for _, item := range invoices {
		foodsMap := map[string]string{
			"id":         strconv.Itoa(item.Id),
			"client_id":  strconv.Itoa(item.ClientId),
			"invoice_id": item.InvoiceId,
			"cont_date":  item.ContDate,
			"exec_date":  item.ExecDate,
			"sum_total":  strconv.Itoa(item.SumTotal),
			"handed":     item.Handed,
			"accepted":   item.Accepted,
			"add_info":   item.AddInfo,
			"basis_doc":  item.BasisDoc,
		}
		invoicesSlice = append(invoicesSlice, foodsMap)
	}
	return invoicesSlice, nil
}

func (s *InvoiceService) InsertInvoice(invoiceId, contDate, execDate, handed, accepted, addInfo, basisDoc string, clientId, sumTotal int) error {
	return s.repo.InsertInvoice(entity.Invoice{Id: 0, ClientId: clientId, InvoiceId: invoiceId, ContDate: contDate, ExecDate: execDate, SumTotal: sumTotal, Handed: handed, AddInfo: addInfo, Accepted: accepted, BasisDoc: basisDoc})
}

func (s *InvoiceService) UpdateInvoice(invoiceId, contDate, execDate, handed, accepted, addInfo, basisDoc string, id, sumTotal int) error {
	return s.repo.UpdateInvoice(entity.Invoice{Id: id, ClientId: 0, InvoiceId: invoiceId, ContDate: contDate, ExecDate: execDate, SumTotal: sumTotal, Handed: handed, AddInfo: addInfo, Accepted: accepted, BasisDoc: basisDoc})
}

func (s *InvoiceService) DeleteInvoice(id int) error {
	return s.repo.DeleteInvoice(id)
}
