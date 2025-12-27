package handlers

import (
    "net/http"
    "gin-quickstart/db"
    "gin-quickstart/models"
    "github.com/gin-gonic/gin"
)

// @Summary Create a new book
// @Router /books [post]
func CreateBook(c *gin.Context) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&book)
    c.JSON(http.StatusOK, book)
}

// @Summary Get all books
// @Router /books [get]
func GetBooks(c *gin.Context) {
    var books []models.Book
    db.DB.Find(&books)
    c.JSON(http.StatusOK, books)
}

// @Summary Get a book by ID
// @Router /books/{id} [get]
func GetBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := db.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// @Summary Update a book
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := db.DB.First(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&book)
    c.JSON(http.StatusOK, book)
}

// @Summary Delete a book
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
    id := c.Param("id")
    var book models.Book
    if err := db.DB.Delete(&book, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
