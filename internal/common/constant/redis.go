package constant

import (
	"fmt"
	"time"
)

const (
	DefaultExpiration = time.Hour * 24
)

type Message struct {
	Event EventType `json:"event"`
	Data  Data      `json:"data"`
}

type Data struct {
	RoomID string `json:"room_id"`
	UserID string `json:"user_id,omitempty"`
	Reason string `json:"reason,omitempty"`
	Ts     string `json:"ts"`
}

type EventType string

const (
	JoinRoomEvent   EventType = "join_room"
	LeaveRoomEvent  EventType = "leave_room"
	ErrorEvent      EventType = "error"
	DisconnectEvent EventType = "disconnect"
	DeleteRoomEvent EventType = "delete_room"
)

func (e EventType) String() string {
	return string(e)
}

// Constants for Redis keys
func RoomUsersKey(roomID string) string {
	return fmt.Sprintf("room:%s:users", roomID)
}
func RoomEventsKey(roomID string) string {
	return fmt.Sprintf("room:%s:events", roomID)
}

func UserKey(userID string) string {
	return fmt.Sprintf("user:%s", userID)
}

func ProfileKey(userID string) string {
	return fmt.Sprintf("profile:%s", userID)
}
