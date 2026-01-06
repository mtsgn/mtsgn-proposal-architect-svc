package utils

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
)

func GenerateID() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := make([]byte, length)

	for i := range id {
		id[i] = charset[r.Intn(len(charset))]
	}

	return string(id)
}

func GenerateUUID() string {
	return uuid.New().String()
}

func CurrentTS() time.Time {
	return time.Now()
}
