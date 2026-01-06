package schemas

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID       `json:"id"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	DeletedAt       *time.Time      `json:"deleted_at"`
	LastLoginAt     *time.Time      `json:"last_login_at"`
	PhoneNumber     string          `json:"phone_number"`
	CountryDialCode CountryDialCode `json:"country_dial_code"`
	Country         Country         `json:"country"`
	EasemobUsername string          `json:"easemob_username"`
	EasemobUUID     string          `json:"easemob_uuid"`
}
