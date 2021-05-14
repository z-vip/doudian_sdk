package order

// orderDetail OrderList方法的参数

//结果
type OrderDetailInfo struct {
	ShopOrderDetail ShopOrderDetail `json:"shop_order_detail" mapstructure:"shop_order_detail"`
}
