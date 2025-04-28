package repositories

import(
	"context"
	"github.com/Zauro25/perpus-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type PerpustakaanRepository struct {
	db *sqlx.DB
}


func (r *PerpustakaanRepository) Create(ctx context.Context, p *models.DataPerpustakaan) error{
	query := `insert into inputdata (id, namaPerpustakaan, alamat, jenisPerpustakaan, nomorInduk, jumlahCetak, jumlahEbook, jumlahSDM, jumlahPengunjung, jumlahAnggota) values ($1. $2. $3. $4. $5)`
	_, err := r.db.ExecContext(ctx, query, p.ID, p.NamaPerpustakaan, p.Alamat, p.JenisPerpustakaan, p.NomorInduk, p.JumlahCetak, p.JumlahEbook, p.JumlahSDM, p.JumlahPengunjung, p.JumlahAnggota)
	return err
}

func (r *PerpustakaanRepository) GetAll(ctx context.Context) ([]models.DataPerpustakaan, error) {
	query := `SELECT * FROM inputdata`
	var data []models.DataPerpustakaan
	err := r.db.SelectContext(ctx, &data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *PerpustakaanRepository) GetByID(ctx context.Context, id string) (*models.DataPerpustakaan, error) {
	query := `SELECT * FROM inputdata WHERE id = $1`
	var data models.DataPerpustakaan
	err := r.db.GetContext(ctx, &data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PerpustakaanRepository) Update(ctx context.Context, p *models.DataPerpustakaan) error {
	query := `UPDATE inputdata SET namaPerpustakaan = $1, alamat = $2, jenisPerpustakaan = $3, nomorInduk = $4, jumlahCetak = $5, jumlahEbook = $6, jumlahSDM = $7, jumlahPengunjung = $8, jumlahAnggota = $9 WHERE id = $10`
	_, err := r.db.ExecContext(ctx, query, p.NamaPerpustakaan, p.Alamat, p.JenisPerpustakaan, p.NomorInduk, p.JumlahCetak, p.JumlahEbook, p.JumlahSDM, p.JumlahPengunjung, p.JumlahAnggota, p.ID)
	return err
}

func (r *PerpustakaanRepository) Delete(ctx context.Context, id string) error {	
	query := `DELETE FROM inputdata WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func NewPerpustakaanRepository(db *sqlx.DB) *PerpustakaanRepository {
	return &PerpustakaanRepository{db: db}
}
