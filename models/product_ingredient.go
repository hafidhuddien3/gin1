package models

type ProductIngredient struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    IngredientID  uint `json:"ingredientId"`
    ProductID uint `json:"productId"`
    Multiply uint `json:"multiply"`
    Ingredient Ingredient `gorm:"foreignKey:IngredientID"`
}


// ingredientId
// productId
// multiply