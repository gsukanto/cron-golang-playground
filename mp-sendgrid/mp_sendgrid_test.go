package mp_sendgrid

import (
	"testing"

	mpModel "../mp-model"
)

var (
	emails []string = []string{"gregory@mokapos.com"}

	substitudeLowInventory []mpModel.Substitude = []mpModel.Substitude{
		{"date", "Date"},
		{"business_name", "Business Name"},
		{"business_phone", "Business Phone"},
		{"business_email", "Business Email"},
		{"list_item_variants", "List Item Variants"}}

	substitudeDailySales []mpModel.Substitude = []mpModel.Substitude{
		{"date", "Date"},
		{"business_name", "Business Name"},
		{"business_phone", "Business Phone"},
		{"business_email", "Business Email"},
		{"gross_sales", "Gross Sales"},
		{"discounts", "Discounts"},
		{"refunds", "Refunds"},
		{"net_sales", "Net Sales"},
		{"gratuity", "Gratuity"},
		{"tax", "Tax"},
		{"total_collected", "Total Collected"},
		{"list_of_top_items", "List Of Top Items"}}
)

func TestSendLowInventory(t *testing.T) {
	statusCode := SendLowInventory(substitudeLowInventory, emails)

	if statusCode < 200 || statusCode > 299 {
		t.Errorf("Send low inventory giving status code %v", statusCode)
	}
}

func TestSendDailySales(t *testing.T) {
	statusCode := SendDailySales(substitudeDailySales, emails)

	if statusCode < 200 || statusCode > 299 {
		t.Errorf("Send daily summary giving status code %v", statusCode)
	}
}
