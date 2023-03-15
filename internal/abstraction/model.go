package abstraction

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Model struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;" param:"id"`

	CreatedAt int64   `json:"created_at"`
	CreatedBy *string `json:"created_by"`
	UpdatedAt int64   `json:"updated_at"`
	UpdatedBy *string `json:"updated_by"`

	DeletedAt soft_delete.DeletedAt `json:"deleted_at"`
}

func (m *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UnixMilli()
	return
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UnixMilli()
	return
}
