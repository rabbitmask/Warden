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
	//Crawl.Getdata("http://192.168.20.30")

	Console.Console()

	//gethosts, err := Console.Gethosts("203.93.136.0/22")
	//if err != nil {
	//	return
	//}
	//for _, num := range gethosts {
	//	println(num)

	//Warden_api.Querytest()
	//test()
	//Console.Getip()
	//Redis.Getip()
	//Console.Ipread()
	//Wash.GetAddress("27.211.25.204")
	//address:=GetAddress(ip)
	//Console.Mynmap()

	//Console.Ipread()

	//
	//func checkErr(err error) {
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//
}
