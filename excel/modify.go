/*
@Time : 2019-01-31 14:51
@Author : caozg
@File : modify
*/
package main

import (
	"github.com/tealeg/xlsx"
)

func main() {
	excelFileName := "test_write.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		panic(err)
	}
	first := xlFile.Sheets[0]
	row := first.AddRow()
	row.SetHeightCM(1)
	cell := row.AddCell()
	cell.Value = "铁锤"
	cell = row.AddCell()
	cell.Value = "99"

	err = xlFile.Save(excelFileName)
	if err != nil {
		panic(err)
	}
}
