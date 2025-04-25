package database

import(
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"database/sql"
)

func database (){
	db_Host := "localhost"
	db_port := "5432"
	db_user := "admin"
	db_password  := "admin"
	db_name := "Manajemen_Perpustakaan"

	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s name=%s", db_Host, db_port, db_user, db_password, db_name)

	db, err := sql.Open("postgres", connect)
	if err != nil{
		log.Fatal("gagal terkoneksi ke postgresql", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil{
		log.Fatal("gagal ping ke database", err)
	}

	fmt.Println("Berhasil terkoneksi ke database")
}