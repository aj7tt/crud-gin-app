package handlers

import (
    "github.com/gin-gonic/gin"
    "crud-gin-app/models"
    "net/http"
    "strconv"
)

var books []models.Book

// Create a new book
func CreateBook(c *gin.Context) {
    var newBook models.Book
    if err := c.ShouldBindJSON(&newBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    newBook.ID = uint(len(books) + 1)
    books = append(books, newBook)
    c.JSON(http.StatusCreated, newBook)
}

// Get all books
func GetBooks(c *gin.Context) {
    c.JSON(http.StatusOK, books)
}

// Get a single book by ID
func GetBookByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    for _, book := range books {
        if int(book.ID) == id {
            c.JSON(http.StatusOK, book)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// Update a book by ID
func UpdateBook(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, book := range books {
        if int(book.ID) == id {
            books[i] = updatedBook
            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// Delete a book by ID
func DeleteBook(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
        return
    }

    for i, book := range books {
        if int(book.ID) == id {
            books = append(books[:i], books[i+1:]...)
            c.JSON(http.StatusNoContent, gin.H{"message": "Book deleted"})
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
