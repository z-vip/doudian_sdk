package Msg

import "github.com/z-vip/doudian_sdk/unit"

//订单创建消息,当买家下单，系统生成订单时，推送此消息
type TradeCreate struct {
	Pid         uint64   `mapstructure:"p_id"`         // 父订单ID
	SIds        []uint64 `mapstructure:"s_ids"`        // 子订单ID列表
	ShopID      uint64   `mapstructure:"shop_id"`      //  店铺ID
	CreateTime  uint64   `mapstructure:"create_time"`  //  创建时间
	OrderStatus uint8    `mapstructure:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8    `mapstructure:"order_type"`   //订单类型
	Biz         uint8    `mapstructure:"biz"`          //订单业务类型
}

//订单支付/确认消息,
//以下两种场景会推送此消息：
//在线支付订单，当买家付款成功时
//货到付款订单，当商家确认订单时（实际并无支付行为）
type TradePaid struct {
	Pid         uint64     `mapstructure:"p_id"`         // 父订单ID
	SIds        []uint64   `mapstructure:"s_ids"`        // 子订单ID列表
	ShopID      uint64     `mapstructure:"shop_id"`      //  店铺ID
	OrderStatus uint8      `mapstructure:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8      `mapstructure:"order_type"`   //订单类型
	PayType     uint8      `mapstructure:"pay_type"`     //订单支付方式： 0: 货到付款 1: 微信 2: 支付宝
	PayTime     uint64     `mapstructure:"pay_time"`     //1: 在线订单支付时间 2: 货到付款订单确认时间
	PayAmount   unit.Price `mapstructure:"pay_amount"`   //订单实付金额
	Biz         uint8      `mapstructure:"biz"`          //订单业务类型
}

//卖家发货消息
//买家付款后，卖家对所有商品发货完成，且父订单状态为「已发货」时，推送此消息
type TradeSellerShip struct {
	Pid          uint64       `mapstructure:"p_id"`          // 父订单ID
	SIds         []uint64     `mapstructure:"s_ids"`         // 子订单ID列表
	ShopID       uint64       `mapstructure:"shop_id"`       //  店铺ID
	OrderStatus  uint8        `mapstructure:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8        `mapstructure:"order_type"`    //订单类型
	UpdateTime   uint64       `mapstructure:"update_time"`   //订单发货时间
	PayType      uint8        `mapstructure:"pay_type"`      //订单支付方式： 0: 货到付款 1: 微信 2: 支付宝
	PayTime      uint64       `mapstructure:"pay_time"`      //1: 在线订单支付时间 2: 货到付款订单确认时间
	PayAmount    unit.Price   `mapstructure:"pay_amount"`    //订单实付金额
	ReceiverMsg  ReceiverMsg  `mapstructure:"receiver_msg"`  //收货人详细信息
	LogisticsMsg LogisticsMsg `mapstructure:"logistics_msg"` //发货物流信息
	Biz          uint8        `mapstructure:"biz"`           //订单业务类型
}

//收货人详细信息
type ReceiverMsg struct {
	Name string `mapstructure:"name"` //  收货人姓名
	Tel  string `mapstructure:"tel"`  //  收货人手机号
	Addr string `mapstructure:"addr"` //  收货地址
}

//发货物流信息
type LogisticsMsg struct {
	CompanyId     string `mapstructure:"express_company_id"` //  发货快递公司
	LogisticsCode string `mapstructure:"logistics_code"`     //  发货物流单号
}

//交易完成消息
//买家确认收货或系统自动确认收货，且父订单状态变为「已完成」时，推送此消息
type TradeSuccess struct {
	Pid          uint64   `mapstructure:"p_id"`          // 父订单ID
	SIds         []uint64 `mapstructure:"s_ids"`         // 子订单ID列表
	ShopID       uint64   `mapstructure:"shop_id"`       //  店铺ID
	OrderStatus  uint8    `mapstructure:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8    `mapstructure:"order_type"`    //订单类型
	CompleteTime uint64   `mapstructure:"complete_time"` //交易完成时间
	Biz          uint8    `mapstructure:"biz"`           //订单业务类型
}

//发货物流变更消息
//订单已发货，发货物流信息变更时，推送此消息
type TradeLogisticsChanged struct {
	Pid          uint64       `mapstructure:"p_id"`          // 父订单ID
	SIds         []uint64     `mapstructure:"s_ids"`         // 子订单ID列表
	ShopID       uint64       `mapstructure:"shop_id"`       //  店铺ID
	OrderStatus  uint8        `mapstructure:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8        `mapstructure:"order_type"`    //订单类型
	UpdateTime   uint64       `mapstructure:"update_time"`   //发货物流变更时间
	ReceiverMsg  ReceiverMsg  `mapstructure:"receiver_msg"`  //收货人详细信息
	LogisticsMsg LogisticsMsg `mapstructure:"logistics_msg"` //发货物流信息
	Biz          uint8        `mapstructure:"biz"`           //订单业务类型
}

//买家收货信息变更消息
//当买家收货地址被修改时，推送此消息
type TradeAddressChanged struct {
	Pid         uint64      `mapstructure:"p_id"`         // 父订单ID
	SIds        []uint64    `mapstructure:"s_ids"`        // 子订单ID列表
	ShopID      uint64      `mapstructure:"shop_id"`      //  店铺ID
	OrderStatus uint8       `mapstructure:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8       `mapstructure:"order_type"`   //订单类型
	UpdateTime  uint64      `mapstructure:"update_time"`  //发货物流变更时间
	ReceiverMsg ReceiverMsg `mapstructure:"receiver_msg"` //收货人详细信息
	Biz         uint8       `mapstructure:"biz"`          //订单业务类型
}

