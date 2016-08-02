package mp_helper

import (
	"fmt"

	mpModel "../mp-model"
	mpUtil "../mp-util"
)

func IsTotalEql(succeed int64, failed int64, total int64) bool {
	return (succeed + failed) == total
}

func IsHttpSucceed(statusCode int) bool {
	if statusCode >= 200 && statusCode < 300 {
		return true
	}
	return false
}

func MakeTimeSubstitude(substitudes []mpModel.Substitude) []mpModel.Substitude {
	substitudes = append(substitudes, mpModel.Substitude{"date", fmt.Sprintf("%s", mpUtil.GetJakartaWeekdayDateYesterday())})
	return substitudes
}

func MakeBusinessDataSubstitude(substitudes []mpModel.Substitude, businessData mpModel.BusinessDataRecipients) []mpModel.Substitude {
	substitudes = append(substitudes, mpModel.Substitude{"business_name", businessData.Name})
	substitudes = append(substitudes, mpModel.Substitude{"business_phone", businessData.Phone})
	substitudes = append(substitudes, mpModel.Substitude{"business_email", businessData.Email})
	return substitudes
}
