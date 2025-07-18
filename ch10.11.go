package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(cus *customer,transac transaction) error{
	if transac.transactionType==transactionWithdrawal{
		if cus.balance<transac.amount{
			return errors.New("insufficient funds")
		}
		cus.balance-=transac.amount
		return nil
	}
	if transac.transactionType==transactionDeposit{
		cus.balance+=transac.amount
		return nil
	}
	return errors.New("unknown transaction type")
}