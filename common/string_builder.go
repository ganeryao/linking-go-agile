package common

import (
	"strconv"
	"strings"
)

type StringBuilder struct {
	buffer strings.Builder
}

func IsEmpty(str string) bool {
	return len(str) == 0
}

func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StrToInt64(str string) int64 {
	i, _ := strconv.ParseInt(str, 0, 64)
	return i
}

func NewStringBuilder() *StringBuilder {
	var builder StringBuilder
	return &builder
}

func NewStringBuilderString(str string) *StringBuilder {
	var builder StringBuilder
	builder.buffer.WriteString(str)
	return &builder
}

func (builder *StringBuilder) Append(s string) *StringBuilder {
	builder.buffer.WriteString(s)
	return builder
}

func (builder *StringBuilder) AppendStrings(ss ...string) *StringBuilder {
	for i := range ss {
		builder.buffer.WriteString(ss[i])
	}
	return builder
}

func (builder *StringBuilder) AppendInt(i int) *StringBuilder {
	builder.buffer.WriteString(strconv.Itoa(i))
	return builder
}

func (builder *StringBuilder) AppendInt64(i int64) *StringBuilder {
	builder.buffer.WriteString(strconv.FormatInt(i, 10))
	return builder
}

func (builder *StringBuilder) AppendFloat64(f float64) *StringBuilder {
	builder.buffer.WriteString(strconv.FormatFloat(f, 'f', -1, 32))
	return builder
}

func (builder *StringBuilder) Clear() *StringBuilder {
	builder.Clear()
	return builder
}

func (builder *StringBuilder) ToString() string {
	return builder.buffer.String()
}

func (builder *StringBuilder) IsEmpty() bool {
	return builder.buffer.Len() == 0
}
