package models

import (
	"boilerplate-api/internal/schemas"
	"time"
)

type Country struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Name      string    `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Code      string    `json:"code" gorm:"column:iso_code;type:char(3);not null"`
}

func (c *Country) TableName() string {
	return "countries"
}

func (c *Country) ToSchema() schemas.Country {
	return schemas.Country{
		ID:   c.ID,
		Name: c.Name,
		Code: c.Code,
	}
}

type CountryDialCode struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	DialCode  string    `json:"dial_code" gorm:"column:dial_code;type:varchar(10);not null"`
	CountryID uint64    `json:"country_id" gorm:"column:country_id;not null"`
	Country   Country   `gorm:"foreignKey:CountryID" json:"country"`
}

func (c *CountryDialCode) TableName() string {
	return "country_dial_codes"
}

func (c *CountryDialCode) ToSchema() schemas.CountryDialCode {
	return schemas.CountryDialCode{
		ID:       c.ID,
		DialCode: c.DialCode,
		Country:  c.Country.ToSchema(),
	}
}
