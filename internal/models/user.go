package models

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var JWT_KEY = os.Getenv("JWT_SECRET_KEY")

type User struct {
	ID        uint       `gorm:"primary_key;"`
	Name      string     `json:"name" xml:"name" form:"name" query:"name"`
	Email     string     `json:"email" xml:"email" form:"email" query:"email"`
	Password  string     `json:"password" xml:"password" form:"password" query:"password"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	DeletedAt *time.Time `gorm:"autoDeleteTime"`
}

func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.Password = string(bytes)
}

func (u *User) CheckPasswordHash(password, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password))
	return err == nil
}

// GenerateToken is a method for struct User for creating new jwt token
func (u *User) GenerateToken() (string, error) {
	// var (
	// 	jwtKey = os.Getenv("JWT_SECRET_KEY")
	// )

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
		"name":  u.Name,
		"exp":   time.Now().Add(time.Hour * 72).Unix(), // we set expired in 72 hour
	})

	tokenString, err := token.SignedString([]byte(JWT_KEY))
	return tokenString, err
}
