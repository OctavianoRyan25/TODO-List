package model

import "time"

type GormModel struct {
	CreatedAt *time.Time 	`json:"created_at"`
	UpdatedAt time.Time 	`json:"updated_at"`
}