package Wash

import (
	"Warden/Crawl"
	"Warden/Elasticsearch"
	"fmt"
	"time"
)

//res["HOST"] = host.Addresses[0].Addr
//res["PORT"] = strconv.Itoa(int(port.ID))

//res["Protocol"] = port.Protocol
//res["State"] = port.State.State
//res["ServiceName"] = port.Service.Name

//fmt.Println(data.Title)
//fmt.Println(strconv.Itoa(data.StatusCode))
//fmt.Println(data.Header)

func Mytask(ip string, ch chan int) {
	fmt.Println(ip)
	//ris := [](map[string]string){}
	//res := make(map[string]string)
	address := GetAddress(ip)
	ris := Mynmap(ip)
	if len(ris) > 0 {
		for _, i := range ris {
			//for k, v := range i {
			//	fmt.Printf("%s - %s\n", k, v)
			//}
			//fmt.Println(i["PORT"])
			if i["ServiceName"] == "http" {
				httpres := Crawl.Getdata("http://" + i["HOST"] + ":" + i["PORT"])
				data := Elasticsearch.Data2{HOST: i["HOST"], PORT: i["PORT"], Country: address.Country, Subdivision: address.Subdivision, City: address.City, ISO: address.ISO, Protocol: i["Protocol"], State: i["State"], ServiceName: i["ServiceName"], Title: httpres.Title, StatusCode: httpres.StatusCode, Header: httpres.Header, Date: time.Now().Format("2006-01-02 15:04:05")}
				fmt.Println(data)
				Elasticsearch.Datasave2(data)
			} else if i["ServiceName"] == "https" {
				httpres := Crawl.Getdata("https://" + i["HOST"] + ":" + i["PORT"])
				data := Elasticsearch.Data2{HOST: i["HOST"], PORT: i["PORT"], Country: address.Country, Subdivision: address.Subdivision, City: address.City, ISO: address.ISO, Protocol: i["Protocol"], State: i["State"], ServiceName: i["ServiceName"], Title: httpres.Title, StatusCode: httpres.StatusCode, Header: httpres.Header, Date: time.Now().Format("2006-01-02 15:04:05")}
				fmt.Println(data)
				Elasticsearch.Datasave2(data)
			} else {
				data := Elasticsearch.Data1{HOST: i["HOST"], PORT: i["PORT"], Country: address.Country, Subdivision: address.Subdivision, City: address.City, ISO: address.ISO, Protocol: i["Protocol"], State: i["State"], ServiceName: i["ServiceName"], Date: time.Now().Format("2006-01-02 15:04:05")}
				fmt.Println(data)
				Elasticsearch.Datasave1(data)
			}
			ch <- 1
		}
	} else {
		ch <- -1
	}

}
