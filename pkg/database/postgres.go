package database

import(
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func PostgresConnection(dsn string) ( *sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil{
		return nil, fmt.Errorf("gagal terhubung ke database: %v", err)
	}
	return db, nil
}