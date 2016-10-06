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

func appendSubstitude(substitudes []mpModel.Substitude, key string, value string) []mpModel.Substitude {
	return append(substitudes, mpModel.Substitude{key, value})
}

// TODO: use proper html and css
var tdIndexStyling = `style='border: none;text-align: left;padding: 8px;-webkit-text-size-adjust: 100%;-ms-text-size-adjust: 100%;mso-table-lspace: 0pt;mso-table-rspace: 0pt;'`
var tdItemNameStyling = `style='border: none;text-align: left;padding: 8px;-webkit-text-size-adjust: 100%;-ms-text-size-adjust: 100%;mso-table-lspace: 0pt;mso-table-rspace: 0pt;'`
var tdRevenueStyling = `style='border: none;text-align: left;padding: 8px;color: #969696;-webkit-text-size-adjust: 100%;-ms-text-size-adjust: 100%;mso-table-lspace: 0pt;mso-table-rspace: 0pt;'`
var tdQuantityStyling = `style='border: none;text-align: center;padding: 8px;color: #969696;-webkit-text-size-adjust: 100%;-ms-text-size-adjust: 100%;mso-table-lspace: 0pt;mso-table-rspace: 0pt;'`

func makeTopItemsSubstitude(substitudes []mpModel.Substitude, paymentIds []int64) []mpModel.Substitude {
	var html string
	topItems := mpPg.GetTopItems(paymentIds)

	for i, v := range topItems {
		if i > 2 {
			break
		}
		tdIndex := fmt.Sprintf("<td %s'>%d.</td>",
			tdIndexStyling,
			i+1)

		tdItemName := fmt.Sprintf("<td %s'>%s</td>",
			mpUtil.TruncateString(tdItemNameStyling, mpModel.LimitNameLength),
			v.ItemName)

		tdRevenue := fmt.Sprintf("<td %s'>%v</td>",
			tdRevenueStyling,
			mpUtil.Int64ToRupiah(v.Revenue))

		tdQuantity := fmt.Sprintf("<td %s'>%v</td>",
			tdQuantityStyling,
			v.Quantity)

		tableTopItems := fmt.Sprintf("<tr>%s%s%s%s</tr>",
			tdIndex, tdItemName, tdRevenue, tdQuantity)

		html = fmt.Sprintf("%s%s", html, tableTopItems)
	}

	return appendSubstitude(substitudes, "list_of_top_items", html)
}

func makeDailySalesPaymentDatasSubstitute(substitudes []mpModel.Substitude, paymentIds []int64) []mpModel.Substitude {
	paymentDatas := mpPg.GetPaymentDatasFromPaymentIds(paymentIds)
	grossSalesDatas := mpPg.GetGrossSalesFromPaymentIds(paymentIds)

	var (
		discounts      int64 = 0
		gratuity       int64 = 0
		tax            int64 = 0
		grossSales     int64 = 0
		refunds        int64 = 0
		netSales       int64 = 0
		totalCollected int64 = 0
	)

	for _, paymentData := range paymentDatas {
		discounts += paymentData.TotalDiscountAmount
		gratuity += paymentData.TotalGratuityAmount
		tax += paymentData.TotalTaxAmount

		tempGrossSale := mpHelper.MakeGrossSales(grossSalesDatas, paymentData)
		grossSales += tempGrossSale

		tempRefund := mpHelper.MakeRefundsPayment(paymentData)
		refunds += tempRefund

		tempNetSales := tempGrossSale - tempRefund - paymentData.TotalDiscountAmount
		netSales += tempNetSales

		if paymentData.TotalCollectedAmount == 0 {
			totalCollected += tempNetSales + paymentData.TotalGratuityAmount + paymentData.TotalTaxAmount
		} else {
			totalCollected += paymentData.TotalCollectedAmount
		}
	}

	substitudes = appendSubstitude(substitudes, "discounts", mpUtil.Int64ToRupiah(discounts))
	substitudes = appendSubstitude(substitudes, "gratuity", mpUtil.Int64ToRupiah(gratuity))
	substitudes = appendSubstitude(substitudes, "tax", mpUtil.Int64ToRupiah(tax))
	substitudes = appendSubstitude(substitudes, "gross_sales", mpUtil.Int64ToRupiah(grossSales))
	substitudes = appendSubstitude(substitudes, "refunds", mpUtil.Int64ToRupiah(refunds))
	substitudes = appendSubstitude(substitudes, "net_sales", mpUtil.Int64ToRupiah(netSales))
	return appendSubstitude(substitudes, "total_collected", mpUtil.Int64ToRupiah(totalCollected))
}

func makeYesterdayPaymentIds(businessId int64) []int64 {
	now := mpUtil.GetJakartaTimeNow()
	today := now.Format("2006-01-02")
	yesterday := now.AddDate(0, 0, -1).Format("2006-01-02")

	return mpPg.GetPaymentIdsFromBusinessId(businessId, yesterday, today)
}

func makeDailySalesSubstitute(businessData mpModel.BusinessDataRecipients) []mpModel.Substitude {
	var substitudes []mpModel.Substitude

	paymentIds := makeYesterdayPaymentIds(businessData.BusinessId)

	substitudes = mpHelper.MakeTimeSubstitude(substitudes)
	substitudes = mpHelper.MakeBusinessDataSubstitude(substitudes, businessData)
	substitudes = makeDailySalesPaymentDatasSubstitute(substitudes, paymentIds)
	return makeTopItemsSubstitude(substitudes, paymentIds)
}

func setEmailRedisStatus(statusCode int, businessId int64) {
	if mpHelper.IsHttpSucceed(statusCode) {
		mpRedis.SetDailyBusinessIdEmailSucceed(businessId)
	} else {
		mpRedis.SetDailyBusinessIdEmailFailed(businessId)
	}
}

func sendDailySales(businessId int64) {
	businessData := mpRedis.GetBusinessEmailsByBusinessIdProfile(businessId)
	substitudes := makeDailySalesSubstitute(businessData)

	statusCode := mpSendgrid.SendDailySales(substitudes, businessData.Emails)
	setEmailRedisStatus(statusCode, businessId)
}

// Special request from our VP, please do not use this function except for operational
func SendDailySalesToSpecificEmails(businessId int64, emails []string) {
	businessData := mpRedis.GetBusinessEmailsByBusinessIdProfile(businessId)
	substitudes := makeDailySalesSubstitute(businessData)

	mpSendgrid.SendDailySales(substitudes, emails)
}

func main() {
	businessIds := mpRedis.GetDailySalesBusinessIds()

	mpRedis.ClearDailyBusinessIdsEmailSucceed()
	mpRedis.ClearDailyBusinessIdsEmailFailed()

	for i, v := range businessIds {
		if (i % mpModel.DailySalesThread) == 0 {
			sendDailySales(v)
		} else {
			go sendDailySales(v)
		}
	}

	for {
		countSucceed := mpRedis.CountDailyBusinessIdsEmailSucceed()
		countFailed := mpRedis.CountDailyBusinessIdsEmailFailed()
		total := mpRedis.CountDailySalesBusinessIds()

		if mpHelper.IsTotalEql(countSucceed, countFailed, total) {
			mpUtil.SlackDailySalesSummary(total,countSucceed, countFailed)
			break
		}

		time.Sleep(mpModel.CronSleepTime)
	}
}
