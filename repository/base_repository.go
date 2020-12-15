/**
 * @Title  仓库包
 * @Description 存放仓库的类
 * @Author YaoWeiXin
 * @Update 2020/11/20 17:24
 */
package repository

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/mysql"
	"github.com/ganeryao/linking-go-agile/redis"
	"github.com/ganeryao/linking-go-agile/strs"
)

type BaseRepository struct {
}

func GetAlways(db string, table string, dest interface{}, sql string, id int64) interface{} {
	var key = "cache:" + table + ":" + strs.Int64ToStr(id)
	var str = redis.RGet(db, key)
	if strs.IsEmpty(str) {
		mysql.MFindOne(db, dest, sql, id)
		redis.RSet(db, key, common.ConvertJson(dest))
	} else {
		common.ParseJson(str, dest)
	}
	return dest
}
