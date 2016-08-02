package main

import (
	"fmt"
	"time"

	mpHelper "./mp-helper"
	mpModel "./mp-model"
	mpPg "./mp-pg"
	mpRedis "./mp-redis"
	mpSendgrid "./mp-sendgrid"
	mpUtil "./mp-util"
)

// TODO: use proper html and css

func makeLowInventoryData(listLowInventory []mpModel.ItemVariantData) []mpModel.ItemVariantData {
	for i, v := range listLowInventory {
		listLowInventory[i].AvgDailySales = mpPg.GetLowInventoryAverageDailySales(v.ItemVariantId)
	}

	return listLowInventory
}

func makeListLowInventorySubstitude(substitudes []mpModel.Substitude, businessId int64, count int64) []mpModel.Substitude {
	var html string
	lowInventory := mpPg.GetLimitedListLowInventory(businessId)
	listLowInventory := makeLowInventoryData(lowInventory)

	for _, v := range listLowInventory {
		tdItemName := fmt.Sprintf("<td %s'>%s</td>",
			tdItemNameStyling,
			mpUtil.TruncateString(v.ItemName+" "+v.ItemVariantName, mpModel.LimitNameLength))

		tdOutletName := fmt.Sprintf("<td %s'>%s</td>",
			tdOutletNameStyling,
			mpUtil.TruncateString(v.OutletName, mpModel.LimitNameLength))

		tdInStock := fmt.Sprintf("<td %s'>%v</td>",
			tdInStockStyling,
			v.InStock)

		tdAvgDailySales := fmt.Sprintf("<td %s'>%v</td>",
			tdAvgDailySalesStyling,
			fmt.Sprintf("%.2f", v.AvgDailySales))

		tableInventory := fmt.Sprintf("<tr>%s%s%s%s</tr>",
			tdItemName, tdOutletName, tdInStock, tdAvgDailySales)

		html = fmt.Sprintf("%s%s", html, tableInventory)
	}

	if count > mpModel.LimitListLowInventory {
		unrenderedItemVariants := fmt.Sprintf("<tr><td %s' colspan='4'><strong><em>( %v more items )</em></strong></td></tr>",
			unrenderedItemVariantsStyling,
			count-mpModel.LimitListLowInventory)

		html = fmt.Sprintf("%s%s", html, unrenderedItemVariants)
	}

	substitudes = append(substitudes, mpModel.Substitude{"list_item_variants", html})

	return substitudes
}

func makeLowInventorySubstitute(count int64, businessData mpModel.BusinessDataRecipients) []mpModel.Substitude {
	var substitudes []mpModel.Substitude

	substitudes = mpHelper.MakeTimeSubstitude(substitudes)
	substitudes = mpHelper.MakeBusinessDataSubstitude(substitudes, businessData)
	substitudes = makeListLowInventorySubstitude(substitudes, businessData.BusinessId, count)

	return substitudes
}

func setEmailRedisStatus(statusCode int, businessId int64) {
	if mpHelper.IsHttpSucceed(statusCode) {
		mpRedis.SetLowBusinessIdEmailSucceed(businessId)
	} else {
		mpRedis.SetLowBusinessIdEmailFailed(businessId)
	}
}

func sendLowInventory(businessId int64) {
	CountListLowInventory := mpPg.CountListLowInventory(businessId)
	if CountListLowInventory == 0 {
		return
	}

	businessData := mpRedis.GetBusinessEmailsByBusinessIdProfile(businessId)
	substitudes := makeLowInventorySubstitute(CountListLowInventory, businessData)

	statusCode := mpSendgrid.SendLowInventory(substitudes, businessData.Emails)
	setEmailRedisStatus(statusCode, businessId)
}

// Special request from our VP, please do not use this function except for operational
func SendLowInventoryToSpecificEmails(businessId int64, emails []string) {
	CountListLowInventory := mpPg.CountListLowInventory(businessId)
	if CountListLowInventory == 0 {
		return
	}

	businessData := mpRedis.GetBusinessEmailsByBusinessIdProfile(businessId)
	substitudes := makeLowInventorySubstitute(CountListLowInventory, businessData)

	mpSendgrid.SendLowInventory(substitudes, emails)
}

func main() {
	businessIds := mpRedis.GetLowInventoryBusinessIds()

	mpRedis.DelLowBusinessIdEmailSucceed()
	mpRedis.DelLowBusinessIdEmailFailed()

	for i, v := range businessIds {
		if (i % mpModel.LowInventoryThread) == 0 {
			sendLowInventory(v)
		} else {
			go sendLowInventory(v)
		}
	}

	for {
		if mpHelper.IsTotalEql(mpRedis.CountLowBusinessIdEmailSucceed(),
			mpRedis.CountLowBusinessIdEmailFailed(),
			mpRedis.CountListLowInventoryBusinessIds()) {

			mpUtil.SlackLowInventorySummary(mpRedis.CountListLowInventoryBusinessIds(),
				mpRedis.CountLowBusinessIdEmailSucceed(), mpRedis.CountLowBusinessIdEmailFailed())
			break
		}

		time.Sleep(mpModel.CronSleepTime)
	}
}

var tdItemNameStyling = `style='border: none;text-align: left;padding: 8px; color:black;  -webkit-line-clamp: $lines-to-show;-webkit-box-orient: vertical;overflow: hidden;text-overflow: ellipsis;`
var tdOutletNameStyling = `style='border: none;text-align: left;padding: 8px; color: #969696;`
var tdInStockStyling = `style='border: none;text-align: center;padding: 8px; color: #969696;`
var tdAvgDailySalesStyling = `style='border: none;text-align: center;padding: 8px; color: #969696;`
var unrenderedItemVariantsStyling = `style='border: none;text-align: left;padding: 8px;`
