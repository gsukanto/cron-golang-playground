package mp_util

import (
	"github.com/leekchan/accounting"
)

func Int64ToRupiah(money int64) string {
	ac := accounting.Accounting{Symbol: "Rp. ", Precision: 2, Thousand: ",", Decimal: "."}
	return ac.FormatMoney(money)
}
