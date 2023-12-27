package entity

type Invoice struct {
	Id        int    `db:"id"`
	ClientId  int    `db:"client_id"`
	InvoiceId string `db:"invoice_id"`
	ContDate  string `db:"cont_date"`
	ExecDate  string `db:"exec_date"`
	SumTotal  int    `db:"sum_total"`
	Handed    string `db:"handed"`
	Accepted  string `db:"accepted"`
	AddInfo   string `db:"add_info"`
	BasisDoc  string `db:"basis_doc"`
}
