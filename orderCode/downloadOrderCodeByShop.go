package orderCode

//  DownloadOrderCodeByShop 方法的参数
type ArgDownloadOrderCodeByShop struct {
	OrderId string `paramName:"order_id,optional" json:"order_id,optional"`
}

type DownloadOrderCodeByShopInfo struct {
	IsSuccess    string `json:"is_success" mapstructure:"is_success"`
	ErrorDesc    string `json:"error_desc" mapstructure:"error_desc"`
	DeliveryType int    `json:"delivery_type" mapstructure:"delivery_type"`
	ShipType     int    `json:"ship_type" mapstructure:"ship_type"`
	OrderCode    string `json:"order_code" mapstructure:"order_code"`
}
