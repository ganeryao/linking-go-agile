package common

import (
	"encoding/base64"
	"encoding/json"
	"github.com/alecthomas/log4go"
	"github.com/golang/protobuf/proto"
	"github.com/kataras/iris"
	"linking/linking-go-agile/model"
)

func ConvertJson(data interface{}) string {
	jsonStr, _ := json.Marshal(data)
	return string(jsonStr)
}

func ParseJson(str string, data interface{}) interface{} {
	_ = json.Unmarshal([]byte(str), data)
	return data
}

func ConvertDTO(ctx iris.Context, m proto.Message) {
	var param = ctx.PostValue("param")
	if len(param) == 0 {
		_ = ctx.ReadBody(&param)
	}
	var dataType = SelfRuntime.GetProtocolType()
	switch dataType {
	case Protocol_PROTOBUF.String():
		var data, _ = base64.StdEncoding.DecodeString(param)
		_ = proto.Unmarshal(data, m)
	case Protocol_JSON.String():
		_ = json.Unmarshal([]byte(param), m)
	}
}

func ConvertResult(result *model.LResult) string {
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

func convertJsonResult(result *model.LResult) LResult {
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
