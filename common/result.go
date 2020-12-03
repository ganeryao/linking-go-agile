package common

import (
	"encoding/json"
	"github.com/ganeryao/linking-go-agile/protos"
)

type LResult struct {
	Ok   bool        `json:"ok,omitempty"`
	Code string      `json:"code,omitempty"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

var ResultOk = &protos.LResult{Ok: true, Code: "0", Msg: ""}

func TestFail(lResult *protos.LResult) bool {
	return !lResult.Ok
}

func OfResultData(data interface{}) *protos.LResult {
	if data == nil {
		return &protos.LResult{Ok: true, Code: "0", Msg: ""}
	} else {
		jsonB, _ := json.Marshal(data)
		return &protos.LResult{Ok: true, Code: "0", Msg: "", Data: string(jsonB)}
	}
}

func OfResultFail(code string, msg string) *protos.LResult {
	return &protos.LResult{Ok: false, Code: code, Msg: msg}
}
