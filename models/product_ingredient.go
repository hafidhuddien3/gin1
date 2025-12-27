package models

type ProductIngredient struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    IngredientId  string `json:"ingredientId"`
    ProductId string `json:"productId"`
    Multiply int `json:"multiply"`
}


// ingredientId
// productId
// multiply