package controllers

import (
	"net/http"
	"github.com/google/uuid"

	"github.com/Zauro25/perpus-app/internal/models"
	"github.com/Zauro25/perpus-app/internal/repositories"
	"github.com/gin-gonic/gin"
)

type inputdataController struct {
	repo *repositories.InputDataRepository
}

func NewPerpustakaanController(repo *repositories.InputDataRepository) *inputdataController{
	return &inputdataController{repo: repo}
}

func (c *inputdataController) CreatePerpustakaan(ctx *gin.Context) {
	var input models.DataPerpustakaan
	if err:= ctx.ShouldBindBodyWithJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = uuid.New().String()
	
	if err := c.repo.Create(ctx.Request.Context(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H {"error": "Data Gagal Disimpan"})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}