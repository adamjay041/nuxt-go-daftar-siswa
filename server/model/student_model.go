package model

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Student struct {
	Id        int        `json:"id" gorm:"primary_key"`
	NPM       string     `json:"npm" gorm:"not null"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	FullName  string     `json:"full_name"`
	Class     string     `json:"class" `
	Email     string     `json:"email" `
	Password  string     `json:"password" gorm:"not null"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy time.Time  `json:"updated_by"`
	CreateBy  time.Time  `json:"create_by"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

type Login struct {
	Email    string
	Password string
}

type TokenJwt struct {
	jwt.StandardClaims
	Password string
}

// type StudentRequest struct {
// 	NPM       int    `json:"npm" gorm:"not null, unique"`
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	FullName  string `json:"full_name"`
// 	Email     string `json:"email" gorm:"not null"`
// 	Class     string `json:"class" gorm:"not null"`
// }

type StudentRepository interface {
	Fetch() (*[]Student, error)
	Create(student *Student) (*Student, error)
	FetchBy(value string) (*[]Student, error)
	DeleteBy(id string) (*Student, error)
	EditBy(filter string, value string, student map[string]interface{}) (*Student, error)
}

type StudentUsecase interface {
	Fecth() (*[]Student, error)
	Create(student *Student) (*Student, error)
	FetchBy(value string) (*[]Student, error)
	DeleteBy(id string) (*Student, error)
	EditBy(filter string, value string, student map[string]interface{}) (*Student, error)
}