/**
订单取消消息
订单被取消时推送此消息。取消订单的场景如下：
	货到付款订单且订单状态为「待确认」，买家和商家可取消订单
	货到付款订单且订单状态为「备货中」，买家、商家和平台客服可取消订单
	货到付款订单且订单状态为「已发货」，物流状态为拒收或退回
	在线支付订单且订单状态为「待付款」，买家或平台客服可取消该订单
	在线支付订单且订单状态为「待付款」，超时未支付，订单自动取消
	在线支付订单且订单状态为「备货中」，触发平台风控规则，订单被取消
*/
type TradeCanceled struct {
	Pid          uint64   `mapstructure:"p_id"`          // 父订单ID
	SIds         []uint64 `mapstructure:"s_ids"`         // 子订单ID列表
	ShopID       uint64   `mapstructure:"shop_id"`       //  店铺ID
	OrderStatus  uint8    `mapstructure:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8    `mapstructure:"order_type"`    //订单类型
	CancelTime   uint64   `mapstructure:"cancel_time"`   //订单取消时间
	CancelReason string   `mapstructure:"cancel_reason"` //取消原因
	Biz          uint8    `mapstructure:"biz"`           //订单业务类型
}

////////////////////////////////////////////////////
//以下为售后相关消息
////////////////////////////////////////////////////

//退款消息必定包含的通用信息
type RefundBase struct {
	Sid              uint64 `mapstructure:"s_id"`               // 子订单ID
	Pid              uint64 `mapstructure:"p_id"`               // 父订单ID
	ShopID           uint64 `mapstructure:"shop_id"`            //  店铺ID
	AfterSaleId      uint64 `mapstructure:"aftersale_id"`       //  售后单id
	AfterSaleStatus  uint8  `mapstructure:"aftersale_status"`   //售后状态
	AfterSaleType    uint8  `mapstructure:"aftersale_type"`     //售后类型
	RefundAmount     uint8  `mapstructure:"refund_amount"`      //申请退款的金额（含运费）
	RefundPostAmount uint8  `mapstructure:"refund_post_amount"` //申请退的运费金额
	RefundVoucherNum uint8  `mapstructure:"refund_voucher_num"` //申请退款的卡券的数量
	ReasonCode       uint8  `mapstructure:"reason_code"`        //申请售后原因码
	LatestRecord     string `mapstructure:"latest_record"`      //最近一条操作记录
}

/**
买家发起售后申请消息
买家发起交易逆向申请时，推送此消息。具体场景如下：
	订单未发货，买家申请整单退款时
	订单已发货，买家申请售后仅退款时
	订单已发货，买家申请售后退货时
*/

type RefundTradeCreated struct {
	RefundBase
	ApplyTime uint64 `mapstructure:"apply_time"` //售后申请时间
}

/**
同意退款消息
以下场景会推送此消息：
	买家在发货前申请整单退款，卖家同意退款或超时自动同意退款时
	买家在发货后申请仅退款，卖家同意退款或超时自动同意退款时
	买家在发货后申请申请退货，卖家确认收货或系统超时自动确认收货时
*/
type RefundAgreed struct {
	RefundBase
	AgreeTime uint64 `mapstructure:"agree_time"` //同意退款时间
}

/**
同意退货申请消息
已发货订单，买家申请退货，卖家同意或系统超时同意该退货申请时，推送此消息
*/
type ReturnApplyAgreed struct {
	RefundBase
	Addr      uint8  `mapstructure:"addr"`       //退货地址
	AgreeTime uint64 `mapstructure:"agree_time"` //同意退款时间
}

/**
买家退货给卖家消息
已发货订单，买家退货申请被同意，买家提交退货物流成功时，推送此消息
*/
type BuyerReturnGoods struct {
	RefundBase
	Logistics  uint64 `mapstructure:"logistics"`   //退货物流单号
	ReturnTime uint64 `mapstructure:"return_time"` //退货物流提交时间
}

/**
拒绝退款消息
买家在发货前申请整单退款，卖家拒绝
买家在发货后申请仅退款，卖家拒绝
买家在发货后申请申请退货，买家发货后，卖家拒绝收货
*/
type RefundRefused struct {
	RefundBase
	RefuseTime uint64 `mapstructure:"refuse_time"` //拒绝时间
}

/**
拒绝退货申请消息
当买家在发货后申请退货退款，卖家拒绝申请时，推送此消息
*/
type ReturnApplyRefused struct {
	RefundBase
	RefuseTime uint64 `mapstructure:"refuse_time"` //拒绝时间
}

/**
退款成功消息
当商家同意退款后，实际退款到账时，会推送此消息
*/
type RefundSuccess struct {
	RefundBase
	SuccessTime uint64 `mapstructure:"success_time"` //退款成功时间¬
}

/**
售后关闭消息
当买家取消申请或系统超时机制导致退款取消时，会推送此消息
*/
type RefundClosed struct {
	RefundBase
	CloseTime uint64 `mapstructure:"close_time"` //售后关闭时间
}
