package common

import (
	"encoding/json"
	"github.com/ganeryao/linking-go-agile/protos"
)

type LResult struct {
	Api  string      `json:"api,omitempty"`
	Ok   bool        `json:"ok,omitempty"`
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func TestFail(lResult *protos.LResult) bool {
	return !lResult.Ok
}

func OfResultData(api string, data interface{}) *protos.LResult {
	if data == nil {
		return &protos.LResult{Api: api, Ok: true, Code: "0", Msg: ""}
	} else {
		jsonB, _ := json.Marshal(data)
		return &protos.LResult{Api: api, Ok: true, Code: "0", Msg: "", Data: string(jsonB)}
	}
}

func OfResultFail(api string, code string, msg string) *protos.LResult {
	return &protos.LResult{Api: api, Ok: false, Code: code, Msg: msg}
}
