package Crawl

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

type data struct {
	StatusCode int
	Server     string
}

func Getdata(url string) {
	url = strings.Replace(url, "\n", "", -1)
	client := &http.Client{Timeout: time.Duration(2 * time.Second)}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:67.0) Gecko/20100101 Firefox/67.0")
	resp, err := client.Do(req)

	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		StatusCode := resp.StatusCode
		fmt.Println(StatusCode)
	}

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
}
