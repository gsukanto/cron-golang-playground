package mp_pg

import (
	"fmt"
	"os"
	"time"

	mpModel "../mp-model"
	mpUtil "../mp-util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

var (
	DB_RDS_HOST     = os.Getenv("DB_RDS_HOST")
	DB_RDS_PORT     = os.Getenv("DB_RDS_PORT")
	DB_RDS_USER     = os.Getenv("DB_RDS_USER")
	DB_RDS_PASSWORD = os.Getenv("DB_RDS_PASSWORD")
	DB_RDS_NAME     = os.Getenv("DB_RDS_NAME")
	DB_RDS_SSLMODE  = os.Getenv("DB_RDS_SSLMODE")

	db *gorm.DB
)

func init() {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		DB_RDS_HOST, DB_RDS_PORT, DB_RDS_USER, DB_RDS_PASSWORD, DB_RDS_NAME, DB_RDS_SSLMODE)

	var err error

	db, err = gorm.Open("postgres", dbinfo)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetBusinessIdsWhereDailyOrLowIsTrue() []int64 {
	var businessIds []int64
	db.Table("settings").Where("daily_sales_summary IS TRUE OR inventory_alerts IS TRUE").Pluck("business_id", &businessIds)
	return businessIds
}

func GetDailyBusinessIdRecipients() []int64 {
	var businessIds []int64
	db.Table("settings").Where("daily_sales_summary IS TRUE").Pluck("business_id", &businessIds)
	return businessIds
}

func GetLowInventoryBusinessIds() []int64 {
	var businessIds []int64

	query := fmt.Sprintf("SELECT %s FROM %s INNER JOIN %s INNER JOIN %s INNER JOIN %s WHERE %s AND %s AND %s AND %s",
		"DISTINCT o.business_id",
		"item_variants iv",
		"items i on i.id = iv.item_id",
		"outlets o on o.id = i.outlet_id",
		"settings s on s.business_id = o.business_id",
		"iv.in_stock <= iv.stock_alert",
		"iv.track_stock IS TRUE",
		"iv.alert IS TRUE",
		"s.inventory_alerts IS TRUE")
	db.Raw(query).Pluck("business_id", &businessIds)

	return businessIds
}

func GetBusinessProfileByBusinessIds(ids []int64) []mpModel.BusinessProfile {
	var businessProfiles []mpModel.BusinessProfile

	db.Raw("SELECT id as business_id, name, phone, email FROM businesses WHERE id in (?)", ids).Scan(&businessProfiles)

	return businessProfiles
}

func GetEmailsByBusinessId(id int64) []string {
	var emails []string

	query := fmt.Sprintf("SELECT %s FROM %s INNER JOIN %s ON %s WHERE %s = %v AND %s",
		"email",
		"setting_alert_recipients",
		"settings",
		"settings.id = setting_alert_recipients.setting_id",
		"settings.business_id",
		id,
		"setting_alert_recipients.is_deleted IS FALSE")

	db.Raw(query).Pluck("email", &emails)

	return emails
}

func GetLowInventoryAverageDailySales(itemVariantId int64) float64 {
	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)

	var sliceSumQuantity []int64
	db.Raw("SELECT SUM(quantity) AS TOTAL FROM checkouts WHERE item_variant_id = ? AND created_at >= ?",
		itemVariantId, lastMonth).Pluck("TOTAL", &sliceSumQuantity)

	sumQuantity := sliceSumQuantity[0]
	floatSumQuantity := float64(sumQuantity)
	result := floatSumQuantity / 30

	return result
}

