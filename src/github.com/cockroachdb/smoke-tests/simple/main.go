package main

// Taken from https://www.cockroachlabs.com/docs/build-a-test-app.html

import (
	"flag"
	"database/sql"
	"fmt"
	"log"

	"github.com/cockroachdb/cockroach-go/crdb"
)

func main() {
	flag.Parse()

	db, err := sql.Open("postgres", flag.Arg(0))
	if err != nil {
		log.Fatalf("Connecting to DB: %s", err)
	}

	defer db.Close()

	err = makeAccounts(db)
	if err != nil {
		log.Fatalf("Making accounts: %s", err)
	}

	err = transferFunds(db, 1 /* from acct# */, 2 /* to acct# */, 100 /* amount */);
	if err != nil {
		log.Fatalf("Transfering funds: %s", err)
	}

	fmt.Println("Success")
}

func makeAccounts(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE accounts (id INT PRIMARY KEY, balance INT)")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO accounts (id, balance) VALUES (1, 1000), (2, 250)")
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT id, balance FROM accounts")
	if err != nil {
		return err
	}

	defer rows.Close()

	fmt.Println("Initial balances:")

	for rows.Next() {
		var id, balance int

		err := rows.Scan(&id, &balance)
		if err != nil {
			return err
		}

		fmt.Printf("- %d: %d\n", id, balance)
	}

	return nil
}

func transferFunds(db *sql.DB, from int, to int, amount int) error {
	return crdb.ExecuteTx(db, func(tx *sql.Tx) error {
		var fromBalance int

		err := tx.QueryRow("SELECT balance FROM accounts WHERE id = $1", from).Scan(&fromBalance)
		if err != nil {
			return err
		}

		if fromBalance < amount {
			return fmt.Errorf("insufficient funds")
		}

		_, err = tx.Exec("UPDATE accounts SET balance = balance - $1 WHERE id = $2", amount, from)
		if err != nil {
			return err
		}

		_, err = tx.Exec("UPDATE accounts SET balance = balance + $1 WHERE id = $2", amount, to)

		return err
	})
}
