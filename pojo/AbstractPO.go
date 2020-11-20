/**
 * @Title  pojo包
 * @Description 存放结构体对象po vo等
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:09
 */
package pojo

import (
	"strconv"
	"time"
)

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
