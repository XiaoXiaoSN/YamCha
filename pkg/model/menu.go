package model

import (
	"encoding/json"
	"time"
)

// Price Object
// type Price struct {
// 	Small string `gorm:"small" json:"small"`
// 	Large string `gorm:"large" json:"large"`
// }

// Menu Object
type Menu struct {
	ID      int             `gorm:"id" json:"id"`
	Name    string          `gorm:"name" json:"name"`
	StoreID string          `gorm:"store_id" json:"store_id"`
	Price   json.RawMessage `gorm:"price" json:"price"`
	// PriceStruct Price
	// Size      string    `gorm:"size" json:"size"`
	ImgURI    string    `gorm:"img_uri" json:"img_uri"`
	CreatedAt time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at" json:"updated_at"`
}

// TableName of Menu
func (*Menu) TableName() string {
	return "menus"
}
