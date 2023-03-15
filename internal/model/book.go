package model

import (
	"context"
	"time"

	"github.com/falahlaz/boilerplate-golang/internal/abstraction"
	"gorm.io/gorm"
)

type BookEntity struct {
	Title     string `json:"title" validate:"required" gorm:"index:idx_book_title;size:100;not null"`
	Author    string `json:"author" validate:"required" gorm:"index:idx_book_author;size:50;not null"`
	Publisher string `json:"publisher" validate:"required" gorm:"index:idx_book_publisher;size:50;not null"`
	Synopsis  string `json:"synopsis" validate:"required" gorm:"type:text"`
	Year      string `json:"year" validate:"required"`
}

type BookModel struct {
	// abstraction
	abstraction.Model

	// entity
	BookEntity

	// context
	Context context.Context `json:"-" gorm:"-"`
}

func (m *BookModel) TableName() string {
	return "master.books"
}

func (m *BookModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UnixMilli()
	return
}

func (m *BookModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UnixMilli()
	return
}
