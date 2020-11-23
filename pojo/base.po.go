package pojo

import (
	"strconv"
	"time"
)

type PO interface {
	Initial()
}

type AbstractPO struct {
	Id         int64  `db:"id,key"`
	CreateTime uint64 `db:"create_time"`
	UpdateTime uint64 `db:"update_time"`
}

func (po *AbstractPO) Initial() {
	timeStr := time.Now().Format("20060102150405")
	po.CreateTime, _ = strconv.ParseUint(timeStr, 10, 64)
	po.UpdateTime = po.CreateTime
}