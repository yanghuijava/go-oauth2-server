package myuuid

import (
	"github.com/google/uuid"
	"strings"
)

func SimpleUUID() string {
	uuid := uuid.New().String()
	return strings.Replace(uuid, "-", "", -1)
}

func UUID() string {
	return uuid.New().String()
}
