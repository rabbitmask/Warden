package Elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func Estest() {
	Client, err := elastic.NewClient(elastic.SetURL("http://192.168.20.31:9200"))
	checkErr(err)
	name := "test111"
	data := `{
	"ip": "1.8.8.8",
	"country": "china",
	"port": "80",
	"date": "2022-01-01"
	"city":"shandong"
	}`
	_, err = Client.Index().Index(name).BodyJson(data).Do(context.Background())
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
