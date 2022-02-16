package Console

import (
	"Warden/Redis"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Ipread() {
	//file, err := ioutil.ReadFile(`Target/ip.txt`)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(file))

	f, err := os.Open("Target/ip.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		fmt.Println(line)
		line = strings.Replace(line, "\n", "", -1)
		Redis.Saveip(line)
	}
}

func Console() {
	//Ipread()
	Redis.Getip()

}
