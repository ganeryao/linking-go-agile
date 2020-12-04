/**
 * @Title  公用包
 * @Description 公用的属性和方法
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:07
 */
package common

type ProtocolType string

const (
	ProtocolProtobuf ProtocolType = "1"
	ProtocolJson     ProtocolType = "2"
)

func (p ProtocolType) ValueOf(value string) ProtocolType {
	switch value {
	case "1":
		return ProtocolProtobuf
	case "2":
		return ProtocolJson
	default:
		return "0"
	}
}

func (p ProtocolType) String() string {
	switch p {
	case ProtocolProtobuf:
		return "1"
	case ProtocolJson:
		return "2"
	default:
		return "0"
	}
}
