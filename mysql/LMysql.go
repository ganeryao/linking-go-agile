package mysql

import (
	"github.com/alecthomas/log4go"
	"linking/linking-go-agile/pojo"
)

func Init(mysqlConfig MConfig) {
	initMysql(mysqlConfig)
}

func MFindOne(db string, dest interface{}, sql string, args ...interface{}) interface{} {
	err := getConn(db).Get(dest, sql, args...)
	if err != nil {
		_ = log4go.Error("MFindOne error=========", err)
		return nil
	}
	return dest
}

func MFind(db string, dest interface{}, sql string, args ...interface{}) interface{} {
	err := getConn(db).Select(dest, sql, args...)
	if err != nil {
		_ = log4go.Error("MFind error=========", err)
		return nil
	}
	return dest
}

// 插入数据
func MAdd(db string, po pojo.AbstractPO, sql string, args ...interface{}) int64 {
	ret, err := getConn(db).Exec(sql, args...)
	if err != nil {
		_ = log4go.Error("MAdd1 error================", err)
		return 0
	}
	// 新插入数据的id
	theID, err := ret.LastInsertId()
	if err != nil {
		_ = log4go.Error("MAdd2 error================", err)
		return 0
	}
	po.Id = theID
	return theID
}
