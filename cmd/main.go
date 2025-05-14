package main

import (
	"github.com/Zauro25/perpus-app/pkg/database" 
	"github.com/Zauro25/perpus-app/internal/controllers"	
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/Zauro25/perpus-app/internal/repositories"


	"log"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}
	authRepo := repositories.NewAuthRepository(db)
	authController := controllers.NewAuthController(authRepo)
	perpusRepo := repositories.NewPerpustakaanRepository(db)
	perpusController := controllers.NewPerpustakaanController(perpusRepo)
	defer db.Close()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Aku sigma")
	})
	router.POST("/inputdata", perpusController.CreatePerpustakaan)
	router.GET("/data", perpusController.GetAllPerpustakaan)
	router.GET("/data/:id", perpusController.GetPerpustakaanByID)
	router.GET("/data/nama/:namaperpustakaan", perpusController.GetPerpustakaanByName)
	router.PUT("/updatedata/:id", perpusController.UpdatePerpustakaan)
	router.DELETE("/deletedata/:id", perpusController.DeletePerpustakaan)
	router.POST("/login", authController.Login)
	router.POST("/register", authController.CreateUser)
	router.GET("/users", authController.GetAllUsers)
	fmt.Println("Server berjalan di port 8080")
	router.Run(":8080")
}