/*
@Time : 2019-02-19 14:47
@Author : caozg
@File : test
*/
package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	var m Movie
	//打印结构体的值
	fmt.Print("M初始化前的数据：")
	fmt.Printf("%+v\n", m)
	// 对结构体作初始化
	m = Movie{
		"红高粱",
		9,
	}
	//打印结构体的值
	fmt.Printf("M初始化后的数据：%+v\n", m)
	fmt.Print("M初始化后的内容：")
	fmt.Printf("电影名称：%s,    豆瓣评分:%f", m.Name, m.Rating)
	fmt.Println()

}
