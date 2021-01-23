package model

// Extra Object
type Extra struct {
	ID      int    `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	StoreID string `gorm:"column:store_id" json:"store_id"`
	Price   int8   `gorm:"column:price" json:"price"`
}
