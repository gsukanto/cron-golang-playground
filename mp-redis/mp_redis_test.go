package mp_redis

import (
	"testing"
)

var (
	ids    []int64  = []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	emails []string = []string{"email1", "email2", "email3"}

	businessName  = "Nama bisnis"
	businessPhone = "0812345679"
	businessEmail = "business@email.com"
)

func TestAllBusinessRecipient(t *testing.T) {
	DelAllBusinessIdsRecipient()
	SetAllBusinessIdsRecipient(ids)
	allBusinessRecipientIds := GetAllBusinessIdsRecipient()

	for i, v := range allBusinessRecipientIds {
		if v != ids[i] {
			t.Errorf("AllBusinessRecipient failed in index %v, at value %v", i, v)
		}
	}
}

func TestBusinessProfile(t *testing.T) {
	for _, id := range ids {
		SetBusinessProfileByBusinessId(id, businessName, businessPhone, businessEmail)
		gettedProfiles := GetBusinessProfileByBusinessIds(id)

		if gettedProfiles[0] != businessName {
			t.Errorf("BusinessProfile failed in business name")
		}
		if gettedProfiles[1] != businessPhone {
			t.Errorf("BusinessProfile failed in business phone")
		}
		if gettedProfiles[2] != businessEmail {
			t.Errorf("BusinessProfile failed in business email")
		}
	}
}

func TestBusinessEmail(t *testing.T) {
	for _, id := range ids {
		SetBusinessEmailsByBusinessId(id, emails)
		gettedEmails := GetBusinessEmailsByBusinessId(id)
		for i, v := range gettedEmails {
			if v != emails[i] {
				t.Errorf("BusinessEmail failed in index %v with string %s", i, v)
			}
		}
	}
}

func TestBusinessEmailProfile(t *testing.T) {
	for _, id := range ids {
		SetBusinessProfileByBusinessId(id, businessName, businessPhone, businessEmail)
		SetBusinessEmailsByBusinessId(id, emails)
		gettedBusinessDatas := GetBusinessEmailsByBusinessIdProfile(id)

		if gettedBusinessDatas.BusinessId != id {
			t.Errorf("BusinessEmailProfile failed in with id %v", gettedBusinessDatas.BusinessId)
		}
		if gettedBusinessDatas.Name != businessName {
			t.Errorf("BusinessEmailProfile failed in with name %s", gettedBusinessDatas.Name)
		}
		if gettedBusinessDatas.Phone != businessPhone {
			t.Errorf("BusinessEmailProfile failed in with phone %s", gettedBusinessDatas.Phone)
		}
		if gettedBusinessDatas.Email != businessEmail {
			t.Errorf("BusinessEmailProfile failed in with business email %s", gettedBusinessDatas.Email)
		}
		for i, email := range gettedBusinessDatas.Emails {
			if email != emails[i] {
				t.Errorf("BusinessEmailProfile failed in index %v with email %s", i, email)
			}
		}
	}
}

func TestLowInventoryBusinessIds(t *testing.T) {
	DelLowInventoryBusinessIds()
	SetLowInventoryBusinessIds(ids)
	lowBusinessIds := GetLowInventoryBusinessIds()

	for i, v := range lowBusinessIds {
		if v != ids[i] {
			t.Errorf("Low business id failed in index %v, at value %v", i, v)
		}
	}

	if int64(len(lowBusinessIds)) != CountListLowInventoryBusinessIds() {
		t.Errorf("Count low business ids failed at %v", CountListLowInventoryBusinessIds())
	}
}

func TestLowBusinessIdEmailSucceed(t *testing.T) {
	DelLowBusinessIdEmailSucceed()

	for _, id := range ids {
		SetLowBusinessIdEmailSucceed(id)
		gettedSucceedEmailIds := GetLowBusinessIdsEmailSucceed()
		for i, v := range gettedSucceedEmailIds {
			if v != ids[i] {
				t.Errorf("Low business ids succeed email failed in index %v with string %s", i, v)
			}
		}
		if int64(len(gettedSucceedEmailIds)) != CountLowBusinessIdEmailSucceed() {
			t.Errorf("Low business ids succeed email failed at %v", CountLowBusinessIdEmailSucceed())
		}
	}
}

func TestLowBusinessIdEmailFailed(t *testing.T) {
	DelLowBusinessIdEmailFailed()

	for _, id := range ids {
		SetLowBusinessIdEmailFailed(id)
		gettedFailedEmailIds := GetLowBusinessIdsEmailFailed()
		for i, v := range gettedFailedEmailIds {
			if v != ids[i] {
				t.Errorf("Low business ids failed email failed in index %v with string %s", i, v)
			}
		}
		if int64(len(gettedFailedEmailIds)) != CountLowBusinessIdEmailFailed() {
			t.Errorf("Low business ids failed email failed at %v", CountLowBusinessIdEmailFailed())
		}
	}
}

func TestDailySalesBusinessId(t *testing.T) {
	DelDailySalesBusinessId()
	SetDailySalesBusinessIds(ids)
	gettedFailedEmailIds := GetDailySalesBusinessIds()
	for i, v := range gettedFailedEmailIds {
		if v != ids[i] {
			t.Errorf("Daily summary business ids failed in index %v with string %s", i, v)
		}
	}
	if int64(len(gettedFailedEmailIds)) != CountDailySalesBusinessIds() {
		t.Errorf("Daily summary business ids failed at %v", CountDailySalesBusinessIds())
	}
}

func TestDailyBusinessIdEmailSucceed(t *testing.T) {
	DelDailyBusinessIdsEmailSucceed()

	for _, id := range ids {
		SetDailyBusinessIdEmailSucceed(id)
		gettedFailedEmailIds := GetDailyBusinessIdsEmailSucceed()
		for i, v := range gettedFailedEmailIds {
			if v != ids[i] {
				t.Errorf("Daily business ids succeed email failed in index %v with string %s", i, v)
			}
		}
		if int64(len(gettedFailedEmailIds)) != CountDailyBusinessIdsEmailSucceed() {
			t.Errorf("Daily business ids succeed email failed at %v", CountDailyBusinessIdsEmailSucceed())
		}
	}
}

func TestDailyBusinessIdEmailFailed(t *testing.T) {
	DelDailyBusinessIdsEmailFailed()

	for _, id := range ids {
		SetDailyBusinessIdEmailFailed(id)
		gettedFailedEmailIds := GetDailyBusinessIdsEmailFailed()
		for i, v := range gettedFailedEmailIds {
			if v != ids[i] {
				t.Errorf("Daily business ids failed email failed in index %v with string %s", i, v)
			}
		}
		if int64(len(gettedFailedEmailIds)) != CountDailyBusinessIdsEmailFailed() {
			t.Errorf("Daily business ids failed email failed at %v", CountDailyBusinessIdsEmailFailed())
		}
	}
}
