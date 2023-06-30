package model

import "time"

type Commune struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	DistrictCode string `json:"districtCode"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}
