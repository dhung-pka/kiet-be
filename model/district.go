package model

import "time"

type District struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Code string `json:"code"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}
