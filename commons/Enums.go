package commons

type ProtocolType string

const (
	Protocol_PROTOBUF ProtocolType = "1"
	Protocol_JSON     ProtocolType = "2"
)

func (p ProtocolType) ValueOf(value string) ProtocolType {
	switch value {
	case "1":
		return Protocol_PROTOBUF
	case "2":
		return Protocol_JSON
	default:
		return "0"
	}
}

func (p ProtocolType) String() string {
	switch p {
	case Protocol_PROTOBUF:
		return "1"
	case Protocol_JSON:
		return "2"
	default:
		return "0"
	}
}
