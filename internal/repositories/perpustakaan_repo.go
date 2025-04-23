package repositories

import (
	"context"
	"github.com/Zauro25/perpus-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type PerpustakaanRepository struct {
	db *sqlx.DB
}

func (r *PerpustakaanRepository) Create(ctx context.Context, p *models.Perpustakaan) error{
	query := `insert into perpustakaan (id, nama, alamat, koleksi_buku, status) values ($1. $2. $3. $4. $5)`
	_, err := r.db.ExecContext(ctx, query, p.ID, p.Nama, p.Alamat, p.KoleksiBuku, p.Status)
	return err
}