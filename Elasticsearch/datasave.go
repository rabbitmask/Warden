package Elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Data struct {
	ip      string
	port    int
	country string
	city    string
	date    string
}

func Estest() {

	Client, err := elastic.NewClient(elastic.SetURL("http://192.168.20.31:9200"))
	checkErr(err)
	name := "test"
	data := Data{ip: "8.8.8.8", port: 80, country: "china", city: "shandong", date: "2022-01-01"}
	res, err := json.Marshal(data)
	checkErr(err)
	fmt.Println(res)
	_, err = Client.Index().Index(name).BodyJson(res).Do(context.Background())
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
