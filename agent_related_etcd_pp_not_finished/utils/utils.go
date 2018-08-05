package utils

import (
	"time"
)

func GetCurrentTimeStr() string {
	return time.Now().Format(time.RFC3339)
}
