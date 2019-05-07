package usecase

import (
	"encoding/json"
	"log"

	product "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction"
	transaction "github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction"
	"github.com/orgmatileg/dapur-online-akuntansi-backend/module/transaction/model"
)

type transactionUsecase struct {
	transactionRepo transaction.Repository
}

func NewTransactionUsecase(r product.Repository) transaction.Usecase {
	return &transactionUsecase{
		transactionRepo: r,
	}
}

func (u *transactionUsecase) Save(m *model.Transaction) (err error) {
	m.TransactionDataD, err = json.Marshal(m.TransactionDataH)

	if err != nil {
		return err
	}

	return u.transactionRepo.Save(m)
}

func (u *transactionUsecase) FindByID(id string) (m *model.Transaction, err error) {
	m, err = u.transactionRepo.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = json.Unmarshal(m.TransactionDataD, &m.TransactionDataH)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

func (u *transactionUsecase) FindAll(limit, offset, order string) (ml model.TransactionList, count int64, err error) {
	ml, err = u.transactionRepo.FindAll(limit, offset, order)
	if err != nil {
		log.Println(err)
		return nil, -1, err
	}

	for i, v := range ml {
		err = json.Unmarshal(v.TransactionDataD, &ml[i].TransactionDataH)
		if err != nil {
			log.Println(err)
			return nil, -1, err
		}
	}

	count, err = u.transactionRepo.Count()
	return
}

func (u *transactionUsecase) Update(id string, m *model.Transaction) (rowAffected *string, err error) {
	return u.transactionRepo.Update(id, m)
}

func (u *transactionUsecase) Delete(id string) (err error) {
	return u.transactionRepo.Delete(id)
}

func (u *transactionUsecase) IsExistsByID(id string) (isExist bool, err error) {
	return u.transactionRepo.IsExistsByID(id)
}

func (u *transactionUsecase) Count() (count int64, err error) {
	return u.transactionRepo.Count()
}
