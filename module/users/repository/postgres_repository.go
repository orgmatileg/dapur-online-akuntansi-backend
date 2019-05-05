package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/users/model"
)

// postgresUsersRepository struct
type postgresUsersRepository struct {
	db *sql.DB
}

// NewUserRepositoryPostgres
func NewUserRepositoryPostgres(db *sql.DB) users.Repository {
	return &postgresUsersRepository{db}
}

// Save User
func (r *postgresUsersRepository) Save(u *model.User) error {

	query := `
	INSERT INTO tbl_users (
	user_role_id,
	email,
	password,
	first_name,
	last_name,
	photo_profile,
	created_at,
	updated_at )
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	RETURNING user_id`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	err = statement.QueryRow(u.UserRoleID, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.CreatedAt, u.UpdatedAt).Scan(&u.UserID)
	if err != nil {
		return err
	}

	return nil
}

// FindByID User
func (r *postgresUsersRepository) FindByID(id string) (*model.User, error) {
	query := `
	SELECT *
	FROM tbl_users 
	WHERE user_id = $1`

	var user model.User

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&user.UserID, &user.UserRoleID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

// FindAll User
func (r *postgresUsersRepository) FindAll(limit, offset, order string) (model.Users, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_users
	ORDER BY created_at %s
	LIMIT %s 
	OFFSET %s`, order, limit, offset)

	var users model.Users

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User

		err = rows.Scan(&user.UserID, &user.UserRoleID, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Update User
func (r *postgresUsersRepository) Update(id string, u *model.User) (rowAffected *string, err error) {

	query := `
			UPDATE tbl_users SET 
			user_role_id = $1,
			email = $2, 
			password = $3,
			first_name = $4,
			last_name = $5,
			photo_profile = $6,
			updated_at = $7
			WHERE user_id = $8
		`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(u.UserRoleID, u.Email, u.Password, u.FirstName, u.LastName, u.PhotoProfile, u.UpdatedAt, id)

	if err != nil {
		return nil, err
	}

	rowsAffectedInt64, err := result.RowsAffected()

	if err != nil {
		return nil, err
	}

	rowsAffectedStr := strconv.FormatInt(rowsAffectedInt64, 10)

	rowAffected = &rowsAffectedStr

	return rowAffected, nil

}

// Delete User
func (r *postgresUsersRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_users
	WHERE user_id = $1`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	_, err = statement.Exec(id)

	if err != nil {
		return err
	}

	return nil
}

// IsExistsByID User
func (r *postgresUsersRepository) IsExistsByID(idUser string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_users WHERE user_id = $1)"

	statement, err := r.db.Prepare(query)

	if err != nil {
		return false, err
	}

	defer statement.Close()

	err = statement.QueryRow(idUser).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return isExist, nil
}

// Count Posts
func (r *postgresUsersRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_users
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return -1, err
	}

	defer statement.Close()

	err = statement.QueryRow().Scan(&count)

	if err != nil {
		return -1, err
	}

	return count, nil
}
