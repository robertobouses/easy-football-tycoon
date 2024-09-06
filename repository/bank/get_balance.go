package bank

import (
	"database/sql"
	"log"
)

func (r *repository) GetBalance() (int, error) {

	row := r.getBalance.QueryRow()
	var balance int

	err := row.Scan(&balance)
	log.Print("balance inicial en repository", balance)
	if err != nil {
		if err == sql.ErrNoRows {

			log.Print("No se encontraron resultados para el balance")
			return 0, nil
		}
		log.Print("Error fetching balance:", err)
		return 0, err
	}

	return balance, nil
}
