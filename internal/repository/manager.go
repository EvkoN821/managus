package repository

import (
	"context"
	"github.com/IlyaZayats/managus/internal/entity"
	"github.com/IlyaZayats/managus/internal/interfaces"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresManagerRepository struct {
	db *pgxpool.Pool
}

func NewPostgresManagerRepository(db *pgxpool.Pool) (interfaces.ManagerRepository, error) {
	return &PostgresManagerRepository{
		db: db,
	}, nil
}

func (r *PostgresManagerRepository) GetManagers() ([]entity.Manager, error) {
	var managers []entity.Manager
	q := "SELECT id, name FROM Managers"
	rows, err := r.db.Query(context.Background(), q)
	if err != nil && err.Error() != "no rows in result set" {
		return managers, err
	}
	//faculties, err =
	return r.parseRowsToSlice(rows)

}

func (r *PostgresManagerRepository) InsertManager(manager entity.Manager) error {
	q := "INSERT INTO Managers (name) VALUES ($1)"
	if _, err := r.db.Exec(context.Background(), q, manager.Name); err != nil {
		return err
	}
	return nil
}

func (r *PostgresManagerRepository) UpdateManager(manager entity.Manager) error {
	q := "UPDATE Managers SET name=$1 WHERE id=$2"
	if _, err := r.db.Exec(context.Background(), q, manager.Name, manager.Id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresManagerRepository) DeleteManager(id int) error {
	q := "DELETE FROM Managers WHERE id=$1"
	if _, err := r.db.Exec(context.Background(), q, id); err != nil {
		return err
	}
	return nil
}

func (r *PostgresManagerRepository) parseRowsToSlice(rows pgx.Rows) ([]entity.Manager, error) {
	var slice []entity.Manager
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return slice, err
		}
		slice = append(slice, entity.Manager{Id: id, Name: name})
	}
	return slice, nil
}
