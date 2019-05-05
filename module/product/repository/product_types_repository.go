package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	productTypes "github.com/orgmatileg/SOLO-YOLO-BACKEND/module/product_types"
	"github.com/orgmatileg/SOLO-YOLO-BACKEND/module/product_types/model"
)

// postgresProductTypesRepository struct
type postgresProductTypesRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewProductTypesRepositoryPostgres(db *sql.DB) productTypes.Repository {
	return &postgresProductTypesRepository{db}
}

// Save
func (r *postgresProductTypesRepository) Save(m *model.ProductTypes) error {
	query := `
	INSERT INTO tbl_product_types
	(
		product_types_name
	)
	VALUES ( $1 )
	RETURNING product_types_id`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	return statement.QueryRow(m.ProductTypesName).Scan(&m.ProductTypesID)
}

// FindByID Example
func (r *postgresProductTypesRepository) FindByID(id string) (*model.ProductTypes, error) {

	query := `
	SELECT *
	FROM tbl_product_types 
	WHERE product_types_id = $1`

	var m model.ProductTypes

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(&m.ProductTypesID, &m.ProductTypesName)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// FindAll Example
func (r *postgresProductTypesRepository) FindAll(limit, offset, order string) (model.ProductTypesList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM tbl_product_types
	ORDER BY product_types_id %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var ml model.ProductTypesList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m model.ProductTypes

		err = rows.Scan(&m.ProductTypesID, &m.ProductTypesName)

		if err != nil {
			return nil, err
		}
		ml = append(ml, m)
	}

	return ml, nil
}

// Update Example
func (r *postgresProductTypesRepository) Update(id string, m *model.ProductTypes) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_product_types
	SET
		product_types_name = $1
	WHERE product_types_id = $2`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(m.ProductTypesName, id)

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
func (r *postgresProductTypesRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_product_types
	WHERE product_types_id = $1`

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
func (r *postgresProductTypesRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_product_types WHERE product_types_id = $1)"
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
func (r *postgresProductTypesRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM tbl_product_types
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
