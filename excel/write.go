/*
@Time : 2019-01-31 14:39
@Author : caozg
@File : write
*/
package main

import (
	"database/sql"
	"fmt"
	"github.com/tealeg/xlsx"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	db, err := sql.Open("mysql", "root:5900000@(127.0.0.1:3306)/cao?charset=utf8")
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
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row, row1, row2 *xlsx.Row
	var cell *xlsx.Cell
	var err2 error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		fmt.Printf(err.Error())
	}
	row = sheet.AddRow()
	row.SetHeightCM(1)
	cell = row.AddCell()
	cell.Value = "姓名"
	cell = row.AddCell()
	cell.Value = "年龄"
	row1 = sheet.AddRow()
	row1.SetHeightCM(1)
	cell = row1.AddCell()
	cell.Value = "狗子"
	cell = row1.AddCell()
	cell.Value = "18"
	row2 = sheet.AddRow()
	row2.SetHeightCM(1)
	cell = row2.AddCell()
	cell.Value = "蛋子"
	cell = row2.AddCell()
	cell.Value = "28"
	err2 = file.Save("test_write.xlsx")
	if err2 != nil {
		fmt.Printf(err2.Error())
	}
}
