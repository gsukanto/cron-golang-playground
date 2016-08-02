package mp_pg

import (
	"testing"
	"time"
)

var (
	id int64 = 1

	ids []int64 = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
)

func TestGetBusinessIdsWhereDailyOrLowIsTrue(t *testing.T) {
	GetBusinessIdsWhereDailyOrLowIsTrue()
}

func TestGetDailyBusinessIdRecipients(t *testing.T) {
	GetDailyBusinessIdRecipients()
}

func TestGetLowInventoryBusinessIds(t *testing.T) {
	GetLowInventoryBusinessIds()
}

func TestGetBusinessProfileByBusinessIds(t *testing.T) {
	GetBusinessProfileByBusinessIds(ids)
}

func TestGetEmailsByBusinessId(t *testing.T) {
	GetEmailsByBusinessId(id)
}

func TestGetLowInventoryAverageDailySales(t *testing.T) {
	GetLowInventoryAverageDailySales(id)
}

func TestGetLimitedListLowInventory(t *testing.T) {
	GetLimitedListLowInventory(id)
}

func TestCountListLowInventory(t *testing.T) {
	CountListLowInventory(id)
}

func TestGetOutletIdsfromBusinessId(t *testing.T) {
	GetOutletIdsfromBusinessId(id)
}

func TestGetPaymentIdsFromBusinessId(t *testing.T) {
	indonesia, _ := time.LoadLocation("Asia/Jakarta")
	timeNow := time.Now().In(indonesia)
	endDate := timeNow.Format("2006-01-02")
	startDate := timeNow.AddDate(0, 0, -1).Format("2006-01-02")

	GetPaymentIdsFromBusinessId(id, startDate, endDate)
}

func TestGetPaymentDatasFromPaymentIds(t *testing.T) {
	GetPaymentDatasFromPaymentIds(ids)
}

func TestGetGrossSalesFromPaymentIds(t *testing.T) {
	GetGrossSalesFromPaymentIds(ids)
}

func TestGetTopItems(t *testing.T) {
	GetTopItems(ids)
}
