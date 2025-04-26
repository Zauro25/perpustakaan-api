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

func (r *AuthRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `insert into users (id, username, password, role) values ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Password, user.Role)
	return err
}
func (r *AuthRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `update users set username = $1, password = $2, role = $3 where id = $4`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password, user.Role, user.ID)
	return err
}
func (r *AuthRepository) DeleteUser(ctx context.Context, id string) error {
	query := `delete from users where id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
func (r *AuthRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	query := `select id, username, password, role from users`
	var users []models.User
	err := r.db.SelectContext(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (r *AuthRepository) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	query := `select id, username, password, role from users where id = $1`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndPassword(ctx context.Context, username, password string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and password = $2`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRole(ctx context.Context, username, role string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByRole(ctx context.Context, role string) ([]models.User, error) {
	query := `select id, username, password, role from users where role = $1`
	var users []models.User
	err := r.db.SelectContext(ctx, &users, query, role)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (r *AuthRepository) GetUserByIDAndRole(ctx context.Context, id, role string) (*models.User, error) {
	query := `select id, username, password, role from users where id = $1 and role = $2`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, id, role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndID(ctx context.Context, username, id string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and id = $2`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndPasswordAndRole(ctx context.Context, username, password, role string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and password = $2 and role = $3`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, password, role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndPasswordAndID(ctx context.Context, username, password, id string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and password = $2 and id = $3`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, password, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndID(ctx context.Context, username, role, id string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndPassword(ctx context.Context, username, role, password string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and password = $3`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndIDAndPassword(ctx context.Context, username, role, id, password string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3 and password = $4`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id, password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndIDAndPasswordAndEmail(ctx context.Context, username, role, id, password, email string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3 and password = $4 and email = $5`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id, password, email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndIDAndPasswordAndEmailAndPhone(ctx context.Context, username, role, id, password, email, phone string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3 and password = $4 and email = $5 and phone = $6`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id, password, email, phone)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndIDAndPasswordAndEmailAndPhoneAndAddress(ctx context.Context, username, role, id, password, email, phone, address string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3 and password = $4 and email = $5 and phone = $6 and address = $7`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id, password, email, phone, address)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *AuthRepository) GetUserByUsernameAndRoleAndIDAndPasswordAndEmailAndPhoneAndAddressAndCity(ctx context.Context, username, role, id, password, email, phone, address, city string) (*models.User, error) {
	query := `select id, username, password, role from users where username = $1 and role = $2 and id = $3 and password = $4 and email = $5 and phone = $6 and address = $7 and city = $8`
	var user models.User
	err := r.db.GetContext(ctx, &user, query, username, role, id, password, email, phone, address, city)
	if err != nil {
		return nil, err
	}
	return &user, nil
}