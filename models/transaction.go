package models

import (
	"errors"
	"time"
)

type Transaction struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	BookID     int       `json:"book_id"`
	BorrowDate time.Time `json:"borrow_date"`
	ReturnDate time.Time `json:"return_date,omitempty"`
}

// BorrowBook records a transaction for borrowing a book
func BorrowBook(userID, bookID int) error {
	_, err := db.Exec(`
        INSERT INTO transactions (user_id, book_id, borrow_date) 
        VALUES (?, ?, NOW())`, userID, bookID)
	return err
}

// ReturnBook updates the transaction to mark a book as returned and updates the stock
func ReturnBook(transactionID int) error {
	// Get the book ID associated with the transaction
	var bookID int
	err := db.QueryRow(`
        SELECT book_id 
        FROM transactions 
        WHERE id = ? AND return_date IS NULL`, transactionID).
		Scan(&bookID)
	if err != nil {
		return errors.New("transaction not found or already returned")
	}

	// Mark the book as returned in the transaction
	_, err = db.Exec(`
        UPDATE transactions 
        SET return_date = NOW() 
        WHERE id = ?`, transactionID)
	if err != nil {
		return err
	}

	// Update the stock of the book
	return ReturnBookStock(bookID)
}

// GetAllTransactions retrieves all transactions from the database
func GetAllTransactions() ([]Transaction, error) {
	rows, err := db.Query(`
        SELECT id, user_id, book_id, borrow_date, return_date 
        FROM transactions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.ID, &t.UserID, &t.BookID, &t.BorrowDate, &t.ReturnDate); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}
	return transactions, nil
}
