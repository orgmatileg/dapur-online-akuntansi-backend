package model

import (
	"time"

	m "github.com/orgmatileg/dapur-online-akuntansi-backend/module/product/model"
)

type Transaction struct {
	TransactionID    int64            `json:"transaction_id"`
	TransactionDataD []byte           `json:"-"`
	TransactionDataH TransactionDataH `json:"transaction_data"`
	CreatedBy        CreatedBy        `json:"created_by"`
	CreatedAt        time.Time        `json:"created_at"`
	UpdatedAt        time.Time        `json:"updated_at"`
}

type CreatedBy struct {
	UserID       int64  `json:"user_id"`
	UserFullName string `json:"user_fullname"`
}

type TransactionDataH struct {
	Carts                  []ProductInCart `json:"carts"`
	GrandTotalCapitalPrice int64           `json:"grand_total_capital_price"`
	GrandTotalSellingPrice int64           `json:"grant_total_selling_price"`
	GrandTotalProfit       int64           `json:"grand_total_profit"`
	NotesTransaction       string          `json:"notes_transaction"`
}

type ProductInCart struct {
	m.Product
	ActualSellingPrice int64  `json:"actual_selling_price"`
	NotesProduct       string `json:"notes_products"`
}

type TransactionList []Transaction

func NewTransaction() *Transaction {
	return &Transaction{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
