package models

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Gift        bool   `json:"is_gift"`
}
