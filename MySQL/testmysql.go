/*
@Time : 2019-01-16 15:33 
@Author : caozg
@File : testmysql
*/
package main

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
//	db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:5900000@(127.0.0.1:3306)/cao?charset=utf8&parseTime=True&loc=Local")
	if ( err != nil){
		println("数据库连接失败")
		print(err)
	}else{
		print("数据库连接成功")
		// 获取第一条记录，按主键排序
		user := print
		db.First(&user)
		//// SELECT * FROM users ORDER BY id LIMIT 1;

		// 获取最后一条记录，按主键排序
		db.Last(&user)
		//// SELECT * FROM users ORDER BY id DESC LIMIT 1;

		// 获取所有记录
		db.Find(&users)
		//// SELECT * FROM users;

		// 使用主键获取记录
		db.First(&user, 10)
		//// SELECT * FROM users WHERE id = 10;
	}

	defer db.Close()
}
