package Elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
)

var client *elastic.Client
var eshost = "http://192.168.20.31:9200"

type Data1 struct {
	HOST        string
	PORT        string
	Country     string
	Subdivision string
	City        string
	ISO         string
	//Coordinates string
	Protocol    string
	State       string
	ServiceName string
	Date        string
}

type Data2 struct {
	HOST        string
	PORT        string
	Country     string
	Subdivision string
	City        string
	ISO         string
	//Coordinates string
	Protocol    string
	State       string
	ServiceName string
	Title       string
	StatusCode  int
	Header      string
	Date        string
}

//数据存储
func Datasave1(data Data1) {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	_, err := client.Index().Index("test").BodyJson(data).Do(context.Background())
	checkErr(err)
}

//数据存储
func Datasave2(data Data2) {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	_, err := client.Index().Index("test").BodyJson(data).Do(context.Background())
	checkErr(err)

}

//精确查询
func Query_City() {
	////字段相等
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	q := elastic.NewQueryStringQuery("City:Shandong")
	res, err := client.Search("test").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintData(res)
}

//精确查询
func Query_IP(HOST string) {
	////字段相等
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	q := elastic.NewQueryStringQuery("HOST:" + HOST)
	res, err := client.Search("test").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintData(res)
}

//精确查询
func Query_PORT(PORT string) {
	////字段相等
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	q := elastic.NewQueryStringQuery("PORT:" + PORT)
	res, err := client.Search("test").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintData(res)
}

//精确查询
func Query_ServiceName(ServiceName string) {
	////字段相等
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	q := elastic.NewQueryStringQuery("ServiceName:" + ServiceName)
	res, err := client.Search("test").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintData(res)
}

//模糊查询
func Query_Body() {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("Date", "2022")
	res, err := client.Search("test").Query(matchPhraseQuery).Size(1000).Do(context.Background())
	checkErr(err)
	PrintData(res)
}

func PrintData(res *elastic.SearchResult) {
	var typ Data2
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Data2)
		fmt.Printf("%#v\n", t)
	}
}

func Estest() {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(eshost))
	Datasave1(Data1{"42.245.198.22", "53", "China", "Chongqing", "Chongqing", "CN", "tcp", "closed", "domain", "2022-02-16 15:39:25"})
	//Query_City()
	//Query_Body()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
}
