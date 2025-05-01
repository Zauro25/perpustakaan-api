package models

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Role string

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     Role   `json:"role" binding:"required"`
}


const(
	RoleAdminPerpus Role = "admin_perpustakaan"
	RoleAdminDPK Role = "admin_dpk"
)

type Permission map[Role][]string

type User struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"password.omitempty` // Di-exclude dari JSON
	Role     Role      `db:"role" json:"role"`
}

func (u * User) hashPassword() {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password", err)
	}
	u.Password = string(hashedPassword)
	return
}
func (u *User) cekPassword(password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return errors.New("password tidak sesuai")
	}
	return nil
}
func validateRole(role Role) error {
	switch role {
	case RoleAdminPerpus, RoleAdminDPK:
		return nil
	default:
		return errors.New("role tidak valid")
}
}
var role Permission = map[Role][]string{
	RoleAdminPerpus: {"create", "read", "update", "delete", "submit"},
	RoleAdminDPK:   {"read_all, verify", "submit_to_perpusnas"},
}
