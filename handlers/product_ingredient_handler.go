package handlers

import (
    "net/http"
    "gin-quickstart/db"
    "gin-quickstart/models"
    "github.com/gin-gonic/gin"
)

// @Tags Product Ingredient
// @Summary Create a new productIngredient
// @Router /productIngredients [post]
func CreateProductIngredient(c *gin.Context) {
    var productIngredients []models.ProductIngredient

    if err := c.ShouldBindJSON(&productIngredients); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // if len(productIngredients) == 0 {
    //     c.JSON(http.StatusBadRequest, gin.H{"error": "No product ingredients provided"})
    //     return
    // }

    // Get productId from the first item (assuming all belong to same product)
    productId := productIngredients[0].ProductID

    // Delete existing rows for this productId
    if err := db.DB.
    Where("product_id = ?", productId).
    Delete(&models.ProductIngredient{}).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Insert new rows
    if err := db.DB.Create(&productIngredients).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, productIngredients)
}

// func CreateProductIngredient(c *gin.Context) {
//     var productIngredient []models.ProductIngredient
//     if err := c.ShouldBindJSON(&productIngredient); err != nil {
//         c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
//     }
//     db.DB.Create(&productIngredient)
//     c.JSON(http.StatusOK, productIngredient)
// }

// @Tags Product Ingredient
// @Summary Get all productIngredients
// @Router /productIngredients [get]
func GetProductIngredients(c *gin.Context) {
    var productIngredients []models.ProductIngredient
    db.DB.
    Joins("Ingredient").
    Find(&productIngredients)
    c.JSON(http.StatusOK, productIngredients)
}

// @Tags Product Ingredient
// @Summary Get a productIngredient by ID
// @Router /productIngredients/{id} [get]
func GetProductIngredient(c *gin.Context) {
    id := c.Param("id")
    var productIngredient models.ProductIngredient
    if err := db.DB.First(&productIngredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ProductIngredient not found"})
        return
    }
    c.JSON(http.StatusOK, productIngredient)
}

// @Tags Product Ingredient
// @Summary Get a productIngredient by productID
// @Router /productIngredients/product/{productId} [get]
func GetProductIngredientByProductId(c *gin.Context) {
    productId := c.Param("productId")
    // var productIngredients models.ProductIngredient
    var productIngredients []models.ProductIngredient
    // var productIngredients []ProductIngredient
    if err := db.DB.
    Joins("Ingredient").
    Where("product_id = ?", productId).
    Find(&productIngredients).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No ingredients found for this product"})
        return
    }

    // if len(productIngredients) == 0 { c.JSON(http.StatusNotFound, gin.H{"error": "No ingredients found for this product"}) return }

    c.JSON(http.StatusOK, productIngredients)
}

// type Result struct {
//     ProductID    uint
//     IngredientID uint
//     IngredientName string
// }

// var results []Result
// db.DB.Table("product_ingredients").
//     Select("product_ingredients.product_id, product_ingredients.ingredient_id, ingredients.name as ingredient_name").
//     Joins("inner join ingredients on product_ingredients.ingredient_id = ingredients.id").
//     Where("product_ingredients.product_id = ?", productId).
//     Scan(&results)


// @Tags Product Ingredient
// @Summary Update a productIngredient
// @Router /productIngredients/{id} [put]
func UpdateProductIngredient(c *gin.Context) {
    id := c.Param("id")
    var productIngredient models.ProductIngredient
    if err := db.DB.First(&productIngredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ProductIngredient not found"})
        return
    }
    if err := c.ShouldBindJSON(&productIngredient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&productIngredient)
    c.JSON(http.StatusOK, productIngredient)
}

// @Tags Product Ingredient
// @Summary Delete a productIngredient
// @Router /productIngredients/{id} [delete]
func DeleteProductIngredient(c *gin.Context) {
    id := c.Param("id")
    var productIngredient models.ProductIngredient
    if err := db.DB.Delete(&productIngredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "ProductIngredient not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "ProductIngredient deleted"})
}
