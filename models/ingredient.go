package models

type Ingredient struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Price string `json:"price"`
    Qty int `json:"qty"`
    QtyName string `json:"qtyName"`
}

// ingredientId
// name
// price
// qty
// qtyName default pieces