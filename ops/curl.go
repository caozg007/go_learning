/*
@Time : 2019-01-30 14:26
@Author : caozg
@File : curl
*/

package main

//使用的包
import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
)

//定义脚本的版本
const VERSION = "0.1.0"

//使用 flag 来定义命令输入的参数
var Input_protocol = flag.String("p", "tcp", "The protocol you want to check")

//定义检测 tcp 服务的脚本，用到 net 包
func tcp(url string) int {
	_, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		return 1
	}
}

//定义检测 udp 服务的脚本，用到 net 包
func udp(url string) int {
	_, err := net.Dial("udp", url)
	if err != nil {
		fmt.Println(err)
		return 0
	} else {
		return 1
	}
}

//定义检测 http 服务的脚本，成功返回1，失败返回0
func chttp(url, checkword string) int {
	res, err := http.Get(url)
	if err != nil {
		//如果连接失败返回错误
		log.Fatal(err)
		return 0
	}

	//使用 ioutil 读取得到的响应
	robots, err := ioutil.ReadAll(res.Body)
	//关闭资源
	res.Body.Close()

	//失败返回原因
	if err != nil {
		log.Fatal(err)
		return 0
	}

	//调用 regexp 函数查找 checkword
	word, err := regexp.MatchString(checkword, string(robots))
	if err != nil {
		log.Println(err)
		return 0
	}
	if word {
		log.Printf("The `%s`  find in `%s`", checkword, url)
		return 1
	}
	log.Printf("The `%s`  not find in `%s`", checkword, url)
	return 0
}

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		useage := "使用示例: check -p tcp 192.168.7.26:22 或者 check -p udp 192.168.7.23:123 或者 check -p http http://dockerpool.com  dockerpool"
		fmt.Println(useage)
		return
	}
	p := *Input_protocol
	switch {
	case p == "tcp":
		r := tcp(flag.Args()[0])
		fmt.Println(r)
	case p == "udp":
		r := udp(flag.Args()[0])
		fmt.Println(r)
	case p == "http":
		chttp(flag.Args()[0], flag.Args()[1])
	}
}
