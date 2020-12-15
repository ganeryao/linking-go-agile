package common

import (
	"strconv"
)

func Int2Str(value int) string {
	return strconv.Itoa(value)
}

func Str2Init(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Int642Str(value int64) string {
	return strconv.FormatInt(value, 10)
}

func Str2Int64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

func Float2Str(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}

func Str2Float(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func Bool2Str(value bool) string {
	return strconv.FormatBool(value)
}

func Str2Bool(value string) bool {
	v, _ := strconv.ParseBool(value)
	return v
}

func Interface2Str(value interface{}) string {
	return value.(string)
}

func Byte2Str(value interface{}) string {
	return string(value.([]byte))
}
