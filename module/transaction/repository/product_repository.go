package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"
)

// postgresProductRepository struct
type postgresProductRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewProductRepositoryPostgres(db *sql.DB) product.Repository {
	return &postgresProductRepository{db}
}

// Save
func (r *postgresProductRepository) Save(m *model.Product) error {
	query := `
	INSERT INTO tbl_products
	(
		product_types_id,
		product_name,
		product_desc,
		product_capital_price,
		product_selling_price,
		product_image,
		created_at,
		updated_at
	)
	VALUES ( $1, $2, $3, $4, $5, $6, $7, $8 )
	RETURNING product_id`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	return statement.QueryRow(m.ProductTypes.ProductTypesID, m.Name, m.Description, m.CapitalPrice, m.SellingPrice, m.Image, m.CreatedAt, m.UpdatedAt).Scan(&m.ProductID)
}

// FindByID Example
func (r *postgresProductRepository) FindByID(id string) (*model.Product, error) {

	query := `
	SELECT *
	FROM v_products 
	WHERE product_id = $1`

	var m model.Product

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&m.ProductID,
		&m.Name,
		&m.Description,
		&m.CapitalPrice,
		&m.SellingPrice,
		&m.Image,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.ProductTypes.ProductTypesID,
		&m.ProductTypes.ProductTypesName,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// FindAll Example
func (r *postgresProductRepository) FindAll(limit, offset, order string) (model.ProductList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM v_products
	ORDER BY created_at %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var ml model.ProductList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m model.Product

		err = rows.Scan(
			&m.ProductID,
			&m.Name,
			&m.Description,
			&m.CapitalPrice,
			&m.SellingPrice,
			&m.Image,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.ProductTypes.ProductTypesID,
			&m.ProductTypes.ProductTypesName,
		)

		if err != nil {
			return nil, err
		}
		ml = append(ml, m)
	}

	return ml, nil
}

// Update Example
func (r *postgresProductRepository) Update(id string, m *model.Product) (rowAffected *string, err error) {

	query := `
	UPDATE tbl_products
	SET
		product_types_id = $1,
		product_name = $2,
		product_desc = $3,
		product_capital_price = $4,
		product_selling_price = $5,
		product_image = $6,
		updated_at = $7
	WHERE product_id = $8
	`

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.Exec(
		m.ProductTypes.ProductTypesID,
		m.Name,
		m.Description,
		m.CapitalPrice,
		m.SellingPrice,
		m.Image,
		m.UpdatedAt,
		id,
	)

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
func (r *postgresProductRepository) Delete(id string) error {

	query := `
	DELETE FROM tbl_products
	WHERE product_id = $1`

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
func (r *postgresProductRepository) IsExistsByID(id string) (isExist bool, err error) {

	query := "SELECT EXISTS(SELECT TRUE from tbl_products WHERE product_id = $1)"
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
func (r *postgresProductRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM v_products
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
