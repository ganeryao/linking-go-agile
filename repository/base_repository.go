/**
 * @Title  仓库包
 * @Description 存放仓库的类
 * @Author YaoWeiXin
 * @Update 2020/11/20 17:24
 */
package repository

import (
	"github.com/ganeryao/linking-go-agile/common"
	"github.com/ganeryao/linking-go-agile/errors"
	"github.com/ganeryao/linking-go-agile/mysql"
	"github.com/ganeryao/linking-go-agile/redis"
)

type BaseRepository struct {
}

func GetAlways(db string, table string, dest interface{}, sql string, id int64) (interface{}, *errors.Error) {
	var key = "cache:" + table + ":" + common.Int64ToStr(id)
	var str = redis.RGet(db, key)
	if common.IsEmpty(str) {
		mysql.MFindOne(db, dest, sql, id)
		str, err := common.ConvertJson(dest)
		if err != nil {
			return nil, err
		}
		redis.RSet(db, key, str)
	} else {
		_, err := common.ParseJson(str, dest)
		if err != nil {
			return nil, err
		}
	}
	return dest, nil
}
