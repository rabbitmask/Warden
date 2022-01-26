package main

import (
	"Warden/Elasticsearch"
	"fmt"
)

func main() {
	//Redis.Saveip("1.1.8.8")
	//Redis.Redis_test()
	Elasticsearch.Estest()
	//Crawl.Getdata("http://www.baidu.com")

	//type User struct {
	//	name string
	//	age  int
	//	city string
	//}
	//
	//user := User{name: "rabbit", age: 18}
	//ptr := &user
	//fmt.Print(ptr.name)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
