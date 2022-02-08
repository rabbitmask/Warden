package main

import (
	"Warden/Crawl"
	"fmt"
)

func main() {
	//Redis.Saveip("1.1.8.8")
	//Redis.Redis_test()
	//Elasticsearch.Estest()
	Crawl.Getdata("http://www.baidu.com")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
