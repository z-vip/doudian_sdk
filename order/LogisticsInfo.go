package order

type LogisticsInfo struct {
	TrackingNo  string `json:"tracking_no" mapstructure:"tracking_no"`
	Company     string `json:"company" mapstructure:"company"`
	ShipTime    int    `json:"ship_time" mapstructure:"ship_time"`
	DeliveryId  string `json:"delivery_id" mapstructure:"delivery_id"`
	CompanyName string `json:"company_name" mapstructure:"company_name"`
	ProductInfo string `json:"product_info" mapstructure:"product_info"`
}
