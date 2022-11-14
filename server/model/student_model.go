package model

import (
	"server/dto"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type Student struct {
	Id         int64           `json:"id" gorm:"primary_key;auto_increment;not_null"`
	Npm        string          `json:"npm" gorm:"type:varchar;not null"`
	Name       string          `json:"name" gorm:"type:varchar;not null"`
	Kelas      string          `json:"kelas" gorm:"type:varchar;not null"`
	Password   string          `json:"password" gorm:"type:varchar;not null"`
	Periode    string          `json:"periode" gorm:"type:varchar;not null"`
	L_pict     string          `json:"l_pict" gorm:"type:varchar"`
	Status     string          `json:"status" gorm:"type:varchar"`
	Generation string          `json:"generation" gorm:"type:varchar"`
	Email      string          `json:"email" gorm:"type:varchar;not null;"`
	Role       string          `json:"role" gorm:"type:varchar;not null"`
	Major      string          `json:"major" gorm:"type:varchar;not null"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	UpdatedBy  time.Time       `json:"updated_by"`
	CreateBy   time.Time       `json:"create_by"`
	DeletedAt  *gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

type Login struct {
	Email    string
	Password string
}

type TokenJwt struct {
	jwt.StandardClaims
	Password string
}

type StudentRepository interface {
	Fetch() (*[]Student, error)
	Create(student *dto.CreateRequestStudent) (*dto.CreateRequestStudent, error)
	FetchBy(value string) (*[]Student, error)
	DeleteBy(id string) (*Student, error)
	EditBy(filter string, value string, student *dto.EditRequestStudent) (*Student, error)
}

type StudentUsecase interface {
	Fecth() (*[]Student, error)
	Create(student *dto.CreateRequestStudent) (*dto.CreateRequestStudent, error)
	FetchBy(value string) (*[]Student, error)
	DeleteBy(id string) (*Student, error)
	EditBy(filter string, value string, student *dto.EditRequestStudent) (*Student, error)
}
