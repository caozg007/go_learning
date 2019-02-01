/*
@Time : 2019-01-31 16:02
@Author : caozg
@File : 每日必抢
*/
package main

import (
	"bytes"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/tealeg/xlsx"
	_ "github.com/xuri/excelize"
)

var (
	file  *xlsx.File
	sheet *xlsx.Sheet
	cell  *xlsx.Cell
	row   *xlsx.Row
	_err  error
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func querydata() {
	var buffer bytes.Buffer
	begindate := "2017-1-1"
	enddate := "2019-7-10"
	buffer.WriteString("SELECT mae.activity_name '活动名称', ml.goods_name '活动商品', sc.catalog_name '一级类目名称'")
	buffer.WriteString(",ps.vip_price '日常售价',ps.act_price '活动售价', COALESCE(msw.cost_price, '') '商品成本价',")
	buffer.WriteString("ps.act_store_money '活动时店主返利',sum(ml.goods_qty) '活动销量'")
	buffer.WriteString(",sum(ml.goods_qty * ml.sell_price) '活动销售额' ")
	buffer.WriteString("FROM ybl_order.m_order m JOIN ybl_order.m_order_detail ml ON m.order_id = ml.order_id ")
	buffer.WriteString("JOIN ybl_main.marketing_activity_everyday mae ON mae.id = ml.activity_id ")
	buffer.WriteString("JOIN ybl_spu.product_sku ps ON ml.sku_id = ps.sku_id ")
	buffer.WriteString("LEFT JOIN ybl_spu.product_base pb ON ps.product_id = pb.product_id ")
	buffer.WriteString("LEFT JOIN ybl_spu.catalog sc ON sc.id = pb.f_catalog_id ")
	buffer.WriteString("left join ybl_order.m_sup_wh msw on msw.order_detail_id=ml.order_detail_id ")
	buffer.WriteString("WHERE m.order_status IN (3, 4, 5, 6, 7) AND m.order_time >= mae.activity_start_time ")
	buffer.WriteString("AND m.order_time <= mae.activity_end_time AND mae.activity_name REGEXP '首页活动' ")
	buffer.WriteString("AND m.order_time >= str_to_date('" + begindate + "', '%Y-%m-%d %H')  ")
	buffer.WriteString("AND m.order_time < str_to_date('" + enddate + "', '%Y-%m-%d %H')")
	buffer.WriteString("GROUP BY ml.goods_id, ml.activity_id ORDER BY mae.activity_name DESC ")
	file = xlsx.NewFile()
	sheet, _err = file.AddSheet("Sheet1")
	db, _ := sql.Open("mysql", "111:111@(1:3309)/ybl_order?charset=utf8&parseTime=true")
	rows, err := db.Query(buffer.String())
	if err != nil {
		fmt.Printf(err.Error())
	}
	var a, b, c, d, e, f, g, h, i string
	row := sheet.AddRow()
	var item = [9]string{"活动名称", "活动商品", "一级类目名称", "日常售价", "活动售价", "商品成本价", "活动时店主返利", "活动销量", "活动销售额"}
	for j := 0; j < len(item); j++ {
		cell = row.AddCell()
		cell.Value = item[j]
	}
	for rows.Next() {
		err = rows.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
		row = sheet.AddRow()
		row.SetHeightCM(1)
		var content = [9]string{a, b, c, d, e, f, g, h, i}
		for k := 0; k < len(content); k++ {
			cell = row.AddCell()
			cell.Value = content[k]
		}
	}
	_err = file.Save("/Users/caozg/Desktop/BI/每日必抢998.xlsx")

}

func main() {
	querydata()
}
