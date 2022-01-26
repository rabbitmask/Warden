package main

import (
	"Warden/Redis"
	"fmt"
)

func main() {
	Redis.Saveip("1.1.8.8")
	//Redis.Redis_test()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
