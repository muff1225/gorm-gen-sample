// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID            int32          `gorm:"column:id;type:int unsigned;primaryKey;autoIncrement:true" json:"id"`
	Name          string         `gorm:"column:name;type:varchar(255);not null" json:"name"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3)" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3)" json:"deleted_at"`
	PurchaseOrder PurchaseOrder  `json:"purchase_order"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}