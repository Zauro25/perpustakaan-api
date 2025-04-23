package repositories

import(
	"context"
	"errors"
	"github.com/Zauro25/perpus-app/internal/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}
func (r *AuthRepository) GetUserByUername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	query := `select id, username, password, role from users where username = $1`
	err := r.db.GetContext(ctx, &user, query, username)

	if err != nil {
		return nil, errors.New("User Tidak DItemukan")
	}
	return &user, nil
}