func GetLimitedListLowInventory(businessId int64) []mpModel.ItemVariantData {
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s, %s FROM %s INNER JOIN %s INNER JOIN %s WHERE %s = %v AND %s ORDER BY %s LIMIT %v",
		"o.business_id AS business_id",
		"iv.id AS item_variant_id",
		"i.name AS item_name",
		"iv.name AS item_variant_name",
		"o.name AS outlet_name",
		"iv.in_stock",
		"item_variants iv",
		"items i on i.id = iv.item_id AND i.is_deleted IS FALSE",
		"outlets o on o.id = i.outlet_id AND o.is_deleted IS FALSE",
		"o.business_id",
		businessId,
		"iv.in_stock <= iv.stock_alert AND iv.track_stock IS TRUE AND iv.alert IS TRUE AND iv.is_deleted IS FALSE",
		"o.name, i.name",
		mpModel.LimitListLowInventory)

	var listLowInventoryData []mpModel.ItemVariantData
	db.Raw(query).Scan(&listLowInventoryData)

	return listLowInventoryData
}

func CountListLowInventory(businessId int64) int64 {
	query := fmt.Sprintf("SELECT COUNT(%s) as count FROM %s INNER JOIN %s INNER JOIN %s WHERE %s = %v AND %s",
		"*",
		"item_variants iv",
		"items i on i.id = iv.item_id",
		"outlets o on o.id = i.outlet_id",
		"o.business_id",
		businessId,
		"iv.in_stock <= iv.stock_alert AND iv.track_stock IS TRUE AND iv.alert IS TRUE AND iv.is_deleted IS FALSE")

	var count []int64
	db.Raw(query).Pluck("count", &count)

	return count[0]
}

func GetOutletIdsfromBusinessId(businessId int64) []int64 {
	var outletIds []int64
	db.Table("outlets").Where("business_id = ? and is_deleted IS FALSE", businessId).Pluck("id", &outletIds)
	return outletIds
}

func GetPaymentIdsFromBusinessId(businessId int64, startDate string, endDate string) []int64 {
	outletIds := GetOutletIdsfromBusinessId(businessId)

	query := fmt.Sprintf("SELECT id FROM %s WHERE %s IN (%s) AND %s BETWEEN '%s' AND '%s' AND %s",
		"payments",
		"outlet_id",
		mpUtil.Ints64ToString(outletIds),
		"created_at",
		startDate,
		endDate,
		"is_deleted is false")

	var paymentIds []int64
	db.Raw(query).Pluck("id", &paymentIds)

	return paymentIds
}

func GetPaymentDatasFromPaymentIds(paymentIds []int64) []mpModel.PaymentData {
	query := fmt.Sprintf("SELECT id, %s, %s, %s, %s, %s, %s, %s, %s, %s FROM %s WHERE id IN (%s)",
		"total_collected_amount",
		"total_discount_amount",
		"total_gratuity_amount",
		"total_tax_amount",
		"is_refunded",
		"refund_type",
		"refund_amount",
		"include_gratuity_tax",
		"parent_payment_id",
		"payments",
		mpUtil.Ints64ToString(paymentIds))

	var paymentDatas []mpModel.PaymentData
	db.Raw(query).Scan(&paymentDatas)

	return paymentDatas
}

func GetGrossSalesFromPaymentIds(paymentIds []int64) []mpModel.GrossSales {
	query := fmt.Sprintf("SELECT %s, %s, %s, %s, %s FROM %s WHERE %s IN (%s)",
		"payment_id",
		"item_price_library",
		"item_price",
		"quantity",
		"gross_sales",
		"checkouts",
		"payment_id",
		mpUtil.Ints64ToString(paymentIds))

	var grossSales []mpModel.GrossSales
	db.Raw(query).Scan(&grossSales)

	return grossSales
}

func GetTopItems(paymentIds []int64) []mpModel.TopItem {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE payment_id in (SELECT id FROM %s WHERE id IN (%s) AND %s) GROUP BY %s ORDER BY %s",
		"item_name, SUM(item_price_library * quantity) as revenue, SUM(quantity) as quantity",
		"checkouts",
		"payments",
		mpUtil.Ints64ToString(paymentIds),
		"is_refunded is false",
		"item_name",
		"quantity DESC")

	var topItems []mpModel.TopItem
	db.Raw(query).Scan(&topItems)

	return topItems
}
