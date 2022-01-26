package Crawl

import (
	"fmt"
	"io"
	"net/http"
)

type data struct {
	StatusCode int
	Server     string
}

func Getdata(url string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:67.0) Gecko/20100101 Firefox/67.0")
	resp, err := client.Do(req)
	checkErr(err)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	//body, err := ioutil.ReadAll(resp.Body)
	StatusCode := resp.StatusCode
	fmt.Println(StatusCode)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
