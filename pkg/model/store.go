package model

import "time"

// Store is the publish name of the drink store
type Store struct {
	ID        int    `gorm:"id,primary_key" json:"id"`
	GroupName string `gorm:"group_name" json:"group_name"`
	// BranchStores []BranchStore `gorm:"ForeignKey:store_group_id" json:"branch_stores,omit"`
}

// TableName of Store
func (s *Store) TableName() string {
	return "stores"
}

// BranchStore always is one of the Store
type BranchStore struct {
	ID           int       `gorm:"id" json:"id"`
	Name         string    `gorm:"name" json:"name"`
	StoreGroupID string    `gorm:"store_group_id" json:"store_group_id"`
	Address      string    `gorm:"address" json:"address"`
	Size         string    `gorm:"size" json:"size"`
	ImgURI       string    `gorm:"img_uri" json:"img_uri"`
	Phone        string    `gorm:"phone" json:"phone"`
	CreatedAt    time.Time `gorm:"created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"updated_at" json:"updated_at"`
}

// TableName of BranchStore
func (s *BranchStore) TableName() string {
	return "branch_stores"
}
