package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type MConfig struct {
	MaxIdle   int
	MaxActive int
	Configs   []MServerConfig
}

type MServerConfig struct {
	Name     string
	User     string
	Password string
	Url      string
}

var pools map[string]*sqlx.DB
var mysqlDefaultDb string

func getConn(db string) *sqlx.DB {
	pool, ok := pools[db]
	if ok {
		return pool
	} else {
		return pools[mysqlDefaultDb]
	}
}

func initMysql(mysqlConfig MConfig) {
	pools = make(map[string]*sqlx.DB)
	var i int
	var configs = mysqlConfig.Configs
	var flag = true
	for i = 0; i < len(configs); i++ {
		config := configs[i]
		pool := initDb(config, mysqlConfig.MaxIdle, mysqlConfig.MaxActive)
		if pool != nil {
			pools[config.Name] = pool
			if flag {
				mysqlDefaultDb = config.Name
				flag = false
			}
		}
	}
}

func initDb(config MServerConfig, maxIdle int, maxActive int) *sqlx.DB {
	dsn := config.User + ":" + config.Password + "@" + config.Url
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		panic(`mysql initDb: db(` + config.Url + `): ` + err.Error())
	}
	db.SetMaxOpenConns(maxActive)
	db.SetMaxIdleConns(maxIdle)
	return db
}
