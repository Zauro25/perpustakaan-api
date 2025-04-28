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

func NewAuthController(repo *repositories.AuthRepository) *AuthController {

	return &AuthController{repo: repo}
}

func (c *AuthController) CreateUser(ctx *gin.Context) {
	var input models.User
	if input.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username tidak boleh kosong"})
		return
	}
	if input.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password tidak boleh kosong"})
	}
	if len(input.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password minimal 6 karakter"})
	}
	if input.Role != models.RoleAdminPerpus && input.Role != models.RoleAdminDPK {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Role tidak valid"})
	}
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
	var update models.User
	if update.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Username tidak boleh kosong"})
		return
	}
	if update.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password tidak boleh kosong"})
	}
	if len(update.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Password minimal 6 karakter"})
	}
	if update.Role != models.RoleAdminPerpus && update.Role != models.RoleAdminDPK {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Role tidak valid"})
	}
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

func (c *AuthController) Login(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.repo.GetUserByUsernameAndPassword(ctx.Request.Context(), input.Username, input.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (c *AuthController) GetUserByUsernameAndRole(ctx *gin.Context) {
	username := ctx.Query("username")
	role := ctx.Query("role")

	user, err := c.repo.GetUserByUsernameAndRole(ctx.Request.Context(), username, role)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
func (c *AuthController) GetUserByRole(ctx *gin.Context) {
	role := ctx.Query("role")

	users, err := c.repo.GetUserByRole(ctx.Request.Context(), role)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}
func (c *AuthController) GetUserByIDAndRole(ctx *gin.Context) {
	id := ctx.Param("id")
	role := ctx.Query("role")

	user, err := c.repo.GetUserByIDAndRole(ctx.Request.Context(), id, role)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}