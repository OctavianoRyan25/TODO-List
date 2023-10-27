package model

type Category struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Type  string `gorm:"not null" json:"type" form:"type"`
	Color string `gorm:"not null" json:"color" form:"color" `
	GormModel
	Notes []Note //HasMany
}