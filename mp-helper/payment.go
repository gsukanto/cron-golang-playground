package mp_helper

import (
	"strings"

	mpModel "../mp-model"
	mpPg "../mp-pg"
)

// TODO: use better way for get gross sales, for example store it in database
func MakeGrossSales(grossSalesDatas []mpModel.GrossSales, paymentData mpModel.PaymentData) int64 {
	var grossSales int64 = 0
	grossSales = getGrossSalesFromArray(grossSalesDatas, paymentData, grossSales)
	return grossSales
}

// TODO: use better way to get refunds, for example store it in database
func MakeRefundsPayment(paymentData mpModel.PaymentData) int64 {
	refunds := paymentData.RefundAmount

	if strings.Compare(paymentData.RefundType, mpModel.FullRefund) == 0 {
		if paymentData.IncludeGratuityTax {
			refunds = getParentPaymentGrossSales(paymentData.ParentPaymentId)
		} else {
			refunds += paymentData.TotalTaxAmount
			refunds += paymentData.TotalGratuityAmount
			refunds -= paymentData.TotalDiscountAmount
		}
	}

	return refunds
}

func getParentPaymentGrossSales(parentPaymentId int) int64 {
	var grossSales int64 = 0
	var parentPaymentIds []int64 = []int64{int64(parentPaymentId)}

	grossSalesDatas := mpPg.GetGrossSalesFromPaymentIds(parentPaymentIds)
	paymentDatas := mpPg.GetPaymentDatasFromPaymentIds(parentPaymentIds)
	paymentData := paymentDatas[0]
	grossSales = getGrossSalesFromArray(grossSalesDatas, paymentData, grossSales)

	return grossSales
}

func getGrossSalesFromArray(grossSalesDatas []mpModel.GrossSales, paymentData mpModel.PaymentData, grossSales int64) int64 {
	for _, grossSalesData := range grossSalesDatas {
		if grossSalesData.PaymentId == paymentData.Id {
			if paymentData.IncludeGratuityTax {
				if grossSalesData.GrossSales > 0 {
					grossSales += grossSalesData.GrossSales
				} else {
					grossSales += (grossSalesData.ItemPriceLibrary * int64(grossSalesData.Quantity))
				}
			} else {
				grossSales += (grossSalesData.ItemPrice * int64(grossSalesData.Quantity))
			}
		}
	}
	return grossSales
}
