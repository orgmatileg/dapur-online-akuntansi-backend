package repository

import (
	"database/sql"

	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/auth"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users/model"
)

// postgresAuthRepository struct
type postgresAuthRepository struct {
	db *sql.DB
}

// NewAuthRepositoryPostgres NewUserRepositoryMysql
func NewAuthRepositoryPostgres(db *sql.DB) auth.Repository {
	return &postgresAuthRepository{db}
}

// LoginJWT
func (r *postgresAuthRepository) LoginJWT(u *model.User) (user *model.User, err error) {

	query := `
		SELECT *
		FROM tbl_users 
		WHERE email = $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	var mu model.User

	err = statement.QueryRow(u.Email).Scan(
		&mu.UserID,
		&mu.UserRoleID,
		&mu.Email,
		&mu.Password,
		&mu.FirstName,
		&mu.LastName,
		&mu.PhotoProfile,
		&mu.CreatedAt,
		&mu.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &mu, err
}
