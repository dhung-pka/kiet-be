package model

import "time"

type Credential struct {
	Id        uint    `json:"id" gorm:"primaryKey"`
	Username  string  `json:"username" gorm:"unique"`
	Password  string  `json:"-"`
	Role      string  `json:"role"`
	ProfileId uint    `json:"profileId"`
	Profile   Profile `json:"profile" gorm:"foreignKey:ProfileId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`

	CreatedAt time.Time `json:"createdAt" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updatedAt" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}
