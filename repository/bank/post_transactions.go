package bank

import (
	"log"
)

func (r *repository) PostTransactions(amount int, balance int, signings string, transactionType string) error {
	_, err := r.postTransactions.Exec(amount, balance, signings, transactionType)
	if err != nil {
		log.Print("Error executing PostTransactions statement:", err)
		return err
	}

	return nil
}
