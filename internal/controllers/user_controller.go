package controllers

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/Zauro25/perpus-app/internal/models"
	"github.com/Zauro25/perpus-app/internal/repositories"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	repo *repositories.AuthRepository
}

func NerAuthController(repo *repositories.AuthRepository) *AuthController {

	return &AuthController{repo: repo}
}

func (c *AuthController) CreateUser(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.ID = uuid.New()

	if err := c.repo.CreateUser(ctx.Request.Context(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}
	ctx.JSON(http.StatusCreated, input)
}

func (c *AuthController) GetAllUsers(ctx *gin.Context) {
	data, err := c.repo.GetAllUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
func (c *AuthController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	data, err := c.repo.GetUserByID(ctx.Request.Context(), id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, data)
}
func (c *AuthController) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	parsedID, err := uuid.Parse(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}
	input.ID = parsedID

	if err := c.repo.UpdateUser(ctx.Request.Context(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User update failed"})
		return
	}
	ctx.JSON(http.StatusOK, input)
}
func (c *AuthController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.repo.DeleteUser(ctx.Request.Context(), id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User deletion failed"})
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}