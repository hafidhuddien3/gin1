package models

type ProductIngredient struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    IngredientId  uint `json:"ingredientId"`
    ProductId uint `json:"productId"`
    Multiply uint `json:"multiply"`
}


// ingredientId
// productId
// multiply