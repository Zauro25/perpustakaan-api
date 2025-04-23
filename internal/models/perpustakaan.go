package models

import (
	"github.com/google/uuid"
)

type Perpustakaan struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Nama        string    `db:"nama" json:"nama"`
	Alamat      string    `db:"alamat" json:"alamat"`
	KoleksiBuku int       `db:"koleksi_buku" json:"koleksi_buku"`
	Status      string    `db:"status" json:"status"`
}
