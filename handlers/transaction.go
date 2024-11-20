package handlers

import (
	"library-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BorrowBook(c *gin.Context) {
	var req struct {
		UserID int `json:"user_id"`
		BookID int `json:"book_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := models.BorrowBook(req.UserID, req.BookID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book borrowed successfully"})
}

func ReturnBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	if err := models.ReturnBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book returned successfully"})
}

func GetTransactions(c *gin.Context) {
	transactions, err := models.GetAllTransactions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
		return
	}
	c.JSON(http.StatusOK, transactions)
}
