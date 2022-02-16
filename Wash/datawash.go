package Wash

import (
	"Warden/Crawl"
	"fmt"
)

func Mytask(ip string) {
	fmt.Println(ip)
	//ris := [](map[string]string){}
	//res := make(map[string]string)
	ris := Mynmap("42.245.198.22")
	for _, i := range ris {
		//for k, v := range i {
		//	fmt.Printf("%s - %s\n", k, v)
		//}
		if i["Protocol"] == "http" {
			Crawl.Getdata("http://" + i["HOST"] + ":" + i["PORT"])
		}
		//println(i)
	}

}
