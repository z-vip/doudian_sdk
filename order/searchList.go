package order

// ArgList OrderList方法的参数
type ArgSearchList struct {
	Product             string             `json:"product" paramName:"product"`
	BType               int                `json:"b_type" paramName:"b_type"`
	AfterSaleStatusDesc string             `json:"after_sale_status_desc" paramName:"after_sale_status_desc"`
	TrackingNo          string             `json:"tracking_no" paramName:"tracking_no"`
	PresellType         int                `json:"presell_type" paramName:"presell_type"`
	OrderType           int                `json:"order_type" paramName:"order_type"`
	CreateTimeStart     int                `json:"create_time_start" paramName:"create_time_start"`
	CreateTimeEnd       int                `json:"create_time_end" paramName:"create_time_end"`
	AbnormalOrder       int                `json:"abnormal_order" paramName:"abnormal_order"`
	TradeType           int                `json:"trade_type" paramName:"trade_type"`
	CombineStatus       []ArgCombineStatus `json:"combine_status" paramName:"combine_status"`
	UpdateTimeStart     int                `json:"update_time_start" paramName:"update_time_start"`
	UpdateTimeEnd       int                `json:"update_time_end" paramName:"update_time_end"`
	Size                int                `json:"size" paramName:"size"`
	Page                int                `json:"page" paramName:"page"`
	OrderBy             string             `json:"order_by" paramName:"order_by"`
	OrderAsc            bool               `json:"order_asc" paramName:"order_asc"`
}
type ArgCombineStatus struct {
	OrderStatus string `json:"order_status" paramName:"order_status"`
	MainStatus  string `json:"main_status" paramName:"main_status"`
}

//结果
type SearchListInfo struct {
	Page          int         `json:"page" mapstructure:"page"`
	Total         int         `json:"total" mapstructure:"total"`
	Size          int         `json:"size" mapstructure:"size"`
	ShopOrderList []ShopOrder `json:"shop_order_list" mapstructure:"shop_order_list"`
}
