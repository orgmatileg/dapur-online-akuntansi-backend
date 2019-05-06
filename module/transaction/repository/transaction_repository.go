package repository

import (
	"database/sql"
	"fmt"
	"strconv"

	transaction "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/model"
)

// postgresTransactionRepository struct
type postgresTransactionRepository struct {
	db *sql.DB
}

// NewExampleRepositoryMysql NewUserRepositoryMysql
func NewTransactionRepositoryPostgres(db *sql.DB) transaction.Repository {
	return &postgresTransactionRepository{db}
}

// Save
func (r *postgresTransactionRepository) Save(m *model.Transaction) error {
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

	return statement.QueryRow(m.TransactionID).Scan(&m.TransactionID)
}

// FindByID Example
func (r *postgresTransactionRepository) FindByID(id string) (*model.Transaction, error) {

	query := `
	SELECT *
	FROM v_transaction 
	WHERE transaction_id = $1`

	var m model.Transaction

	statement, err := r.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&m.TransactionID,
		&m.TransactionDataD,
		&m.CreatedAt,
		&m.UpdatedAt,
		&m.CreatedBy.UserID,
		&m.CreatedBy.UserFullName,
	)

	if err != nil {
		return nil, err
	}

	return &m, nil
}

// FindAll Example
func (r *postgresTransactionRepository) FindAll(limit, offset, order string) (model.TransactionList, error) {

	query := fmt.Sprintf(`
	SELECT *
	FROM v_transaction
	ORDER BY created_at %s
	LIMIT %s
	OFFSET %s`, order, limit, offset)

	var ml model.TransactionList

	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var m model.Transaction

		err = rows.Scan(
			&m.TransactionID,
			&m.TransactionDataD,
			&m.CreatedAt,
			&m.UpdatedAt,
			&m.CreatedBy.UserID,
			&m.CreatedBy.UserFullName,
		)

		if err != nil {
			return nil, err
		}
		ml = append(ml, m)
	}

	return ml, nil
}

// Update Example
func (r *postgresTransactionRepository) Update(id string, m *model.Transaction) (rowAffected *string, err error) {

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
func (r *postgresTransactionRepository) Delete(id string) error {

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
func (r *postgresTransactionRepository) IsExistsByID(id string) (isExist bool, err error) {

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
func (r *postgresTransactionRepository) Count() (count int64, err error) {

	query := `
	SELECT COUNT(*)
	FROM v_transaction
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
