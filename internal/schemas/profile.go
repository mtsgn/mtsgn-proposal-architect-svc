package schemas

import (
	"time"

	"github.com/google/uuid"
)

type GetProfileRes struct {
	ID             uuid.UUID  `json:"id"`
	UserID         string     `json:"user_id"`
	Nickname       string     `json:"nickname"`
	Gender         string     `json:"gender"`
	Birthday       *time.Time `json:"birthday"`
	Avatar         string     `json:"avatar"`
	Voice          string     `json:"voice"`
	SelfIntro      string     `json:"self_intro"`
	LikeCount      int        `json:"like_count"`
	FollowingCount int        `json:"following_count"`
	FollowerCount  int        `json:"follower_count"`
}
