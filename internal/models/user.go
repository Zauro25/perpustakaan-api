package models

import "github.com/google/uuid"

type Role string

const(
	RoleAdminPerpus Role = "admin_perpustakaan"
	RoleAdminDPK Role = "admin_dpk"
)

type User struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Username string    `db:"username" json:"username"`
	Password string    `db:"password" json:"-"` // Di-exclude dari JSON
	Role     Role      `db:"role" json:"role"`
}