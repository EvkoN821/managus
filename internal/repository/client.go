package repository

import (
	"context"
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresClientRepository struct {
	db *pgxpool.Pool
}

func NewPostgresClientRepository(db *pgxpool.Pool) (interfaces.ClientRepository, error) {
	return &PostgresClientRepository{
		db: db,
	}, nil
}

func (r *PostgresClientRepository) GetClients() ([]entity.Client, error) {
	var clients []entity.Client
	q := "SELECT id, manager_id, name FROM Clients"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return clients, err
	}
	return r.parseRowsToSlice(rows)

}

func (r *PostgresClientRepository) InsertClient(client entity.Client) error {
	q := "INSERT INTO Clients (manager_id, name) VALUES ($1, $2)"
	if _, err := r.db.Exec(context.Background(), q, client.ManagerId, client.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresClientRepository) UpdateClient(client entity.Client) error {
	q := "UPDATE Clients SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, client.Name, client.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresClientRepository) DeleteClient(id int) error {
	q := "DELETE FROM Clients WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresClientRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Client, error) {
	var slice []entity.Client
	defer rows.Close()
	for rows.Next() {
		var id, managerId int
		var name string
		if err := rows.Scan(&id, &managerId, &name); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Client{Id: id, ManagerId: managerId, Name: name})
	}
	return slice, nil
}
