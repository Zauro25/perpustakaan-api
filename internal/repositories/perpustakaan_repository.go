package repositories

import(
	"context"
	"github.com/Zauro25/perpus-app/internal/models"

	"github.com/jmoiron/sqlx"
)

type PerpustakaanRepository struct {
	db *sqlx.DB
}
func NewDataRepository(db *sqlx.DB) *PerpustakaanRepository {
	return &PerpustakaanRepository{db: db}
}

func (r *PerpustakaanRepository) Create(ctx context.Context, p *models.DataPerpustakaan) error{
	query := `insert into dataperpus (id, namaperpustakaan, alamat, jenisperpustakaan, nomorinduk, jumlahcetak, jumlahebook, jumlahsdm, jumlahpengunjung, jumlahanggota) values (:id, :namaperpustakaan, :alamat, :jenisperpustakaan, :nomorinduk, :jumlahcetak, :jumlahebook, :jumlahsdm, :jumlahpengunjung, :jumlahanggota)`
	_, err := r.db.NamedExecContext(ctx, query, p)
	return err
}

func (r *PerpustakaanRepository) GetAll(ctx context.Context) ([]models.DataPerpustakaan, error) {
	query := `SELECT * FROM dataperpus`
	var data []models.DataPerpustakaan
	err := r.db.SelectContext(ctx, &data, query)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *PerpustakaanRepository) GetByID(ctx context.Context, id string) (*models.DataPerpustakaan, error) {
	query := `SELECT * FROM dataperpus WHERE id = $1`
	var data models.DataPerpustakaan
	err := r.db.GetContext(ctx, &data, query, id)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
func (r *PerpustakaanRepository) GetByName(ctx context.Context, namaperpustakaan string) (*models.DataPerpustakaan, error) {
	query := `SELECT * FROM dataperpus WHERE namaperpustakaan = $1`
	var data models.DataPerpustakaan
	err := r.db.GetContext(ctx, &data, query, namaperpustakaan)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *PerpustakaanRepository) Update(ctx context.Context, p *models.DataPerpustakaan) error {
	query := `UPDATE dataperpus SET namaPerpustakaan = $1, alamat = $2, jenisPerpustakaan = $3, nomorInduk = $4, jumlahCetak = $5, jumlahEbook = $6, jumlahSDM = $7, jumlahPengunjung = $8, jumlahAnggota = $9 WHERE id = $10`
	_, err := r.db.ExecContext(ctx, query, p.NamaPerpustakaan, p.Alamat, p.JenisPerpustakaan, p.NomorInduk, p.JumlahCetak, p.JumlahEbook, p.JumlahSDM, p.JumlahPengunjung, p.JumlahAnggota, p.ID)
	return err
}

func (r *PerpustakaanRepository) Delete(ctx context.Context, id string) error {	
	query := `DELETE FROM dataperpus WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func NewPerpustakaanRepository(db *sqlx.DB) *PerpustakaanRepository {
	return &PerpustakaanRepository{db: db}
}
