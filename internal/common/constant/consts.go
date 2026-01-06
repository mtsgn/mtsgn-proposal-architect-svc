package constant

import "fmt"

type StatusValue string

const (
	StatusActive   StatusValue = "ACTIVE"
	StatusInactive StatusValue = "INACTIVE"
)

type GenderValue string

const (
	GenderMale   GenderValue = "MALE"
	GenderFemale GenderValue = "FEMALE"
)

func (g GenderValue) IsValid() bool {
	return g == GenderMale || g == GenderFemale
}

func (g GenderValue) String() string {
	return string(g)
}

type StatusPostValue string

const (
	StatusPostPublic  StatusPostValue = "PUBLIC"
	StatusPostPrivate StatusPostValue = "PRIVATE"
	StatusPostFriend  StatusPostValue = "FRIEND"
)

func (s StatusPostValue) IsValid() bool {
	return s == StatusPostPublic || s == StatusPostPrivate || s == StatusPostFriend
}

func GetUserRedis(userID string) string {
	return fmt.Sprintf("user:%s", userID)
}
