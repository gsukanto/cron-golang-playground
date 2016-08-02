package mp_util

import (
	"time"
)

func GetJakartaTimeNow() time.Time {
	id, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(id)
}

func GetJakartaDateNow() string {
	now := GetJakartaTimeNow()
	return now.Format("02/01/2006")
}

func GetJakartaWeekdayNow() time.Weekday {
	now := GetJakartaTimeNow()
	return now.Weekday()
}

func GetJakartaWeekdayDateNow() string {
	now := GetJakartaTimeNow()
	return now.Format("Monday, 02/01/2006")
}

func GetJakartaWeekdayDateYesterday() string {
	now := GetJakartaTimeNow()
	return now.AddDate(0, 0, -1).Format("Monday, 02/01/2006")
}
