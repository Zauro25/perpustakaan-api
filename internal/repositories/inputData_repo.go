package repositories

import(
	"context"
	"github.com/Zauro25/perpus-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type InputDataRepository struct {
	db *sqlx.DB
}

func (r *InputDataRepository) Create(ctx context.Context, p *models.DataPerpustakaan) error{
	query := `insert into inputdata (id, namaPerpustakaan, alamat, jenisPerpustakaan, nomorInduk, jumlahCetak, jumlahEbook, jumlahSDM, jumlahPengunjung, jumlahAnggota) values ($1. $2. $3. $4. $5)`
	_, err := r.db.ExecContext(ctx, query, p.ID, p.NamaPerpustakaan, p.Alamat, p.JenisPerpustakaan, p.NomorInduk, p.JumlahCetak, p.JumlahEbook, p.JumlahSDM, p.JumlahPengunjung, p.JumlahAnggota)
	return err
}