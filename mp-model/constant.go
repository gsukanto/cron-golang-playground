package mp_model

const (
	DailySalesBusinessIds   = "daily_sales_business_ids"
	DailyBusinessIdsSuccess = "daily_business_ids_success"
	DailyBusinessIdsFailed  = "daily_business_ids_failed"

	LowInventoryBusinessIds    = "low_inventory_business_ids"
	LowBusinessIdsEmailSucceed = "low_business_ids_success"
	LowBusinessIdsEmailFailed  = "low_business_ids_failed"

	BusinessEmailProfile = "business_email_profile:"
	AllBusinessRecipient = "all_business_recipient"

	LimitListLowInventory = 15
	CronSleepTime         = 2
	DailySalesThread      = 2
	LowInventoryThread    = 2
	LimitNameLength       = 30

	FullRefund = "full"
)
