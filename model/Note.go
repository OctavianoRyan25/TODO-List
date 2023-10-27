package model

import "time"

type Note struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Title  string `gorm:"not null" json:"title" form:"title" validate:"required"`
	Slug   string `gorm:"not null" json:"slug" form:"slug" validate:"required"`
	Body   string `gorm:"not null" json:"body" form:"body" validate:"required"`
	UserID uint
	User   *UserResponse
	GormModel
}

type NoteResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title" form:"title"`
	Slug      string `json:"slug" form:"slug"`
	Body      string `json:"body" form:"body"`
	UserID    uint
	UpdatedAt time.Time `json:"updated_at"`
}

func (NoteResponse) TableName() string {
	return "notes"
}