package utils

import uuid "github.com/satori/go.uuid"

func GetUuid() string {
	v4 := uuid.NewV4()
	return v4.String()
}
