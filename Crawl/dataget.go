package Crawl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Crawldata struct {
	Title      string
	StatusCode int
	Header     string
}

func Getdata(url string) Crawldata {
	fmt.Println(url + "\tstart")
	url = strings.Replace(url, "\n", "", -1)
	client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36")
	resp, err := client.Do(req)
	var data Crawldata
	if err != nil {
		fmt.Println(err.Error())
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		var title string
		var serverheader string
		if len(body) > 0 {
			tmp := regexp.MustCompile(`<title.*?>(.*?)</title>`).FindAllStringSubmatch(string(body), -1)

			if len(tmp) > 0 {
				title = tmp[0][1]
			}

			serverheader = serverheader + resp.Proto + "\t" + resp.Status + "\n"

			for k, v := range resp.Header {

				serverheader = serverheader + k + "=" + v[0] + "\n"
			}

		}

		data = Crawldata{Title: title, StatusCode: resp.StatusCode, Header: serverheader}
		//fmt.Println(title)
		//fmt.Println(string(body))

		//fmt.Println(resp.Proto)
		//fmt.Println(resp.Status)

		//fmt.Println(data.Title)
		//fmt.Println(strconv.Itoa(data.StatusCode))
		//fmt.Println(data.Header)

	}
	return data
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//return
	}
}
