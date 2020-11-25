package common

import (
	"encoding/json"
	"linking/linking-go-agile/model"
)

type LResult struct {
	OK   bool
	Code string
	Msg  string
	Data interface{} `json:"Data,omitempty"`
}

var ResultOk = &model.LResult{OK: true, Code: "0", Msg: ""}

func TestFail(lResult *model.LResult) bool {
	return !lResult.OK
}

func OfResultData(data interface{}) *model.LResult {
	if data == nil {
		return &model.LResult{OK: true, Code: "0", Msg: ""}
	} else {
		jsonB, _ := json.Marshal(data)
		return &model.LResult{OK: true, Code: "0", Msg: "", Data: string(jsonB)}
	}
}

func OfResultFail(code string, msg string) *model.LResult {
	return &model.LResult{OK: false, Code: code, Msg: msg}
}
