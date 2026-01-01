package handlers

import (
    "net/http"
    "gin-quickstart/db"
    "gin-quickstart/models"
    "github.com/gin-gonic/gin"
)

// @Tags Ingredient
// @Summary Create a new ingredient
// @Router /ingredients [post]
func CreateIngredient(c *gin.Context) {
    var ingredient models.Ingredient

if ingredient.ID != 0 { // client provided an ID 
        if err := db.DB.First(&ingredient, ingredient.ID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
        return
    }
    if err := c.ShouldBindJSON(&ingredient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&ingredient)
} else {
    if err := c.ShouldBindJSON(&ingredient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Create(&ingredient)
}


    c.JSON(http.StatusOK, ingredient)
}

// @Tags Ingredient
// @Summary Get all ingredients
// @Router /ingredients [get]
func GetIngredients(c *gin.Context) {
    var ingredients []models.Ingredient
    db.DB.Find(&ingredients)
    c.JSON(http.StatusOK, ingredients)
}

// @Tags Ingredient
// @Summary Get a ingredient by ID
// @Router /ingredients/{id} [get]
func GetIngredient(c *gin.Context) {
    id := c.Param("id")
    var ingredient models.Ingredient
    if err := db.DB.First(&ingredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
        return
    }
    c.JSON(http.StatusOK, ingredient)
}

// @Tags Ingredient
// @Summary Update a ingredient
// @Router /ingredients/{id} [put]
func UpdateIngredient(c *gin.Context) {
    id := c.Param("id")
    var ingredient models.Ingredient
    if err := db.DB.First(&ingredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
        return
    }
    if err := c.ShouldBindJSON(&ingredient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.DB.Save(&ingredient)
    c.JSON(http.StatusOK, ingredient)
}

// @Tags Ingredient
// @Summary Delete a ingredient
// @Router /ingredients/{id} [delete]
func DeleteIngredient(c *gin.Context) {
    id := c.Param("id")
    var ingredient models.Ingredient
    if err := db.DB.Delete(&ingredient, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Ingredient not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Ingredient deleted"})
}
