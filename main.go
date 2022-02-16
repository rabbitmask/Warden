package main

import (
	"Warden/Console"
	"fmt"
)

func test() {
	ris := [](map[string]string){}
	res := make(map[string]string)
	res["test"] = "123"
	ris = append(ris, res)

	fmt.Print(ris)

}

func main() {
	//Redis.Saveip("1.1.8.8")
	//Redis.Redis_test()
	//Elasticsearch.Estest()
	//Crawl.Getdata("http://42.245.198.188")
	Console.Console()
	//test()
	//Console.Getip()
	//Redis.Getip()
	//Console.Ipread()
	//Console.GetAddress("27.211.25.204")
	//Console.Mynmap()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
