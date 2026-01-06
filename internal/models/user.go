package models

import (
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	UserTableName = "users"
)

type User struct {
	BaseModel
	LastLoginAt *time.Time `gorm:"index" json:"last_login_at"`

	PhoneNumber       string `json:"phone_number" gorm:"uniqueIndex:idx_phone_dial_code"`
	Password          string `json:"-"`
	CountryID         uint64 `json:"country_id" gorm:"references:id,cascade:delete;"`
	CountryDialCodeID uint64 `json:"country_dial_code_id" gorm:"references:id,cascade:delete;uniqueIndex:idx_phone_dial_code"` // Updated tag here

	Country         Country         `gorm:"foreignKey:CountryID" json:"country"`
	CountryDialCode CountryDialCode `gorm:"foreignKey:CountryDialCodeID" json:"countryDialCode"`

	Profile Profile `gorm:"foreignKey:UserID" json:"profile"`

	EasemobUsername string `json:"easemob_username"`
	EasemobUUID     string `json:"easemob_uuid"`
	EasemobPassword string `json:"easemob_password"`
}

func (u *User) TableName() string {
	return UserTableName
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		if u.Password, err = u.HashPassword(u.Password); err != nil {
			return err
		}
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.Password != "" {
		if !isHashed(u.Password) {
			if u.Password, err = u.HashPassword(u.Password); err != nil {
				return err
			}
		}
	}
	return nil
}

func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func (u *User) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func isHashed(value string) bool {
	return strings.HasPrefix(value, "$2a$") || strings.HasPrefix(value, "$2b$")
}
