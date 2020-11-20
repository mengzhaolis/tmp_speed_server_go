package Model

type Order struct {
	Tid string
	Title string
	Num	string
	Payment string
	PayTime string
}

// 设置Order的表名为`ss_order_products`
func (Order) TableName() string {
  return "ss_order_products"
}