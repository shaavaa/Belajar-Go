package dao

import (
	"base-gin/app/domain"
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	AccountID *uint              `gorm:"uniqueIndex;"`//kalo ada * nya boleh null, artinya user tidak harus memiliki akun untuk melihat
	Account   *Account           `gorm:"foreignKey:AccountID;"`
	Fullname  string             `gorm:"size:56;not null;"`
	Gender    *domain.TypeGender `gorm:"type:enum('f','m');"`
	BirthDate *time.Time
}

func (Person) TableName() string {
	return "persons"
}
