package utils

import (
	"errors"
	"strings"
)

func RemoveFromSlice(slice []string, item string) []string {
	for i, v := range slice {
		if v == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func ConvKeyValue(field string) (string, string, error) {
	fieldArr := strings.Split(field, "_")
	if len(fieldArr) != 2 {
		return "", "", errors.New("field must be in the format of field_value")
	}
	return fieldArr[0], fieldArr[1], nil
}

func GetValueSort(sort string) int {
	if sort == "desc" {
		return -1
	}
	return 1
}
