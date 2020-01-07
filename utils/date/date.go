package date

import (
	"time"
)

const (
	LAYOUT    = "2006-01-02T15:04:05.000Z"
	LAYOUT_DB = "2006-01-02 15:04:05"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(LAYOUT)
}

func GetNowDB() string {
	return GetNow().Format(LAYOUT_DB)
}
