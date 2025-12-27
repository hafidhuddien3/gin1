package models

type Book struct {
	// ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
    ID     uint   `json:"id" gorm:"primaryKey"`
    Title  string `json:"title"`
    Author string `json:"author"`
}
