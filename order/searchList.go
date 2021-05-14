package order

// ArgList OrderList方法的参数
type ArgSearchList struct {
	AbnormalOrder       int                `json:"abnormal_order" paramName:"abnormal_order"`
	AfterSaleStatusDesc string             `json:"after_sale_status_desc" paramName:"after_sale_status_desc"`
	BType               int                `json:"b_type" paramName:"b_type"`
	CombineStatus       []ArgCombineStatus `json:"combine_status" paramName:"combine_status"`
	CreateTimeEnd       int                `json:"create_time_end" paramName:"create_time_end"`
	CreateTimeStart     int                `json:"create_time_start" paramName:"create_time_start"`
	OrderAsc            bool               `json:"order_asc" paramName:"order_asc"`
	OrderBy             string             `json:"order_by" paramName:"order_by"`
	OrderType           int                `json:"order_type" paramName:"order_type"`
	Page                int                `json:"page" paramName:"page"`
	PresellType         int                `json:"presell_type" paramName:"presell_type"`
	Product             string             `json:"product" paramName:"product"`
	Size                int                `json:"size" paramName:"size"`
	TrackingNo          string             `json:"tracking_no" paramName:"tracking_no"`
	TradeType           int                `json:"trade_type" paramName:"trade_type"`
	UpdateTimeEnd       int                `json:"update_time_end" paramName:"update_time_end"`
	UpdateTimeStart     int                `json:"update_time_start" paramName:"update_time_start"`
}
type ArgCombineStatus struct {
	MainStatus  string `json:"main_status" paramName:"main_status"`
	OrderStatus string `json:"order_status" paramName:"order_status"`
}

//结果
type SearchListInfo struct {
	Page          int         `json:"page" mapstructure:"page"`
	Total         int         `json:"total" mapstructure:"total"`
	Size          int         `json:"size" mapstructure:"size"`
	ShopOrderList []ShopOrder `json:"shop_order_list" mapstructure:"shop_order_list"`
}
