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
