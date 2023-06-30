package model

import "time"

type Profile struct {
	Id               uint   `json:"id" gorm:"primaryKey"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	Age              int    `json:"age"`
	PhoneNumber      string `json:"phoneNumber"`
	OtherPhoneNumber string `json:"otherPhoneNumber"`
	Email            string `json:"email"`
	DetailAddress    string `json:"detailAddress"`
	DistrictCode     string `json:"districtCode"`
	ProvinceCode     string `json:"provinceCode"`

	CreatedAt time.Time `json:"-" swaggerignore:"true"`
	UpdatedAt time.Time `json:"-" swaggerignore:"true"`
	DeletedAt time.Time `json:"-" swaggerignore:"true"`
}
