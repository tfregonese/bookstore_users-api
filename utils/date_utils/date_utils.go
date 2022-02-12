package date_utils

import "time"

const (
	apiDateLayout = time.RFC3339
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
