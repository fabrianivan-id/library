package handlers

import (
	"library-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetBooks retrieves all books from the database
func GetBooks(c *gin.Context) {
	// Fetch books from the database through the model
	books, err := models.GetAllBooks()
	if err != nil {
		// Return error if there is a problem fetching the books
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve books"})
		return
	}

	// Send the list of books as a JSON response
	c.JSON(http.StatusOK, books)
}

// AddBook adds a new book to the database
func AddBook(c *gin.Context) {
	var book models.Book

	// Bind incoming JSON data to the Book model
	if err := c.ShouldBindJSON(&book); err != nil {
		// Return an error response if the binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Validate book data (for example, check if stock is non-negative)
	if book.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock cannot be negative"})
		return
	}

	// Create the new book in the database through the model
	if err := models.CreateBook(&book); err != nil {
		// Return an error response if the database operation fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add book"})
		return
	}

	// Return success response if the book is added
	c.JSON(http.StatusCreated, gin.H{"message": "Book added successfully"})
}

// UpdateBookStock updates the stock of a specific book by ID
func UpdateBookStock(c *gin.Context) {
	// Extract the book ID from the URL parameter
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Return error if ID parsing fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var updateData struct {
		Stock int `json:"stock"`
	}

	// Bind the incoming JSON data to the struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		// Return error if binding the JSON fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Ensure the stock value is valid (non-negative)
	if updateData.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Stock cannot be negative"})
		return
	}

	// Call the model method to update the book's stock in the database
	if err := models.UpdateBookStock(id, updateData.Stock); err != nil {
		// Return error if the update fails
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update book stock"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Book stock updated successfully"})
}
