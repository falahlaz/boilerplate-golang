package model

import (
	"context"
	"time"

	"github.com/falahlaz/boilerplate-golang/internal/abstraction"
	"gorm.io/gorm"
)

type BorrowEntity struct {
	UserID int `json:"user_id" gorm:"not null"`
	BookID int `json:"book_id" gorm:"not null"`
}

type BorrowModel struct {
	// abstraction
	abstraction.Model

	// entity
	BorrowEntity

	// relation
	User *UserModel `gorm:"foreignKey:UserID"`
	Book *BookModel `gorm:"foreignKey:BookID"`

	// context
	Context context.Context `json:"-" gorm:"-"`
}

func (m *BorrowModel) TableName() string {
	return "transaction.borrows"
}

func (m *BorrowModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UnixMilli()
	return
}

func (m *BorrowModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UnixMilli()
	return
}
