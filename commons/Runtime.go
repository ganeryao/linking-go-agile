package commons

import (
	"linking/linking-go-agile/redis"
)

type Runtimes struct {
}

func (*Runtimes) GetProtocolType() string {
	return redis.RGet("default", "runtime:protocol_type")
}

func (*Runtimes) SetProtocolType(protocolType string) {
	redis.RSet("default", "runtime:protocol_type", protocolType)
}

var SelfRuntime = &Runtimes{}
