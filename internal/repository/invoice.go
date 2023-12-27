package repository

import (
	"context"
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresInvoiceRepository struct {
	db *pgxpool.Pool
}

func NewPostgresInvoiceRepository(db *pgxpool.Pool) (interfaces.InvoiceRepository, error) {
	return &PostgresInvoiceRepository{
		db: db,
	}, nil
}

func (r *PostgresInvoiceRepository) GetInvoices() ([]entity.Invoice, error) {
	var invoices []entity.Invoice
	q := "SELECT id, client_id, invoice_id, cont_date, exec_date, sum_total, handed, accepted, add_info, basis_doc FROM Invoices"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return invoices, err
	}
	//foods, err =
	return r.parseRowsToSlice(rows)

}

func (r *PostgresInvoiceRepository) InsertInvoice(invoice entity.Invoice) error {
	q := "INSERT INTO Invoices (client_id, invoice_id, cont_date, exec_date, sum_total, handed, accepted, add_info, basis_doc) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	if _, err := r.db.Exec(context.Background(), q, invoice.ClientId, invoice.InvoiceId, invoice.ContDate, invoice.ExecDate, invoice.SumTotal, invoice.Handed, invoice.Accepted, invoice.AddInfo, invoice.BasisDoc); err != nil {
		return err
	}
	return nil
}

func (r *PostgresInvoiceRepository) UpdateInvoice(invoice entity.Invoice) error {
	q := "UPDATE Invoices SET (invoice_id, cont_date, exec_date, sum_total, handed, accepted, add_info, basis_doc) = ($1, $2, $3, $4, $5, $6, $7, $8) WHERE id=$9"
	if _, err := r.db.Exec(context.Background(), q, invoice.InvoiceId, invoice.ContDate, invoice.ExecDate, invoice.SumTotal, invoice.Handed, invoice.Accepted, invoice.AddInfo, invoice.BasisDoc, invoice.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresInvoiceRepository) DeleteInvoice(id int) error {
	q := "DELETE FROM Invoices WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresInvoiceRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Invoice, error) {
	var slice []entity.Invoice
	defer rows.Close()
	for rows.Next() {
		var id, clientId, sumTotal int
		var invoiceId, handed, accepted, addInfo, basisDoc string
		var contDate, execDate pgtype.Date
		if err := rows.Scan(&id, &clientId, &invoiceId, &contDate, &execDate, &sumTotal, &handed, &accepted, &addInfo, &basisDoc); err != nil {
			return slice, err
		}
		contDateString := contDate.Time.Format("2006.01.02")
		execDateString := execDate.Time.Format("2006.01.02")
		slice = append(slice, entity.Invoice{
			Id:        id,
			ClientId:  clientId,
			InvoiceId: invoiceId,
			ContDate:  contDateString,
			ExecDate:  execDateString,
			SumTotal:  sumTotal,
			Handed:    handed,
			Accepted:  accepted,
			AddInfo:   addInfo,
			BasisDoc:  basisDoc,
		})
	}
	return slice, nil
}
