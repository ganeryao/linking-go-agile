/**
 * @Title  redis操作包
 * @Description redis操作的封装
 * @Author YaoWeiXin
 * @Update 2020/11/20 10:09.
 */
package redis

import (
	"strconv"
)

func Init(redisConfig RConfig) {
	initRedis(redisConfig)
}

func RDel(db string, key string) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("del", key)
}

func RExists(db string, key string) bool {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("exists", key)
	return rev.(int) == 1
}

func RExpire(db string, key string, time int) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("exists", key, time)
}

func RGet(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("get", key)
	if rev == nil {
		return ""
	} else {
		return string(rev.([]byte))
	}
}

func RSet(db string, key string, value string) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("set", key, value)
}

func RSetEX(db string, key string, value string, expire int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SETEX", key, value, expire)
}

func RSetNX(db string, key string, value string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("SETNX", key, value)
	return rev.(int) == 1
}

func RIncr(db string, key string) {
	RIncrBy(db, key, 1)
}

func RIncrBy(db string, key string, num int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("incrby", key, num)
}

func RHDel(db string, key string, field string) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HDEL", key, field)
}

func RHExists(db string, key string, field string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HEXISTS", key, field)
	return rev.(int) == 1
}

func RHGet(db string, key string, field string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HGET", key, field)
	if rev == nil {
		return ""
	} else {
		return string(rev.([]byte))
	}
}

func RHSet(db string, key string, field string, value interface{}) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HSET", key, field, value)
}

func RHSetNX(db string, key string, field string, value string) bool {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HSETNX", key, field, value)
	return rev.(int) == 1
}

func RHGetAll(db string, key string) interface{} {
	value := make(map[string]interface{})
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HGETALL", key)
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		num := len(temp)
		for i := 0; i < num; i += 2 {
			value[temp[i].(string)] = temp[i+1].(interface{})
		}
		return value
	}
}

func RHIncrBy(db string, key string, field string, num int) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HINCRBY", key, field, num)
}

func RHLen(db string, key string) int {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HLEN", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int)
	}
}

func RHMGet(db string, key string, field ...string) interface{} {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range field {
		args = append(args, field[i])
	}
	value := make([]interface{}, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HMGET", args...)
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		for i := range temp {
			value = append(value, temp[i].(interface{}))
		}
		return value
	}
}

func RHMSet(db string, key string, fieldValue ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range fieldValue {
		args = append(args, fieldValue[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("HMSET", args...)
}

func RHValues(db string, key string) interface{} {
	value := make([]interface{}, 0)
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("HVALS", key)
	if rev == nil {
		return nil
	} else {
		temp := rev.([]interface{})
		for i := range temp {
			value = append(value, temp[i].(interface{}))
		}
		return value
	}
}

func RLLen(db string, key string) int {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("LLEN", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int)
	}
}

func RLLPop(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("LPOP", key)
	if rev == nil {
		return ""
	} else {
		return string(rev.([]byte))
	}
}

func RLLPush(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("LPUSH", args...)
}

func RLRPop(db string, key string) string {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("RPOP", key)
	if rev == nil {
		return ""
	} else {
		return string(rev.([]byte))
	}
}

func RLRPush(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("RPUSH", args...)
}

func RSAdd(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SADD", args...)
}

func RSCard(db string, key string) int {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("SCARD", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int)
	}
}

func RSRem(db string, key string, value ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range value {
		args = append(args, value[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("SREM", args...)
}

func RSMembers(db string, key string) interface{} {
	var args = make([]interface{}, 0)
	var scan = 0
	conn := getConn(db)
	defer releaseConn(conn)
	for true {
		rev, _ := conn.Do("SSCAN", key, scan)
		if rev == nil {
			return nil
		} else {
			value := rev.([]interface{})
			scan = value[0].(int)
			list := value[1].([]interface{})
			for i := range list {
				args = append(args, list[i])
			}
		}
		if scan == 0 {
			break
		}
	}
	return args
}

func RZAdd(db string, key string, member string, score float64) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZADD", key, score, member)
}

func RZAdds(db string, key string, arg ...interface{}) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range arg {
		args = append(args, arg[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZADD", args...)
}

func RZCard(db string, key string) int {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("SCARD", key)
	if rev == nil {
		return 0
	} else {
		return rev.(int)
	}
}

func RZIncrBy(db string, key string, member string, score float64) {
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZINCRBY", key, score, member)
}

func RZRange(db string, key string, start int, end int, withScore bool, isRev bool) interface{} {
	withScoreStr := ""
	if withScore {
		withScoreStr = "WITHSCORES"
	}
	conn := getConn(db)
	defer releaseConn(conn)
	command := "ZRANGE"
	if isRev {
		command = "ZREVRANGE"
	}
	rev, _ := conn.Do(command, key, start, end, withScoreStr)
	if rev == nil {
		return nil
	} else {
		return rev
	}
}

func RZRangeByScore(db string, key string, min float64, max float64, withScore bool, isRev bool) interface{} {
	return RZRangeByScoreLimit(db, key, min, max, withScore, isRev, -1, -1)
}

func RZRangeByScoreLimit(db string, key string, min float64, max float64, withScore bool, isRev bool, offset int, count int) interface{} {
	withScoreStr := ""
	if withScore {
		withScoreStr = "WITHSCORES"
	}
	conn := getConn(db)
	defer releaseConn(conn)
	var rev interface{}
	command := "ZRANGEBYSCORE"
	if isRev {
		command = "ZREVRANGEBYSCORE"
	}
	if offset > 0 && count > 0 {
		rev, _ = conn.Do(command, key, min, max, withScoreStr, offset, count)
	} else {
		rev, _ = conn.Do(command, key, min, max, withScoreStr)
	}
	if rev == nil {
		return nil
	} else {
		return rev
	}
}

func RZRank(db string, key string, member string, isRev bool) int {
	conn := getConn(db)
	defer releaseConn(conn)
	command := "ZRANK"
	if isRev {
		command = "ZREVRANK"
	}
	rev, _ := conn.Do(command, key, member)
	if rev == nil {
		return 0
	} else {
		val, err := strconv.Atoi(string(rev.([]byte)))
		if err != nil {
			return 0
		}
		return val
	}
}

func RZRem(db string, key string, member ...string) {
	var args = make([]interface{}, 0)
	args = append(args, key)
	for i := range member {
		args = append(args, member[i])
	}
	conn := getConn(db)
	defer releaseConn(conn)
	conn.Do("ZREM", args)
}

func RZScore(db string, key string, member string) float64 {
	conn := getConn(db)
	defer releaseConn(conn)
	rev, _ := conn.Do("ZSCORE", key, member)
	if rev == nil {
		return 0
	} else {
		val, err := strconv.ParseFloat(string(rev.([]byte)), 64)
		if err != nil {
			return 0
		}
		return val
	}
}
