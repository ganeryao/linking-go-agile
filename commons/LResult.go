package commons

import (
	"encoding/json"
)

type LResult struct {
	OK   bool
	Code string
	Msg  string
	Data interface{} `json:"Data,omitempty"`
}

func (result LResult) TestFail() bool {
	return !result.OK
}

func (result LResult) ToString() string {
	b, _ := json.Marshal(result)
	return string(b)
}

func OfResult() LResult {
	return LResult{OK: true, Code: "0", Msg: "", Data: nil}
}

func OfResultData(data interface{}) LResult {
	return LResult{OK: true, Code: "0", Msg: "", Data: data}
}

func OfResultFail(code string, msg string) LResult {
	return LResult{OK: false, Code: code, Msg: msg}
}
