package aftersale

import (
	"github.com/z-vip/doudian_sdk/order"
	"github.com/z-vip/doudian_sdk/unit"
)

type ResponseAfterSaleRefundProcessDetail struct {
	OrderInfo         Order             `mapstructure:"order_info"`
	Process           ProcessInfo       `mapstructure:"process_info"`
	AfterSaleInfo     Info              `mapstructure:"aftersale_info"`
	Logs              []Log             `mapstructure:"logs"`
	RefundTotalAmount unit.Price        `mapstructure:"refund_total_amount"`
	RefundPostAmount  unit.Price        `mapstructure:"refund_post_amount"`
	ProcessList       []ProcessInfoList `mapstructure:"process_info_list"`
}

type ProcessInfoList struct {
	ProcessInfo
	//Process ProcessInfo `mapstructure:"process_info"`
	//Logs    []Log       `mapstructure:"process_info"`
}

type ProcessInfo struct {
	ApplyTime       string   `mapstructure:"apply_time"`       // demo 2020-07-11 00:29:20
	Reason          string   `mapstructure:"reason"`           // demo 其他
	TypeDesc        string   `mapstructure:"type_desc"`        // demo 退货退款
	ApplyRemark     string   `mapstructure:"apply_remark"`     // demo 到货后试了衣服，我太高了，所以衣服不合身
	Evidence        []string `mapstructure:"evidence"`         // 凭证图片列表
	LogisticsTime   string   `mapstructure:"logistics_time"`   // 买家填写退货物流时间
	LogisticsCode   string   `mapstructure:"logistics_code"`   // 退货物流单号
	LogisticsName   string   `mapstructure:"logistics_name"`   // 退货快递公司
	AfterSaleID     uint64   `mapstructure:"aftersale_id"`     // 售后单ID
	AfterSaleStatus ASS      `mapstructure:"aftersale_status"` // 售后单状态
	RefundType      uint8    `mapstructure:"refund_type"`      // 表示金额怎么退给买家
	RefundStatus    uint8    `mapstructure:"refund_status"`    // 表示退款到账进度
}

type Info struct {
	AfterSaleType     uint8  `mapstructure:"aftersale_type"`      //售后类型：0售后退货退款，1售后仅退款，2售前仅退款
	AfterSaleTypeText string `mapstructure:"aftersale_type_text"` //售后类型文案
}

type Log struct {
	CreateTime string   `mapstructure:"create_time"` // demo 2020-07-11 00:29:20
	Action     string   `mapstructure:"action"`      // demo 申请退货
	Desc       string   `mapstructure:"desc"`        // demo 退货理由：其他
	Operator   string   `mapstructure:"operator"`    // demo 67939615271
	Evidence   []string `mapstructure:"evidence"`    // 凭证图片列表
}

// Order 子订单信息
type Order struct {
	OrderID     uint64     `mapstructure:"order_id"`     // 订单ID
	PID         uint64     `mapstructure:"pid"`          // 父订单ID
	OrderStatus order.SS   `mapstructure:"order_status"` // 订单状态
	FinalStatus order.SS   `mapstructure:"final_status"` // 子订单状态
	StatusDesc  string     `mapstructure:"status_desc"`  // 退款状态对应的描述文案
	CreateTime  string     `mapstructure:"create_time"`  // 订单创建时间
	ReceiptTime string     `mapstructure:"receipt_time"` // 订单确认收货时间,可能为空字符串
	ComboNum    uint16     `mapstructure:"combo_num"`    // 下单的sku购买数量
	ComboAmount unit.Price `mapstructure:"combo_amount"` // 下单时的sku单价
	TotalAmount unit.Price `mapstructure:"total_amount"` // 下单时sku对应的总价
	PayAmount   unit.Price `mapstructure:"pay_amount"`   // 下单时改单实际支付的金额(sku总价扣除优惠后的)
	ShopID      uint64     `mapstructure:"shop_id"`      // 店铺ID
	ProductID   uint64     `mapstructure:"product_id"`   // 商品id
	ProductName string     `mapstructure:"product_name"` // 商品名称
	ProductImg  string     `mapstructure:"product_img"`  // 商品图片
}
