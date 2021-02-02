package model

import "time"

// Store is the publish name of the drink store
type Store struct {
	ID           int           `gorm:"column:id;primary_key" json:"id"`
	GroupName    string        `gorm:"column:group_name" json:"group_name"`
	BranchStores []BranchStore `gorm:"foreignKey:store_group_id;references:id" json:"branch_stores,omit"`
}

// TableName of Store
func (s *Store) TableName() string {
	return "stores"
}

// BranchStore always is one of the Store
type BranchStore struct {
	ID           int       `gorm:"column:id" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	StoreGroupID int       `gorm:"column:store_group_id" json:"store_group_id"`
	Address      string    `gorm:"column:address" json:"address"`
	Size         string    `gorm:"column:size" json:"size"`
	ImgURI       string    `gorm:"column:img_uri" json:"img_uri"`
	Phone        string    `gorm:"column:phone" json:"phone"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName of BranchStore
func (s *BranchStore) TableName() string {
	return "branch_stores"
}
