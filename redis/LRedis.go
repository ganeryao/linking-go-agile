package redis

func Init(redisConfig RConfig) {
	initRedis(redisConfig)
}

func RDel(db string, key string) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	conn.Do("del", key)
	defer releaseConn(conn)
}

func RExists(db string, key string) bool {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	rev, _ := conn.Do("exists", key)
	defer releaseConn(conn)
	return string(rev.([]byte)) == "1"
}

func RExpire(db string, key string, time int) {
	//c.radius 即为 Circle 类型对象中的属性
	conn := getConn(db)
	conn.Do("exists", key, time)
	defer releaseConn(conn)
}

func RGet(db string, key string) string {
	conn := getConn(db)
	rev, _ := conn.Do("get", key)
	defer releaseConn(conn)
	return string(rev.([]byte))
}

func RSet(db string, key string, value string) {
	conn := getConn(db)
	conn.Do("set", key, value)
	defer releaseConn(conn)
}

func RSetNX(db string, key string, value string) bool {
	conn := getConn(db)
	rev, _ := conn.Do("SETNX", key, value)
	defer releaseConn(conn)
	return string(rev.([]byte)) == "1"
}

func RIncr(db string, key string) {
	RIncrBy(db, key, 1)
}

func RIncrBy(db string, key string, num int) {
	conn := getConn(db)
	conn.Do("incrby", key, num)
	defer releaseConn(conn)
}
