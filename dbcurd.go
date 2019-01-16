/*
@Time : 2019-01-16 16:05 
@Author : caozg
@File : mysql
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:5900000@(127.0.0.1:3306)/cao?charset=utf8")
	//新增
	stmt1, err := db.Prepare("INSERT age SET age=?,name=?")
	res1, err := stmt1.Exec(32, "yellow")
	affect1, err := res1.RowsAffected()
	fmt.Printf("新增了 %d 条记录\n", affect1)
	defer stmt1.Close()
	//修改
	stmt, err := db.Prepare("update age set age=? where name=?")
	res, err := stmt.Exec(33, "yellow")
	affect, err := res.RowsAffected()
	fmt.Printf("修改了 %d 条记录\n", affect)
	defer stmt.Close()
	// 删除
	stmt, err = db.Prepare("delete from age where name=?")
	res, err = stmt.Exec("caozg")
	delete, err := res.RowsAffected()
	fmt.Printf("删除了 %d 条记录\n", delete)
	//查询
	rows, err := db.Query("SELECT * FROM age")
	for rows.Next() {
		var name string
		var age int
		err = rows.Scan(&name, &age)
		fmt.Printf("姓名: %s ", name)
		fmt.Printf("年龄: %d\n", age)
	}
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
