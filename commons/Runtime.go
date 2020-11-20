package commons

import (
	"linking/linking-go-agile/redis"
)

type Runtimes struct {
}

func (*Runtimes) GetProtocolType() string {
	var dataType = redis.RGet("default", "runtime:protocol_type")
	if IsEmpty(dataType) {
		return Protocol_JSON.String()
	} else {
		return dataType
	}
}

func (*Runtimes) SetProtocolType(protocolType string) {
	redis.RSet("default", "runtime:protocol_type", protocolType)
}

var SelfRuntime = &Runtimes{}
