package utils

import (
	"github.com/google/uuid"
	"strings"
)

func RandomString() string {
	newUUID, _ := uuid.NewRandom()
	return strings.Replace(newUUID.String(), "-", "", -1)
}
