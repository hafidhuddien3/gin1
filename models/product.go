package models

type Product struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    Name  string `json:"name"`
    Price int `json:"price"`
    Qty int `json:"qty"`
    QtyName string `json:"qtyName"`
}

// productId
// name
// ingredientPrice // dinamic
// price
// qty
// qtyName