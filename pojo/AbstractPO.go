package pojo

import (
	"strconv"
	"time"
)

type AbstractPO struct {
	Id         int64
	CreateTime uint64
	UpdateTime uint64
}

func (po *AbstractPO) Initial() {
	timeStr := time.Now().Format("20060102150405")
	po.CreateTime, _ = strconv.ParseUint(timeStr, 10, 64)
}
