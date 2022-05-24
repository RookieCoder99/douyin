package utils

import "github.com/gofrs/uuid"

func GetUUID() string {
	u2, _ := uuid.NewV4()
	return u2.String()
}
