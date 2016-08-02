package mp_util

import (
	"strconv"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func StringsToInts64(s []string) []int64 {
	var result []int64

	for _, v := range s {
		tempInt, e := strconv.ParseInt(v, 10, 64)
		checkErr(e)
		result = append(result, tempInt)
	}

	return result
}

func Ints64ToString(intArray []int64) string {
	var tempString []string

	for _, v := range intArray {
		s := strconv.FormatInt(v, 10)
		tempString = append(tempString, s)
	}

	result := strings.Join(tempString, ",")
	return result
}

func TruncateString(s string, limit int) string {
	if len(s) > limit {
		return s[:limit]
	}

	return s
}
