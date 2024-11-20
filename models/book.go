package models

type Book struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	PublishedYear  int    `json:"published_year"`
	Stock          int    `json:"stock"`
	AvailableStock int    `json:"available_stock"`
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, title, author, published_year, stock, available_stock FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedYear, &book.Stock, &book.AvailableStock); err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

// CreateBook inserts a new book into the database
func CreateBook(book *Book) error {
	_, err := db.Exec("INSERT INTO books (title, author, published_year, stock, available_stock) VALUES (?, ?, ?, ?, ?)",
		book.Title, book.Author, book.PublishedYear, book.Stock, book.AvailableStock)
	return err
}

// UpdateBookStock updates the stock and available stock of a book
func UpdateBookStock(bookID, stockChange int) error {
	_, err := db.Exec(`
        UPDATE books 
        SET stock = stock + ?, available_stock = available_stock + ? 
        WHERE id = ? AND available_stock + ? >= 0`,
		stockChange, stockChange, bookID, stockChange)
	return err
}

// BorrowBook reduces available stock of the book
func BorrowBookStock(bookID int) error {
	_, err := db.Exec(`
        UPDATE books 
        SET available_stock = available_stock - 1 
        WHERE id = ? AND available_stock > 0`, bookID)
	return err
}

// ReturnBook increases available stock of the book
func ReturnBookStock(bookID int) error {
	_, err := db.Exec(`
        UPDATE books 
        SET available_stock = available_stock + 1 
        WHERE id = ? AND available_stock < stock`, bookID)
	return err
}
