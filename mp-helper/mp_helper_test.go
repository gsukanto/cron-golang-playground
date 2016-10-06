package mp_helper

import (
	"testing"
	"time"

	mpModel "../mp-model"
)

var (
	equal    []int64 = []int64{3, 2, 5}
	notEqual []int64 = []int64{3, 2, 6}

	httpInformational int = 100
	httpSuccess       int = 200
	httpRedirection   int = 300
	httpClientError   int = 400
	httpServerError   int = 500

	businessData mpModel.BusinessDataRecipients = mpModel.BusinessDataRecipients{
		1, "businessName", "businessPhoneNumber", "businessEmail", []string{}}
)

func TestIsTotalEql(t *testing.T) {
	if !IsTotalEql(equal[0], equal[1], equal[2]) {
		t.Errorf("Total equal checker failed at the value is equal")
	}
	if IsTotalEql(notEqual[0], notEqual[1], notEqual[2]) {
		t.Errorf("Total equal checker failed at the value is not equal")
	}
}

func TestIsHttpSucceed(t *testing.T) {
	if IsHttpSucceed(httpInformational) {
		t.Errorf("Http succeed checker failed at informational status code")
	}
	if !IsHttpSucceed(httpSuccess) {
		t.Errorf("Http succeed checker failed at success status code")
	}
	if IsHttpSucceed(httpRedirection) {
		t.Errorf("Http succeed checker failed at redirection status code")
	}
	if IsHttpSucceed(httpClientError) {
		t.Errorf("Http succeed checker failed at client error status code")
	}
	if IsHttpSucceed(httpServerError) {
		t.Errorf("Http succeed checker failed at server error status code")
	}
}

func TestMakeTimeSubstitude(t *testing.T) {
	var substitudes []mpModel.Substitude
	id, _ := time.LoadLocation("Asia/Jakarta")
	substitudes = MakeTimeSubstitude(substitudes)
	if substitudes[0].Value != time.Now().In(id).AddDate(0, 0, -1).Format("Monday, 02/01/2006") {
		t.Errorf("Make time substitudes failed with %v", substitudes)
	}
}

func TestBusinessDataSubstitude(t *testing.T) {
	var substitudes []mpModel.Substitude
	substitudes = MakeBusinessDataSubstitude(substitudes, businessData)
	if substitudes[0].Value != businessData.Name {
		t.Errorf("Make business data substitudes failed at name with %v", substitudes)
	}
	if substitudes[1].Value != businessData.Phone {
		t.Errorf("Make business data substitudes failed at phone with %v", substitudes)
	}
	if substitudes[2].Value != businessData.Email {
		t.Errorf("Make business data substitudes failed at email with %v", substitudes)
	}
}
