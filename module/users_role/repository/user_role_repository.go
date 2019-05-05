package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	// "strconv"
	// "time"

	example "github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/users_role/model"
)

// postgresUserRoleRepository struct
type postgresUserRoleRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewUsersRoleRepositoryPostgres(db *sql.DB) example.Repository {
	return &postgresUserRoleRepository{db}
}

// Save
func (r *postgresUserRoleRepository) Save(m *model.UserRole) error {
	query := `
	INSERT INTO tbl_users_role
	(
		user_role_name
	)
	VALUES ( $1 )
	RETURNING user_role_id`

	statement, err := r.db.Prepare(query)

	if err != nil {
		fmt.Println("Error prepare statement")
		return err
	}

	defer statement.Close()

	return statement.QueryRow(m.UserRoleName).Scan(&m.UserRoleID)
}

// FindByID Example
func (r *postgresUserRoleRepository) FindByID(id string) (*model.UserRole, error) {

	query := `
	SELECT *
	FROM tbl_users_role 
	WHERE user_role_id = $1`

	var m model.UserRole

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&m.UserRoleID, &m.UserRoleName)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// FindAll Example
func (r *postgresUserRoleRepository) FindAll(limit, offset, order string) (model.UserRoleList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_users_role
	ORDER BY user_role_id %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var ml model.UserRoleList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m model.UserRole

		err = rows.Scan(&m.UserRoleID, &m.UserRoleName)

		if err != nil {
			return nil, err
		}
		ml = append(ml, m)
	}

	return ml, nil
}

// Update Example
func (r *postgresUserRoleRepository) Update(id string, m *model.UserRole) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_users_role
	SET
		user_role_name = $1
	WHERE user_role_id = $2`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(m.UserRoleName, id)

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

// Delete Example
func (r *postgresUserRoleRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_users_role
	WHERE user_role_id = $1`

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

// IsExistsByID Example
func (r *postgresUserRoleRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_users_role WHERE user_role_id = $1)"
	statement, err := r.db.Prepare(query)

	if err != nil {
		return false, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&isExist)

	if err != nil {
		return false, err
	}

	return isExist, nil
}

// Count Posts
func (r *postgresUserRoleRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_users_role
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
