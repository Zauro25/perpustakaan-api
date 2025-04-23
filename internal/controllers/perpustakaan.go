package controllers

import (
	"net/http"

	"github.com/Zauro25/perpus-app/internal/models"
	"github.com/Zauro25/perpus-app/internal/repositories"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PerpustakaanController struct {
	repo *repositories.PerpustakaanRepository
}

func NewPerpustakaanController(repo *repositories.PerpustakaanRepository) *PerpustakaanController{
	return &PerpustakaanController{repo: repo}
}

func (c *PerpustakaanController) CreatePerpustakaan(ctx *gin.Context) {
	var input models.Perpustakaan
	if err:= ctx.ShouldBindBodyWithJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = uuid.New()
	input.Status = "Draft"

	if err := c.repo.Create(ctx.Request.Context(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Data Gagal Disimpan"})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}