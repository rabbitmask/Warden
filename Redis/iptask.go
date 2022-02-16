package Redis

import (
	"Warden/Wash"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

//func Redis_cli() redis.Conn {
//	conn, err := redis.Dial("tcp", "192.168.20.31:6379")
//	//checkErr(err)
//	fmt.Println("Redis 连接成功")
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Println("Redis 连接失败")
//		}
//	}(conn)
//
//	err = conn.Send("auth", "wings")
//	if err != nil {
//		fmt.Println("Redis 认证失败")
//	}
//	return conn
//}

//func Redis_test() {
//	conn, err := redis.Dial("tcp", "192.168.20.31:6379")
//	//checkErr(err)
//	fmt.Println("Redis 连接成功")
//	defer func(conn redis.Conn) {
//		err := conn.Close()
//		if err != nil {
//			fmt.Println("Redis 连接失败")
//		}
//	}(conn)
//
//	err = conn.Send("auth", "wings")
//	if err != nil {
//		fmt.Println("Redis 认证失败")
//	}
//	_, err = conn.Do("SET", "8.8.8.8", "1", "EX", "60") // 1 天=86400 秒  30 天=2592000 秒
//	fmt.Println("123")
//	checkErr(err)
//}

func Saveip(ip string) {
	conn, err := redis.Dial("tcp", "192.168.20.31:6379")
	checkErr(err)
	//fmt.Println("Redis 连接成功")
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Redis 连接失败")
		}
	}(conn)

	err = conn.Send("auth", "wings")
	if err != nil {
		fmt.Println("Redis 认证失败")
	}

	var exists int
	exists, err = redis.Int(conn.Do("EXISTS", ip))
	if exists == 0 {
		_, err = conn.Do("SET", ip, 0, "EX", "86400") // 1 天=86400 秒  30 天=2592000 秒
	}
	if err != nil {
		fmt.Println(err.Error())
	} else {
	}
}

func Getip() {
	conn, err := redis.Dial("tcp", "192.168.20.31:6379")
	checkErr(err)
	//fmt.Println("Redis 连接成功")
	defer func(conn redis.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("Redis 连接失败")
		}
	}(conn)

	err = conn.Send("auth", "wings")
	if err != nil {
		fmt.Println("Redis 认证失败")
	}

	val, err := redis.Strings(conn.Do("KEYS", "*"))
	for i, _ := range val {
		//fmt.Println(val[i])
		v, err := redis.Int(conn.Do("GET", val[i]))
		checkErr(err)
		//fmt.Println(v)
		if v == 0 {
			Wash.Mytask(val[i])
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
