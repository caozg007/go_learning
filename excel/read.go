/*
@Time : 2019-01-31 15:44
@Author : caozg
@File : read
*/
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize"
)

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

	xlsx, err := excelize.OpenFile("test_write.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	rows_excel := xlsx.GetRows("Sheet1")
	for _, row := range rows_excel {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
