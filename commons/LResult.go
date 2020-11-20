package commons

import (
	"encoding/base64"
	"encoding/json"
	"github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"linking/linking-go-agile/models"
)

type LResult struct {
	OK   bool
	Code string
	Msg  string
	Data interface{} `json:"Data,omitempty"`
}

func ConvertResult(result *models.LResult) string {
	var dataType = SelfRuntime.GetProtocolType()
	switch dataType {
	case Protocol_PROTOBUF.String():
		data, _ := proto.Marshal(result)
		return base64.StdEncoding.EncodeToString(data)
	case Protocol_JSON.String():
		jsonStr, _ := json.Marshal(convertJsonResult(result))
		return string(jsonStr)
	default:
		return ""
	}
}

func convertJsonResult(result *models.LResult) LResult {
	var sResult LResult
	sResult.OK = result.OK
	sResult.Code = result.Code
	sResult.Msg = result.Msg
	var data = result.GetData()
	var obj interface{}
	if !IsEmpty(data) {
		var b = []byte(data)
		if json.Valid(b) {
			if err := json.Unmarshal(b, &obj); err == nil {
				sResult.Data = obj
			} else {
				_ = log4go.Error("ConvertResult string 2 json error: ", err)
			}
		}
	}
	return sResult
}

func TestFail(lResult *models.LResult) bool {
	return !lResult.OK
}

func OfResult() models.LResult {
	return models.LResult{OK: true, Code: "0", Msg: ""}
}

func OfResultData(data interface{}) models.LResult {
	if data == nil {
		return models.LResult{OK: true, Code: "0", Msg: ""}
	} else {
		jsonB, _ := json.Marshal(data)
		return models.LResult{OK: true, Code: "0", Msg: "", Data: string(jsonB)}
	}
}

func OfResultFail(code string, msg string) models.LResult {
	return models.LResult{OK: false, Code: code, Msg: msg}
}
