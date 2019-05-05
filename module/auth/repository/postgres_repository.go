package repository

import (
	"database/sql"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/auth"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users/model"
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
		SELECT  
		user_id,
		username,
		email,
		password,
		first_name,
		last_name,
		photo_profile,
		created_at,
		updated_at 
		FROM tbl_users 
		WHERE email = ?
		LIMIT 1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	var mu model.User

	err = statement.QueryRow(u.Email).Scan(&mu.UserID, &mu.Email, &mu.Password, &mu.FirstName, &mu.LastName, &mu.PhotoProfile, &mu.CreatedAt, &mu.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &mu, err
}
