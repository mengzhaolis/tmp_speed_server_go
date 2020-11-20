package Controller

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	xlsx "github.com/tealeg/xlsx"
	Db "github.com/tmp_speed_server_go/Model"
	// "fmt"
)

type result struct {
	Tid                 string `json:"tid"`
	PaymentSerialNumber string `json:"payment_serial_number"`
	MchId               string `json:"mch_id"`
	Title               string `json:"title"`
	Num                 string `json:"num"`
	Payment             string `json:"payment"`
	PayTime             string `json:"pay_time"`
	Name                string `json:"name"`
	Mobile              string `json:"mobile"`
	ProvinceName        string `json:"province_name"`
	CityName            string `json:"city_name"`
	SubOrderStatus      string `json:"sub_order_status"`
	IsRefundStatus      string `json:"is_refund_status"`
	GmtRefundPay        string `json:"gmt_refund_pay"`
}

// 设置Order的表名为`ss_order_products`
// func (Order) TableName() string {
// 	return "ss_order_products"
// }
/*
* Export 导出excel文件
**/
func Export(c *gin.Context) {
	file := xlsx.NewFile()
	sheet, _ := file.AddSheet("标签")
	// 设置表头
	titleHead := []string{"订单号", "交易号", "商户号", "商品名称", "购买数量", "支付金额", "支付时间", "客户姓名", "手机号", "省份", "城市", "订单状态", "售后状态", "退款时间"}
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleHead {
		cell := titleRow.AddCell()
		cell.Value = v
		//表头字体颜色
		cell.GetStyle().Font.Color = "00FF0000"
		//居中显示
		cell.GetStyle().Alignment.Horizontal = "center"
		cell.GetStyle().Alignment.Vertical = "center"
	}
	// 获取数据库中的内容
	var users []result
	// var user []Db.User
	// data := Db.GetDB().Find(&users)
	data := Db.GetDB().Table("ss_order_products").Select("ss_order_products.tid, ss_order_products.title, ss_order_products.num, ss_order_products.payment, ss_order_products.pay_time as pay_time, ss_uc_receiver_address.name, ss_uc_receiver_address.mobile, ss_uc_receiver_address.province_name, ss_uc_receiver_address.city_name, ss_order_products.sub_order_status, ss_order_products.is_refund_status, ss_sub_orders.payment_serial_number, ss_sub_orders.mch_id, ss_refund_records_details.gmt_refund_pay").Joins("join ss_sub_orders ON ss_sub_orders.tid = ss_order_products.tid").Joins("join ss_uc_receiver_address ON ss_uc_receiver_address.user_id = ss_order_products.buyer_id").Joins("join ss_refund_records_details ON ss_refund_records_details.order_id = ss_order_products.id").Find(&users)
	if data.Error != nil {
		c.JSON(200, gin.H{"code": 200, "status": "error", "message": "暂无数据"})
		return
	}
	// if err != nil {
	// 	c.JSON(200, gin.H{"code": 200, "status": "error", "message": "暂无数据"})
	// 	return
	// }
	// c.JSON(200, gin.H{"code": 200, "status": "success", "message": data})
	// fmt.Println(users)
	// 数据插入到excel
	for _, v := range users {
		row := sheet.AddRow()
		row.WriteStruct(&v, -1)
	}
	// var Filename = "订单导出"
	// 实现下载
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	// c.Writer.Header().Set("Content-Type", "application/octet-stream")
	// disposition := fmt.Sprintf("attachment; filename=\"%s-%s.xlsx\"", Filename, time.Now().Format("2006-01-02 15:04:05"))
	// fmt.Println(disposition)
	// c.Writer.Header().Set("Content-Disposition", disposition)
	_ = file.Write(c.Writer)
}
