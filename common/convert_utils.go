package common

import (
	"encoding/base64"
	"encoding/json"
	"github.com/golang/protobuf/proto"
	"github.com/kataras/iris"
	"linking/linking-go-agile/model"
)

func ConvertJson(data interface{}) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(`ConvertJson: ` + err.Error())
	}
	return string(jsonStr)
}

func ParseJson(str string, data interface{}) interface{} {
	err := json.Unmarshal([]byte(str), data)
	if err != nil {
		panic(`ParseJson: str(` + str + `): ` + err.Error())
	}
	return data
}

func ConvertRequest(ctx iris.Context, m proto.Message) {
	var param = ctx.PostValue("param")
	if len(param) == 0 {
		_ = ctx.ReadBody(&param)
	}
	var dataType = SelfRuntime.GetProtocolType()
	switch dataType {
	case Protocol_PROTOBUF.String():
		var data, err = base64.StdEncoding.DecodeString(param)
		if err != nil {
			panic(`ConvertRequest base64.StdEncoding.DecodeString: str(` + param + `): ` + err.Error())
		}
		err = proto.Unmarshal(data, m)
		if err != nil {
			panic(`ConvertRequest proto.Unmarshal: str(` + param + `): ` + err.Error())
		}
	case Protocol_JSON.String():
		err := json.Unmarshal([]byte(param), m)
		if err != nil {
			panic(`ConvertRequest json.Unmarshal: str(` + param + `): ` + err.Error())
		}
	}
}

func ConvertResult(result *model.LResult) string {
	var protocolType = SelfRuntime.GetProtocolType()
	switch protocolType {
	case Protocol_PROTOBUF.String():
		data, err := proto.Marshal(result)
		if err != nil {
			panic(`ConvertResult proto.Marshal: ` + err.Error())
		}
		return base64.StdEncoding.EncodeToString(data)
	case Protocol_JSON.String():
		jsonStr, err := json.Marshal(convertJsonResult(result))
		if err != nil {
			panic(`ConvertResult json.Marshal: ` + err.Error())
		}
		return string(jsonStr)
	default:
		panic(`ConvertResult protocolType in default: ` + protocolType)
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
				panic(`convertJsonResult string 2 json error: ` + err.Error())
			}
		}
	}
	return sResult
}
