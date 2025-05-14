package controllers

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/Zauro25/perpus-app/internal/models"
	"github.com/Zauro25/perpus-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"log"
)

type DataController struct {
	repo *repositories.PerpustakaanRepository
}

func NewPerpustakaanController(repo *repositories.PerpustakaanRepository) *DataController{
	return &DataController{repo: repo}
}

func (c *DataController) CreatePerpustakaan(ctx *gin.Context) {
	var input models.DataPerpustakaan
	if err:= ctx.ShouldBindBodyWithJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input = models.DataPerpustakaan{
		ID: uuid.New(),
		NamaPerpustakaan: input.NamaPerpustakaan,
		Alamat: input.Alamat,
		JenisPerpustakaan: input.JenisPerpustakaan,
		NomorInduk: input.NomorInduk,
		JumlahCetak: input.JumlahCetak,
		JumlahEbook: input.JumlahEbook,
		JumlahSDM: input.JumlahSDM,
		JumlahPengunjung: input.JumlahPengunjung,
		JumlahAnggota: input.JumlahAnggota,
	}
	if err := c.repo.Create(ctx.Request.Context(), &input); err != nil {
		log.Printf("Error Create: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Menyimpan Data"})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}

func (c *DataController) GetAllPerpustakaan(ctx *gin.Context) {
	data, err := c.repo.GetAll(ctx.Request.Context())
	if err != nil {
		log.Printf("Error GetAll: %v", err) 
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Mengambil Data"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}

func (c *DataController) GetPerpustakaanByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.repo.GetByID(ctx.Request	.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data Tidak Ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
func (c *DataController) GetPerpustakaanByName(ctx *gin.Context) {
	namaperpustakaan := ctx.Param("namaperpustakaan")
	data, err := c.repo.GetByName(ctx.Request.Context(), namaperpustakaan)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data Tidak Ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
func (c *DataController) UpdatePerpustakaan(ctx *gin.Context) {
	id := ctx.Param("id")
	var input models.DataPerpustakaan
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	input.ID = parsedID

	if err := c.repo.Update(ctx.Request.Context(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Memperbarui Data"})
		return
	}
	ctx.JSON(http.StatusOK, input)
}
func (c *DataController) DeletePerpustakaan(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.repo.Delete(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal Menghapus Data"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

func (c *DataController) submit_to_dpk(ctx *gin.Context, id string) {
	data, err := c.repo.GetByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Data Tidak Ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
	