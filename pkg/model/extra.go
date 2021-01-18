package model

// Extra Object
type Extra struct {
	ID      int    `gorm:"id" json:"id"`
	Name    string `gorm:"name" json:"name"`
	StoreID string `gorm:"store_id" json:"store_id"`
	Price   int8   `gorm:"price" json:"price"`
}
