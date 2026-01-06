package utils

import (
	"math/rand"
	"time"
)

func GenerateDefaultProfile() (*string, string, *time.Time) {
	genders := []string{"male", "female", "other"}
	nicknames := []string{"User", "Guest", "Anonymous"}

	rand.Seed(time.Now().UnixNano())
	gender := genders[rand.Intn(len(genders))]
	nickname := nicknames[rand.Intn(len(nicknames))] + "_" + generateRandomString(6)
	birthday := time.Now().AddDate(-20, -rand.Intn(12), -rand.Intn(365))

	return &gender, nickname, &birthday
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
