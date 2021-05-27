package trade

type RefundListSearchInfo struct {
	OrderId                 int64    `json:"order_id"`
	AftersaleId             int64    `json:"aftersale_id"`
	ApplyTime               string   `json:"apply_time"`
	AftersaleType           int      `json:"aftersale_type"`
	StatusDeadline          string   `json:"status_deadline"`
	DeadlineType            int      `json:"deadline_type"`
	AftersaleProcessDesc    string   `json:"aftersale_process_desc"`
	AftersaleStatusDesc     string   `json:"aftersale_status_desc"`
	ReturnStatusDesc        string   `json:"return_status_desc"`
	ReasonDesc              string   `json:"reason_desc"`
	PartType                int      `json:"part_type"`
	Pid                     int64    `json:"pid"`
	AftersaleRefundTypeDesc string   `json:"aftersale_refund_type_desc"`
	RefundType              int      `json:"refund_type"`
	RefundStatus            int      `json:"refund_status"`
	AftersaleStatus         int      `json:"aftersale_status"`
	PostReceiver            string   `json:"post_receiver"`
	ArbitrateStatus         int      `json:"arbitrate_status"`
	UrgeSmsCnt              int      `json:"urge_sms_cnt"`
	AftersaleItems          []string `json:"aftersale_items"`
	AftersaleRecordItems    []string `json:"aftersale_record_items"`
	ProductName             string   `json:"product_name"`
	ProductId               int      `json:"product_id"`
	ProductImg              string   `json:"product_img"`
	Num                     int      `json:"num"`
	PayAmount               int      `json:"pay_amount"`
	PostAmount              int      `json:"post_amount"`
	RefundAmount            int      `json:"refund_amount"`
	RefundPostAmount        int      `json:"refund_post_amount"`
	AftersaleService        []string `json:"aftersale_service"`
	SkuSpec                 []string `json:"sku_spec"`
	Role                    string   `json:"role"`
	OpType                  string   `json:"op_type"`
	OpName                  string   `json:"op_name"`
	Conclusion              int      `json:"conclusion"`
	OpTime                  string   `json:"op_time"`
	Comment                 string   `json:"comment"`
	Evidence                []string `json:"evidence"`
	Remark                  string   `json:"remark"`
	GotPkg                  int      `json:"got_pkg"`
	LogisticsCode           string   `json:"logistics_code"`
	CompanyCode             string   `json:"company_code"`
	CompanyName             string   `json:"company_name"`
	AftersaleOrderType      int      `json:"aftersale_order_type"`
}
