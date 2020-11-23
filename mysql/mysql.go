/**
 * @Title  mysql操作层
 * @Description mysql操作的封装
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:08
 */
package mysql

import (
	"github.com/alecthomas/log4go"
	"strings"
)

func Init(mysqlConfig MConfig) {
	initMysql(mysqlConfig)
}

/**
查找一个数据
*/
func MFindOne(db string, dest interface{}, sql string, args ...interface{}) interface{} {
	err := getConn(db).Get(dest, sql, args...)
	if err != nil {
		return nil
	}
	return dest
}

/**
查找数据列表
*/
func MFind(db string, dest interface{}, sql string, args ...interface{}) interface{} {
	sql = strings.TrimSpace(sql)
	if !strings.Contains(sql, " limit ") {
		sql += " limit 1000"
	}
	err := getConn(db).Select(dest, sql, args...)
	if err != nil {
		_ = log4go.Error("MFind error=========", err)
		return nil
	}
	return dest
}

/**
插入数据
*/
func MAdd(db string, sql string, args ...interface{}) int64 {
	return execSql("add", db, sql, args...)
}

/**
更新数据
*/
func MUpdate(db string, sql string, args ...interface{}) int64 {
	return execSql("update", db, sql, args...)
}

/**
删除数据
*/
func MDel(db string, sql string, args ...interface{}) int64 {
	return execSql("del", db, sql, args...)
}

func execSql(t string, db string, sql string, args ...interface{}) int64 {
	ret, err := getConn(db).Exec(sql, args...)
	if err != nil {
		_ = log4go.Error("execSql error================", err)
		return 0
	}
	switch t {
	case "del":
		num, _ := ret.RowsAffected()
		return num
	case "update":
		num, _ := ret.RowsAffected()
		return num
	case "add":
		theID, _ := ret.LastInsertId()
		return theID
	}
	return 0
}
