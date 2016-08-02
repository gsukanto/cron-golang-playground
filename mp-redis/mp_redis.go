package mp_redis

import (
	"fmt"
	"os"

	mpModel "../mp-model"
	mpUtil "../mp-util"

	"gopkg.in/redis.v4"
)

/* ############################## Init Redis ############################## */
var (
	redisHost = os.Getenv("REDIS_HOST")
	redisPort = os.Getenv("REDIS_PORT")

	redisClient *redis.Client
)

func init() {
	dbAddress := fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisClient = redis.NewClient(&redis.Options{Addr: dbAddress})
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func DelKey(k string) {
	_, e := redisClient.Del(k).Result()
	checkErr(e)
}

func MakeBusinessEmailKey(id int64) string {
	return fmt.Sprintf("%s%v", mpModel.BusinessEmailProfile, id)
}

/* ############################## End Init Redis ############################## */

/* ############################## List Email Recipient Redis ############################## */
func SetAllBusinessIdsRecipient(ids []int64) {
	for _, id := range ids {
		redisClient.RPush(mpModel.AllBusinessRecipient, id)
	}
}

func GetAllBusinessIdsRecipient() []int64 {
	a, err := redisClient.LRange(mpModel.AllBusinessRecipient, 0, -1).Result()
	checkErr(err)

	result := mpUtil.StringsToInts64(a)

	return result
}

func DelAllBusinessIdsRecipient() {
	DelKey(mpModel.AllBusinessRecipient)
}

func SetBusinessEmailsByBusinessId(id int64, emails []string) {
	key := MakeBusinessEmailKey(id)

	for _, email := range emails {
		redisClient.RPush(key, email)
	}
}

func GetBusinessEmailsByBusinessId(id int64) []string {
	key := MakeBusinessEmailKey(id)

	emails, err := redisClient.LRange(key, 3, -1).Result()
	checkErr(err)

	return emails
}

func SetBusinessProfileByBusinessId(id int64, name string, phone string, email string) {
	key := MakeBusinessEmailKey(id)

	DelKey(key)
	redisClient.RPush(key, name)
	redisClient.RPush(key, phone)
	redisClient.RPush(key, email)
}

func GetBusinessProfileByBusinessIds(id int64) []string {
	key := MakeBusinessEmailKey(id)

	profiles, err := redisClient.LRange(key, 0, 2).Result()
	checkErr(err)

	return profiles
}

func GetBusinessEmailsByBusinessIdProfile(id int64) mpModel.BusinessDataRecipients {
	key := MakeBusinessEmailKey(id)

	results, err := redisClient.LRange(key, 0, -1).Result()
	checkErr(err)

	businessEmailProfile := mpModel.BusinessDataRecipients{id, results[0], results[1], results[2], results[3:]}

	return businessEmailProfile
}

func FlushAllBusinessProfileById(id int64) {
	key := MakeBusinessEmailKey(id)

	DelKey(key)
}

/* ############################## End Email Recipient List ############################## */

/* ############################## Low Inventory ############################## */

func SetLowInventoryBusinessIds(ids []int64) {
	for _, id := range ids {
		redisClient.RPush(mpModel.LowInventoryBusinessIds, id)
	}
}

func GetLowInventoryBusinessIds() []int64 {
	a, err := redisClient.LRange(mpModel.LowInventoryBusinessIds, 0, -1).Result()
	checkErr(err)
	result := mpUtil.StringsToInts64(a)
	return result
}

func CountListLowInventoryBusinessIds() int64 {
	result, err := redisClient.LLen(mpModel.LowInventoryBusinessIds).Result()
	checkErr(err)
	return result
}

func DelLowInventoryBusinessIds() {
	DelKey(mpModel.LowInventoryBusinessIds)
}

func SetLowBusinessIdEmailSucceed(id int64) {
	redisClient.RPush(mpModel.LowBusinessIdsEmailSucceed, id)
}

func GetLowBusinessIdsEmailSucceed() []int64 {
	a, err := redisClient.LRange(mpModel.LowBusinessIdsEmailSucceed, 0, -1).Result()
	checkErr(err)
	result := mpUtil.StringsToInts64(a)
	return result
}

func CountLowBusinessIdEmailSucceed() int64 {
	result, err := redisClient.LLen(mpModel.LowBusinessIdsEmailSucceed).Result()
	checkErr(err)
	return result
}

func DelLowBusinessIdEmailSucceed() {
	DelKey(mpModel.LowBusinessIdsEmailSucceed)
}

func SetLowBusinessIdEmailFailed(id int64) {
	redisClient.RPush(mpModel.LowBusinessIdsEmailFailed, id)
}

func GetLowBusinessIdsEmailFailed() []int64 {
	a, err := redisClient.LRange(mpModel.LowBusinessIdsEmailFailed, 0, -1).Result()
	checkErr(err)
	result := mpUtil.StringsToInts64(a)
	return result
}

func CountLowBusinessIdEmailFailed() int64 {
	result, err := redisClient.LLen(mpModel.LowBusinessIdsEmailFailed).Result()
	checkErr(err)
	return result
}

func DelLowBusinessIdEmailFailed() {
	DelKey(mpModel.LowBusinessIdsEmailFailed)
}

func FlushAllLowBusiness() {
	DelLowInventoryBusinessIds()
	DelLowBusinessIdEmailSucceed()
	DelLowBusinessIdEmailFailed()
}

/* ############################## End Low Inventory ############################## */

/* ############################## Daily Sales ############################## */

func SetDailySalesBusinessIds(ids []int64) {
	for _, id := range ids {
		redisClient.RPush(mpModel.DailySalesBusinessIds, id)
	}
}

func GetDailySalesBusinessIds() []int64 {
	a, err := redisClient.LRange(mpModel.DailySalesBusinessIds, 0, -1).Result()
	checkErr(err)

	result := mpUtil.StringsToInts64(a)

	return result
}

func CountDailySalesBusinessIds() int64 {
	result, err := redisClient.LLen(mpModel.DailySalesBusinessIds).Result()
	checkErr(err)
	return result
}

func DelDailySalesBusinessId() {
	DelKey(mpModel.DailySalesBusinessIds)
}

func SetDailyBusinessIdEmailSucceed(id int64) {
	redisClient.RPush(mpModel.DailyBusinessIdsSuccess, id)
}

func GetDailyBusinessIdsEmailSucceed() []int64 {
	a, err := redisClient.LRange(mpModel.DailyBusinessIdsSuccess, 0, -1).Result()
	checkErr(err)

	result := mpUtil.StringsToInts64(a)

	return result
}

func CountDailyBusinessIdsEmailSucceed() int64 {
	result, err := redisClient.LLen(mpModel.DailyBusinessIdsSuccess).Result()
	checkErr(err)
	return result
}

func DelDailyBusinessIdsEmailSucceed() {
	DelKey(mpModel.DailyBusinessIdsSuccess)
}

func SetDailyBusinessIdEmailFailed(id int64) {
	redisClient.RPush(mpModel.DailyBusinessIdsFailed, id)
}

func GetDailyBusinessIdsEmailFailed() []int64 {
	a, err := redisClient.LRange(mpModel.DailyBusinessIdsFailed, 0, -1).Result()
	checkErr(err)

	result := mpUtil.StringsToInts64(a)

	return result
}

func CountDailyBusinessIdsEmailFailed() int64 {
	result, err := redisClient.LLen(mpModel.DailyBusinessIdsFailed).Result()
	checkErr(err)
	return result
}

func DelDailyBusinessIdsEmailFailed() {
	DelKey(mpModel.DailyBusinessIdsFailed)
}

func FlushAllDailyBusiness() {
	DelDailySalesBusinessId()
	DelDailyBusinessIdsEmailSucceed()
	DelDailyBusinessIdsEmailFailed()
}

/* ############################## End Daily Sales ############################## */
