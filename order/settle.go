package order

type ArgSettle struct {
	TimeType  string `json:"time_type"`  //时间类型筛选	0:结算时间 1:下单时间
	StartTime string `json:"start_time"` //查询开始时间
	EndTime   string `json:"end_time"`   //查询结束时间,必须大于等于开始时间 开始时间和结束时间跨度最大30天
	OrderId   string `json:"order_id"`   //子订单ID
	ProductId string `json:"product_id"` //商品ID
	PayType   string `json:"pay_type"`   //结算账户：0全部、1微信(升级前)、2微信、3支付宝、4周期打款
	FlowType  string `json:"flow_type"`  //业务类型，不传则默认为0 0:全部 1:广告 2:联盟 3:频道 4:免费
	Page      uint8  `json:"page"`
	Size      uint8  `json:"size"`
}

// ArgList OrderList方法的参数
type SettleList struct {
	Data  []SettleItem `json:"data"`
	Total int64        `json:"total"` //查询账单条数
}

type SettleItem struct {
	OrderInfo   OrderInfo   `json:"order_info"`   //订单相关信息
	SettleInfo  SettleInfo  `json:"settle_info"`  //结算相关信息
	IncomeInfo  IncomeInfo  `json:"income_info"`  //收入相关信息
	OutcomeInfo OutcomeInfo `json:"outcome_info"` //支出相关信息
}

type OrderInfo struct {
	ShopOrderId  string `json:"shop_order_id"`  //父订单ID
	OrderId      string `json:"order_id"`       //子订单ID
	CreateTime   string `json:"create_time"`    //订单创建时间
	ProductId    string `json:"product_id"`     //商品ID
	PhaseOrderNo string `json:"phase_order_no"` //阶段单号
	PhaseCnt     int    `json:"phase_cnt"`      //阶段数量，预售有定金和尾款两个阶段所以值为2，注：普通单该值固定为1
	PhaseId      int    `json:"phase_id"`       //阶段ID，比如预售：1-定金，2-尾款，注：普通单该值固定为1
}

//结果
type SettleInfo struct {
	SettleTime   string  `json:"settle_time"` //结算时间
	PayType      string  `json:"pay_type"`    //结算账户类型
	FlowType     string  `json:"create_time"` //业务类型
	SettleAmount float64 `json:"product_id"`  //实际结算金额
}

type IncomeInfo struct {
	PayAmount      float64 `json:"pay_amount"`      //实际支付金额
	PlatformCoupon int     `json:"platform_coupon"` //平台券补贴金额，若订单发生退款，平台券补贴金额按占订单总额的比例扣除
}
type OutcomeInfo struct {
	OrderRefund         float64 `json:"order_refund"`           //订单退款金额
	PlatformServiceFee  float64 `json:"platform_service_fee"`   //平台服务费，订单实付金额*平台服务费率
	AuthorCommission    float64 `json:"author_commission"`      //达人佣金，计费公式：实收货款*达人佣金，计费基数不含平台券和运费
	GoodLearnChannelFee float64 `json:"good_learn_channel_fee"` //商品渠道费用，计费公式：(订单实付-平台服务费-达人佣金)*渠道分成比例

}
