package model

import (
	"github.com/OctavianoRyan25/TODO-List/helper"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null" json:"username" form:"username"`
	Email    string `gorm:"not null" json:"email" form:"email" `
	Password string `gorm:"not null" json:"password" form:"password min=8"`
	GormModel
	Notes []Note //Has Many
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username" form:"username"`
	Email    string `json:"email" form:"email" `
}

func (UserResponse) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error){
	validate := validator.New()
	errCreate := validate.Struct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helper.HashPass(u.Password)
	err = nil
	return
}
