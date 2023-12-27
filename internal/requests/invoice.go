package requests

type InsertInvoiceRequest struct {
	ClientId  int    `json:"client_id" binding:"required"`
	InvoiceId string `json:"invoice_id" binding:"required"`
	ContDate  string `json:"cont_date" binding:"required"`
	ExecDate  string `json:"exec_date" binding:"required"`
	SumTotal  int    `json:"sum_total" binding:"required"`
	Handed    string `json:"handed" binding:"required"`
	Accepted  string `json:"accepted" binding:"required"`
	AddInfo   string `json:"add_info" binding:"required"`
	BasisDoc  string `json:"basis_doc" binding:"required"`
}

type UpdateInvoiceRequest struct {
	Id        int    `json:"id" binding:"required"`
	InvoiceId string `json:"invoice_id" binding:"required"`
	ContDate  string `json:"cont_date" binding:"required"`
	ExecDate  string `json:"exec_date" binding:"required"`
	SumTotal  int    `json:"sum_total" binding:"required"`
	Handed    string `json:"handed" binding:"required"`
	Accepted  string `json:"accepted" binding:"required"`
	AddInfo   string `json:"add_info" binding:"required"`
	BasisDoc  string `json:"basis_doc" binding:"required"`
}

type DeleteInvoiceRequest struct {
	Id int `json:"id" binding:"required"`
}
