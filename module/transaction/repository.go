package product

import "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/model"

// Repository interface
type Repository interface {
	Save(*model.Transaction) error
	FindByID(id string) (*model.Transaction, error)
	FindAll(limit, offset, order string) (ml model.TransactionList, err error)
	Update(id string, m *model.Transaction) (*string, error)
	Delete(id string) error
	IsExistsByID(id string) (bool, error)
	Count() (int64, error)
}
