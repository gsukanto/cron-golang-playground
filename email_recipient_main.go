package main

import (
	mpModel "./mp-model"
	mpPg "./mp-pg"
	mpRedis "./mp-redis"

	"github.com/asaskevich/govalidator"
)

func setAllBusinessIdsRecipient() {
	mpRedis.DelAllBusinessIdsRecipient()
	businessIds := mpPg.GetBusinessIdsWhereDailyOrLowIsTrue()
	mpRedis.SetAllBusinessIdsRecipient(businessIds)
}

func setDailySalesBusinessIds() {
	mpRedis.DelDailySalesBusinessId()
	businessIds := mpPg.GetDailyBusinessIdRecipients()
	mpRedis.SetDailySalesBusinessIds(businessIds)
}

func setLowInventoryBusinessIds() {
	mpRedis.DelLowInventoryBusinessIds()
	businessIds := mpPg.GetLowInventoryBusinessIds()
	mpRedis.SetLowInventoryBusinessIds(businessIds)
}

func setBusinessDataProfileByBusinessId() {
	businessIds := mpRedis.GetAllBusinessIdsRecipient()

	var businessProfiles []mpModel.BusinessProfile
	businessProfiles = mpPg.GetBusinessProfileByBusinessIds(businessIds)

	for _, v := range businessProfiles {
		tempEmails := mpPg.GetEmailsByBusinessId(v.BusinessId)
		var emails []string

		for _, recipientEmail := range tempEmails {
			if govalidator.IsEmail(recipientEmail) {
				emails = append(emails, recipientEmail)
			}
		}

		if len(emails) == 0 {
			emails = []string{v.Email}
		}

		mpRedis.SetBusinessProfileByBusinessId(v.BusinessId, v.Name, v.Phone, v.Email)
		mpRedis.SetBusinessEmailsByBusinessId(v.BusinessId, emails)
	}
}

func main() {
	setAllBusinessIdsRecipient()
	setDailySalesBusinessIds()
	setLowInventoryBusinessIds()
	setBusinessDataProfileByBusinessId()
}
