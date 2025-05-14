package models

import (
	"github.com/google/uuid"

)

type DataPerpustakaan struct { 
	ID uuid.UUID `db:"id" json:"id"`
	NamaPerpustakaan string `db:"namaperpustakaan" json:"namaperpustakaan"`
	Alamat string `db:"alamat" json:"alamat"`
	JenisPerpustakaan string `db:"jenisperpustakaan" json:"jenisperpustakaan"`
	NomorInduk int `db:"nomorinduk" json:"nomorinduk"`
	JumlahCetak int `db:"jumlahcetak" json:"jumlahcetak"`
	JumlahEbook int `db:"jumlahebook" json:"jumlahebook"`
	JumlahSDM int `db:"jumlahsdm" json:"jumlahsdm"`
	JumlahPengunjung int `db:"jumlahpengunjung" json:"jumlahpengunjung"`
	JumlahAnggota int `db:"jumlahanggota" json:"jumlahanggota"`
}
