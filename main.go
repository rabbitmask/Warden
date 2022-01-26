package main

import (
	"Warden/Elasticsearch"
	"fmt"
)

func main() {
	//Redis.Saveip("1.1.8.8")
	//Redis.Redis_test()
	Elasticsearch.Estest()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
