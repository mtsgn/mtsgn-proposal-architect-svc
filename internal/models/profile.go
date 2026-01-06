package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ProfileTableName = "profile"
)

type Profile struct {
	BaseModel
	UserID         string     `json:"user_id" gorm:"column:user_id;type:char(36);not null"`
	Nickname       string     `json:"nickname" gorm:"column:nickname;type:varchar(100)"`
	Gender         string     `json:"gender" gorm:"column:gender;type:varchar(100)"`
	Birthday       *time.Time `json:"birthday" gorm:"column:birthday;type:date"`
	Avatar         string     `json:"avatar" gorm:"column:avatar;type:varchar(100)"`
	Voice          string     `json:"voice" gorm:"column:voice;type:varchar(100)"`
	SelfIntro      string     `json:"self_intro" gorm:"column:self_intro;type:varchar(100)"`
	LikeCount      int        `json:"like_count" gorm:"column:like_count;default:0"`
	FollowingCount int        `json:"following_count" gorm:"column:following_count;default:0"`
	FollowerCount  int        `json:"follower_count" gorm:"column:follower_count;default:0"`
}

func (p *Profile) TableName() string {
	return ProfileTableName
}

func (p *Profile) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}
