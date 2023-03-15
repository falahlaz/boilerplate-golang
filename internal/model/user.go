package model

import (
	"context"
	"os"
	"time"

	"github.com/falahlaz/boilerplate-golang/internal/abstraction"
	"github.com/falahlaz/boilerplate-golang/pkg/config"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserEntity struct {
	Name         string `json:"name" validate:"required" gorm:"size:50;not null"`
	Email        string `json:"email" validate:"required,email" gorm:"index:idx_user_email;unique;size:150;not null"`
	PasswordHash string `json:"-"`
	Password     string `json:"password" validate:"required" gorm:"-"`
}

type UserModel struct {
	// abstraction
	abstraction.Model

	// entity
	UserEntity

	// context
	Context context.Context `json:"-" gorm:"-"`
}

func (m *UserModel) TableName() string {
	return "master.users"
}

func (m *UserModel) BeforeUpdate(tx *gorm.DB) (err error) {
	m.UpdatedAt = time.Now().UnixMilli()
	return
}

func (m *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = time.Now().UnixMilli()
	return
}

func (m *UserModel) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	m.PasswordHash = string(bytes)
}

func (m *UserModel) GenerateToken() (string, error) {
	var (
		jwtKey = os.Getenv(config.Config.JWT.Secret)
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    m.ID,
		"email": m.Email,
		"name":  m.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}
