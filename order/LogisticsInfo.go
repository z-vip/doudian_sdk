package order

type LogisticsInfo struct {
	TrackingNo  string        `json:"tracking_no" mapstructure:"tracking_no"`
	Company     string        `json:"company" mapstructure:"company"`
	ShipTime    int           `json:"ship_time" mapstructure:"ship_time"`
	DeliveryId  string        `json:"delivery_id" mapstructure:"delivery_id"`
	CompanyName string        `json:"company_name" mapstructure:"company_name"`
	ProductInfo []ProductInfo `json:"product_info" mapstructure:"product_info"`
}
type ProductInfo struct {
	ProductName  string     `mapstructure:"product_name"`  //商品名称
	Price        int        `mapstructure:"price"`         //商品价格
	OuterSkuId   string     `mapstructure:"outer_sku_id"`  //商家编码
	SkuId        int        `mapstructure:"sku_id"`        //商品skuId
	ProductCount int        `mapstructure:"product_count"` //商品数量
	ProductId    int        `mapstructure:"product_id"`    //商品ID
	SkuOrderId   string     `mapstructure:"sku_order_id"`  //商品单ID
	SkuSpecs     []SkuSpecs `mapstructure:"sku_specs"`     //规格信息
}

type SkuSpecs struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}
