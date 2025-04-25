package models

import (

)

type DataPerpustakaan struct { 
	ID string `db:"ID" json:"ID"`
	NamaPerpustakaan string `db:"namaPerpustakaan "json:"namaPerpustakaan"`
	Alamat string `db:"alamat" json:"alamat"`
	JenisPerpustakaan string `db:"jenisPerpustakaan" json:"jenisPerpustakaan"`
	NomorInduk int `db:"nomorInduk" json:"nomorInduk"`
	JumlahCetak int `db:"jumlahCetak" json:"jumlahCetak"`
	JumlahEbook int `db:"jumlahEbook" json:"jumlahEbook"`
	JumlahSDM int `db:"jumlahSDM" json:"jumlahSDM"`
	JumlahPengunjung int `db:"jumlahPengunjung" json:"jumlahPengunjung"`
	JumlahAnggota int `db:"jumlahAnggota" json:"jumlahAnggota"`
}
