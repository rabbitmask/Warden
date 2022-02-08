package Elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"time"
)

type Data struct {
	Ip      string
	Port    int
	Country string
	City    string
	Date    time.Time
}

var client *elastic.Client
var host = "http://192.168.20.31:9200"

//数据存储
func Datasave() {
	//使用结构体
	data := Data{Ip: "8.8.8.8", Port: 80, Country: "china", City: "shandong", Date: time.Now()}
	_, err := client.Index().Index("test").BodyJson(data).Do(context.Background())
	checkErr(err)

}

//精确查询
func Query_City() {
	////字段相等
	q := elastic.NewQueryStringQuery("City:Shandong")
	res, err := client.Search("test").Query(q).Do(context.Background())
	if err != nil {
		println(err.Error())
	}
	PrintData(res)
}

//模糊查询
func Query_Body() {
	//短语搜索 搜索about字段中有 rock climbing
	matchPhraseQuery := elastic.NewMatchPhraseQuery("Date", "2022")
	res, err := client.Search("test").Query(matchPhraseQuery).Do(context.Background())
	checkErr(err)
	PrintData(res)
}

func PrintData(res *elastic.SearchResult) {
	var typ Data
	for _, item := range res.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(Data)
		fmt.Printf("%#v\n", t)
	}
}

func Estest() {
	client, _ = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	//Datasave()
	//Query_City()
	Query_Body()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
