package main

import (
	"github.com/Zauro25/perpus-app/pkg/database" // Import package database
	"net/http"
	"fmt"
)

func homeLander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Aku sigma")
}
func main() {
	database.ConnectDB()
	http.HandleFunc("/", homeLander)
	fmt.Println("Server berjalan di port  8080")
	http.ListenAndServe(":8080", nil)
}