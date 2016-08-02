package mp_util

import (
	"testing"
	"time"
)

var (
	arrayStrings []string = []string{"1", "2", "3", "4", "5"}
	arrayInt64   []int64  = []int64{1, 2, 3, 4, 5}

	from    string = "Golang Util tester"
	channel string = "@gregory"
	title   string = "Golang unit test"
	color   string = "good"
	text    string = "any text message"

	businessCount int64 = 5
	emailSent     int64 = 3
	emailFailed   int64 = 2
)

func TestJakartaDateNow(t *testing.T) {
	id, _ := time.LoadLocation("Asia/Jakarta")
	if GetJakartaDateNow() != time.Now().In(id).Format("02/01/2006") {
		t.Errorf("GetJakartaDateNow failed")
	}
}

func TestJakartaWeekdayNow(t *testing.T) {
	id, _ := time.LoadLocation("Asia/Jakarta")
	if GetJakartaWeekdayNow() != time.Now().In(id).Weekday() {
		t.Errorf("GetJakartaWeekdayNow failed")
	}
}

func TestJakartaWeekdayDateNow(t *testing.T) {
	id, _ := time.LoadLocation("Asia/Jakarta")
	if GetJakartaWeekdayDateNow() != time.Now().In(id).Format("Monday, 02/01/2006") {
		t.Errorf("GetJakartaDateNow failed")
	}
}

func TestStringsToInts64(t *testing.T) {
	result := StringsToInts64(arrayStrings)
	for i, v := range result {
		if v != arrayInt64[i] {
			t.Errorf("StringsToInts64 failed at index %v, at value %v", i, v)
		}
	}
}

func TestSendChat(t *testing.T) {
	statusCode := SendChat(from, title, color, channel, text)
	if statusCode != 0 {
		t.Errorf("Send Chat return with error!")
	}
}

func TestSlackLowInventorySummary(t *testing.T) {
	statusCode := SlackLowInventorySummary(businessCount, emailSent, emailFailed)
	if statusCode != 0 {
		t.Errorf("Slack Low Inventory Summary return with error!")
	}
}

func TestSlackDailySalesSummary(t *testing.T) {
	statusCode := SlackDailySalesSummary(businessCount, emailSent, emailFailed)
	if statusCode != 0 {
		t.Errorf("Slack Daily Sales Summary return with error!")
	}
}
