package model

import (
	"encoding/json"
	"time"
)

// Price Object
// type Price struct {
// 	Small string `gorm:"column:small" json:"small"`
// 	Large string `gorm:"column:large" json:"large"`
// }

// Menu Object
type Menu struct {
	ID      int             `gorm:"column:id" json:"id"`
	Name    string          `gorm:"column:name" json:"name"`
	StoreID string          `gorm:"column:store_id" json:"store_id"`
	Price   json.RawMessage `gorm:"column:price" json:"price"`
	// PriceStruct Price
	// Size      string    `gorm:"column:size" json:"size"`
	ImgURI    string    `gorm:"column:img_uri" json:"img_uri"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName of Menu
func (*Menu) TableName() string {
	return "menus"
}
