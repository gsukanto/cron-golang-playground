package mp_model

/* ############################## Email Model ############################## */
// email is business_email and emails is the list of email recipients
type BusinessDataRecipients struct {
	BusinessId int64
	Name       string
	Phone      string
	Email      string
	Emails     []string
}

type LowInventoryData struct {
	BusinessId      int64
	ItemName        string
	ItemVariantName string
	OutletName      string
	InStock         int
	AvgDailySales   float64
}

// for sendgrid substitution/section maker key value
type Substitude struct {
	Key   string
	Value string
}

type EmailHeader struct {
	Subject            string
	ListRecipientEmail []string
	ListSubstitutions  []Substitude
}

/* ############################## End Email Model ############################## */

/* ############################## Postgres Query Model ############################## */
type BusinessProfile struct {
	BusinessId int64
	Name       string
	Phone      string
	Email      string
}

type ItemVariantData struct {
	BusinessId      int64
	ItemVariantId   int64
	ItemName        string
	ItemVariantName string
	OutletName      string
	InStock         int
	AvgDailySales   float64
}

type PaymentData struct {
	Id                   int
	TotalCollectedAmount int64
	TotalDiscountAmount  int64
	TotalGratuityAmount  int64
	TotalTaxAmount       int64
	IsRefunded           bool
	RefundType           string
	RefundAmount         int64
	IncludeGratuityTax   bool
	ParentPaymentId      int
}

type GrossSales struct {
	PaymentId        int
	ItemPriceLibrary int64
	ItemPrice        int64
	Quantity         int
	GrossSales       int64
}

type TopItem struct {
	ItemName string
	Revenue  int64
	Quantity int64
}

/* ############################## End Query Model ############################## */
