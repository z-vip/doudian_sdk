package order

const (
	//查询订单账单接口的时间类型
	STT01 int8 = 1 //结算时间
	STT02 int8 = 2 //下单时间
)

// ArgList OrderList方法的参数
type ArgSettle struct {
	TimeType  int8   `paramName:"time_type"`
	StartTime string `paramName:"start_time"`
	EndTime   string `paramName:"end_time"`
	OrderId   string `paramName:"order_id"`
	ProductId string `paramName:"product_id"`
	PayType   string `json:"pay_type"`
	FlowType  string `json:"flow_type"`
	Page      uint8  `paramName:"page,optional"`
	Size      uint8  `paramName:"size,optional"`
}

type SettleOrderInfo struct {
	OrderId    string `mapstructure:"order_id"`
	CreateTime string `mapstructure:"create_time"`
	ProductId  string `mapstructure:"product_id"`
}

type SettleInfo struct {
	SettleTime   string  `mapstructure:"settle_time"`
	PayType      string  `mapstructure:"pay_type"`
	FlowType     string  `mapstructure:"flow_type"`
	SettleAmount float64 `mapstructure:"settle_amount"`
}

type IncomeInfo struct {
	PayAmount      float64 `mapstructure:"pay_amount"`
	PlatformCoupon float64 `mapstructure:"platform_coupon"`
}

type OutcomeInfo struct {
	RefundAmount        float64 `mapstructure:"refund_amount"`
	PlatformServiceFee  float64 `mapstructure:"platform_service_fee"`
	AuthorCommission    float64 `mapstructure:"author_commission"`
	GoodLearnChannelFee float64 `mapstructure:"good_learn_channel_fee"`
}

type SettleDetail struct {
	OrderInfo   SettleOrderInfo `mapstructure:"order_info"`
	SettleInfo  SettleInfo      `mapstructure:"settle_info"`
	IncomeInfo  IncomeInfo      `mapstructure:"income_info"`
	OutcomeInfo OutcomeInfo     `mapstructure:"outcome_info"`
}

// ResponseSettle OrderSettle方法的响应结果
type ResponseSettle struct {
	Data []SettleDetail `mapstructure:"data"`
}