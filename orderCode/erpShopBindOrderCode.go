package orderCode

//  ErpShopBindOrderCode 方法的参数
type ArgErpShopBindOrderCode struct {
	ShopPackageId int                       `paramName:"shop_package_id,optional" json:"shop_package_id,optional"`
	DeliveryType  int                       `paramName:"delivery_type,optional" json:"delivery_type,optional"`
	ShipType      int                       `paramName:"ship_type,optional" json:"ship_type,optional"`
	OrderList     []ArgErpShopBindOrderList `paramName:"order_list,optional" json:"order_list,optional"`
}
type ArgErpShopBindOrderList struct {
	OrderId       int                         `paramName:"order_id,optional" json:"order_id,optional"`
	OrderCode     string                      `paramName:"order_code,optional" json:"order_code,optional"`
	LogisticsCode string                      `paramName:"logistics_code,optional" json:"logistics_code,optional"`
	LogisticsId   int                         `paramName:"logistics_id,optional" json:"logistics_id,optional"`
	OrderDetail   []ArgErpShopBindOrderDetail `paramName:"order_detail,optional" json:"order_detail,optional"`
}
type ArgErpShopBindOrderDetail struct {
	ComboId    int    `paramName:"combo_id,optional" json:"combo_id,optional"`
	ShopQty    int    `paramName:"shop_qty,optional" json:"shop_qty,optional"`
	SkuBarcode string `paramName:"sku_barcode,optional" json:"sku_barcode,optional"`
	UniqueCode string `paramName:"unique_code,optional" json:"unique_code,optional"`
	Lot01      string `paramName:"lot01,optional" json:"lot01,optional"`
}
