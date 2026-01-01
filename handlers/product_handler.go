package handlers

import (
    "net/http"
    "gin-quickstart/db"
    "gin-quickstart/models"
    "github.com/gin-gonic/gin"
)

// @Tags Product
// @Summary Create a new product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    if product.ID != 0 {
    var productU models.Product
    if err := db.DB.First(&productU, product.ID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&productU); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&productU)
    c.JSON(http.StatusOK, productU)
    } else {
    db.DB.Create(&product)
    c.JSON(http.StatusOK, product)
    }
}

// @Tags Product
// @Summary Get all products
// @Router /products [get]
func GetProducts(c *gin.Context) {
    var products []models.Product
    db.DB.Find(&products)
    c.JSON(http.StatusOK, products)
}

// @Tags Product
// @Summary Get a product by ID
// @Router /products/{id} [get]
func GetProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, product)
}

// @Tags Product
// @Summary Update a product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&product)
    c.JSON(http.StatusOK, product)
}

// @Tags Product
// @Summary Delete a product
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
    id := c.Param("id")
    var product models.Product
    if err := db.DB.Delete(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
