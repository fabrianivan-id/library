package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// API Documentation
func GetDocumentation(c *gin.Context) {
	docs := map[string]interface{}{
		"endpoints": []map[string]string{
			{
				"method":      "POST",
				"path":        "/login",
				"description": "Login with username and password to receive a JWT token.",
			},
			{
				"method":      "GET",
				"path":        "/books",
				"description": "Retrieve the list of all books.",
			},
			{
				"method":      "POST",
				"path":        "/books",
				"description": "Add a new book (Admin only).",
			},
			{
				"method":      "PUT",
				"path":        "/books/:id",
				"description": "Update stock of a specific book (Admin only).",
			},
			{
				"method":      "POST",
				"path":        "/transactions",
				"description": "Borrow a book.",
			},
			{
				"method":      "PUT",
				"path":        "/transactions/:id/return",
				"description": "Return a borrowed book.",
			},
			{
				"method":      "GET",
				"path":        "/transactions",
				"description": "Retrieve all transactions (Admin only).",
			},
		},
	}

	c.JSON(http.StatusOK, docs)
}
