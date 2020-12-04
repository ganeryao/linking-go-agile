package common

import (
	"encoding/base64"
	"encoding/json"
	"github.com/ganeryao/linking-go-agile/errors"
	"github.com/ganeryao/linking-go-agile/protos"
	"github.com/ganeryao/linking-go-agile/serialize"
	lkJson "github.com/ganeryao/linking-go-agile/serialize/json"
	"github.com/golang/protobuf/proto"
)

type ConvertUtils struct {
	serializer serialize.Serializer
}

var (
	convert = &ConvertUtils{
		serializer: lkJson.NewSerializer(),
	}
)

// SetSerializer customize application serializer, which automatically Marshal
// and UnMarshal handler payload
func SetSerializer(ser serialize.Serializer) {
	convert.serializer = ser
}

// GetSerializer gets the app serializer
func GetSerializer() serialize.Serializer {
	return convert.serializer
}

func ConvertJson(data interface{}) (string, *errors.Error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return "", errors.NewError(err, errors.ErrInternalCode)
	}
	return string(jsonStr), nil
}

func ParseJson(str string, data interface{}) (interface{}, *errors.Error) {
	err := json.Unmarshal([]byte(str), data)
	if err != nil {
		return "", errors.NewError(err, errors.ErrInternalCode)
	}
	return data, nil
}

func ConvertRequest(param string, m proto.Message) *errors.Error {
	serializerName := convert.serializer.GetName()
	var b []byte
	var err error
	switch serializerName {
	case "protos":
		b, err = base64.StdEncoding.DecodeString(param)
	case "json":
		b = []byte(param)
	default:
		return errors.NewError(errors.ErrWrongSerializer, errors.ErrInternalCode)
	}
	if err != nil {
		return errors.NewError(err, errors.ErrBadRequestCode)
	}
	err = convert.serializer.Unmarshal(b, m)
	if err != nil {
		return errors.NewError(err, errors.ErrBadRequestCode)
	}
	return nil
}

func ConvertResult(result *protos.LResult) (string, *errors.Error) {
	serializerName := convert.serializer.GetName()
	var b []byte
	var err error
	var data string
	switch serializerName {
	case "protos":
		b, err = convert.serializer.Marshal(result)
		data = base64.StdEncoding.EncodeToString(b)
	case "json":
		b, err = convert.serializer.Marshal(convertJsonResult(result))
		data = string(b)
	default:
		panic(`ConvertRequest type error: type(` + serializerName + `): `)
	}
	if err != nil {
		return "", errors.NewError(err, errors.ErrBadRequestCode)
	}
	return data, nil
}

func convertJsonResult(result *protos.LResult) LResult {
	var sResult LResult
	sResult.Ok = result.Ok
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
