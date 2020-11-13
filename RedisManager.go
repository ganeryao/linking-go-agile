package linking_go_agile

import (
	redigo "github.com/gomodule/redigo/redis"
	"time"
)

type RedisConfig struct {
	maxIdle     int
	maxActive   int
	idleTimeout int
	configs     []RedisServerConfig
}

type RedisServerConfig struct {
	name     string
	ip       string
	password string
}

var redisPools map[string]*redigo.Pool
var redisDefaultDb string

func getConn(db string) redigo.Conn {
	return redisPools[db].Get()
}

func releaseConn(conn redigo.Conn) {
	conn.Close()
}

func initRedis(redisConfig RedisConfig) {
	redisPools = make(map[string]*redigo.Pool)
	var i int
	var configs = redisConfig.configs
	for i = 0; i < len(configs); i++ {
		config := configs[i]
		pool := initDb(config, redisConfig.maxIdle, redisConfig.maxActive, redisConfig.idleTimeout)
		redisPools[config.name] = pool
		if i == 0 {
			redisDefaultDb = config.name
		}
	}
}

func initDb(config RedisServerConfig, maxIdle int, maxActive int, idleTimeout int) *redigo.Pool {
	return &redigo.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: time.Duration(idleTimeout) * time.Second,

		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial("tcp", config.ip)
			if err != nil {
				return nil, err
			}
			if config.password != "" {
				if _, err := c.Do("AUTH", config.password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redigo.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
