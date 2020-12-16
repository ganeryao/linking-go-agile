package strs

import (
	"encoding/binary"
	"math"
	"strconv"
)

func IsEmpty(str string) bool {
	return len(str) == 0
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func StrToInit(value string) int {
	v, _ := strconv.Atoi(value)
	return v
}

func Int64ToStr(value int64) string {
	return strconv.FormatInt(value, 10)
}

func StrToInt64(value string) int64 {
	v, _ := strconv.ParseInt(value, 10, 64)
	return v
}

func FloatToStr(value float64) string {
	return strconv.FormatFloat(value, 'E', -1, 64)
}

func StrToFloat(value string) float64 {
	v, _ := strconv.ParseFloat(value, 64)
	return v
}

func BoolToStr(value bool) string {
	return strconv.FormatBool(value)
}

func StrToBool(value string) bool {
	v, _ := strconv.ParseBool(value)
	return v
}

func InterfaceToStr(value interface{}) string {
	return value.(string)
}

func ByteToStr(value interface{}) string {
	return string(value.([]byte))
}

func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func ByteToFloat64(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits)
}
