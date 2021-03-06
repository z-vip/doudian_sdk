package msg

import "github.com/z-vip/doudian_sdk/unit"

//订单信息基本结构体
type TradeBase struct {
	Pid         int64   `json:"p_id"`         // 父订单ID
	SIds        []int64 `json:"s_ids"`        // 子订单ID列表
	ShopID      int64   `json:"shop_id"`      //  店铺ID
	CreateTime  int64   `json:"create_time"`  //  创建时间
	OrderStatus int     `json:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   int     `json:"order_type"`   //订单类型
	Biz         int8    `json:"biz"`          //订单业务类型
}

//订单创建消息,当买家下单，系统生成订单时，推送此消息
type TradeCreate struct {
	Pid         uint64   `json:"p_id"`         // 父订单ID
	SIds        []uint64 `json:"s_ids"`        // 子订单ID列表
	ShopID      uint64   `json:"shop_id"`      //  店铺ID
	CreateTime  uint64   `json:"create_time"`  //  创建时间
	OrderStatus uint8    `json:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8    `json:"order_type"`   //订单类型
	Biz         uint8    `json:"biz"`          //订单业务类型
}

//订单支付/确认消息,
//以下两种场景会推送此消息：
//在线支付订单，当买家付款成功时
//货到付款订单，当商家确认订单时（实际并无支付行为）
type TradePaid struct {
	Pid         uint64     `json:"p_id"`         // 父订单ID
	SIds        []uint64   `json:"s_ids"`        // 子订单ID列表
	ShopID      uint64     `json:"shop_id"`      //  店铺ID
	OrderStatus uint8      `json:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8      `json:"order_type"`   //订单类型
	PayType     uint8      `json:"pay_type"`     //订单支付方式： 0: 货到付款 1: 微信 2: 支付宝
	PayTime     uint64     `json:"pay_time"`     //1: 在线订单支付时间 2: 货到付款订单确认时间
	PayAmount   unit.Price `json:"pay_amount"`   //订单实付金额
	Biz         uint8      `json:"biz"`          //订单业务类型
}

//卖家发货消息
//买家付款后，卖家对所有商品发货完成，且父订单状态为「已发货」时，推送此消息
type TradeSellerShip struct {
	Pid          uint64       `json:"p_id"`          // 父订单ID
	SIds         []uint64     `json:"s_ids"`         // 子订单ID列表
	ShopID       uint64       `json:"shop_id"`       //  店铺ID
	OrderStatus  uint8        `json:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8        `json:"order_type"`    //订单类型
	UpdateTime   uint64       `json:"update_time"`   //订单发货时间
	PayType      uint8        `json:"pay_type"`      //订单支付方式： 0: 货到付款 1: 微信 2: 支付宝
	PayTime      uint64       `json:"pay_time"`      //1: 在线订单支付时间 2: 货到付款订单确认时间
	PayAmount    unit.Price   `json:"pay_amount"`    //订单实付金额
	ReceiverMsg  ReceiverMsg  `json:"receiver_msg"`  //收货人详细信息
	LogisticsMsg LogisticsMsg `json:"logistics_msg"` //发货物流信息
	Biz          uint8        `json:"biz"`           //订单业务类型
}

//收货人详细信息
type ReceiverMsg struct {
	Name string `json:"name"` //  收货人姓名
	Tel  string `json:"tel"`  //  收货人手机号
	Addr string `json:"addr"` //  收货地址
}

//发货物流信息
type LogisticsMsg struct {
	CompanyId     string `json:"express_company_id"` //  发货快递公司
	LogisticsCode string `json:"logistics_code"`     //  发货物流单号
}

//交易完成消息
//买家确认收货或系统自动确认收货，且父订单状态变为「已完成」时，推送此消息
type TradeSuccess struct {
	Pid          uint64   `json:"p_id"`          // 父订单ID
	SIds         []uint64 `json:"s_ids"`         // 子订单ID列表
	ShopID       uint64   `json:"shop_id"`       //  店铺ID
	OrderStatus  uint8    `json:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8    `json:"order_type"`    //订单类型
	CompleteTime uint64   `json:"complete_time"` //交易完成时间
	Biz          uint8    `json:"biz"`           //订单业务类型
}

//发货物流变更消息
//订单已发货，发货物流信息变更时，推送此消息
type TradeLogisticsChanged struct {
	Pid          uint64       `json:"p_id"`          // 父订单ID
	SIds         []uint64     `json:"s_ids"`         // 子订单ID列表
	ShopID       uint64       `json:"shop_id"`       //  店铺ID
	OrderStatus  uint8        `json:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8        `json:"order_type"`    //订单类型
	UpdateTime   uint64       `json:"update_time"`   //发货物流变更时间
	ReceiverMsg  ReceiverMsg  `json:"receiver_msg"`  //收货人详细信息
	LogisticsMsg LogisticsMsg `json:"logistics_msg"` //发货物流信息
	Biz          uint8        `json:"biz"`           //订单业务类型
}

//买家收货信息变更消息
//当买家收货地址被修改时，推送此消息
type TradeAddressChanged struct {
	Pid         uint64      `json:"p_id"`         // 父订单ID
	SIds        []uint64    `json:"s_ids"`        // 子订单ID列表
	ShopID      uint64      `json:"shop_id"`      //  店铺ID
	OrderStatus uint8       `json:"order_status"` //父订单状态，订单创建消息的order_status值为"1"
	OrderType   uint8       `json:"order_type"`   //订单类型
	UpdateTime  uint64      `json:"update_time"`  //发货物流变更时间
	ReceiverMsg ReceiverMsg `json:"receiver_msg"` //收货人详细信息
	Biz         uint8       `json:"biz"`          //订单业务类型
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
	Pid          uint64   `json:"p_id"`          // 父订单ID
	SIds         []uint64 `json:"s_ids"`         // 子订单ID列表
	ShopID       uint64   `json:"shop_id"`       //  店铺ID
	OrderStatus  uint8    `json:"order_status"`  //父订单状态，订单创建消息的order_status值为"1"
	OrderType    uint8    `json:"order_type"`    //订单类型
	CancelTime   uint64   `json:"cancel_time"`   //订单取消时间
	CancelReason string   `json:"cancel_reason"` //取消原因
	Biz          uint8    `json:"biz"`           //订单业务类型
}

/**
买家收货信息变更申请消息
商家开启当买家收货地址修改审核后，当发生修改时，推送此消息
*/
type TradeAddressChangeApplied struct {
	TradeBase
	UpdateTime      int64   `json:"update_time"`       //订单业务类型
	PostReceiverMsg PostMsg `json:"post_receiver_msg"` //订单业务类型
	ReceiverMsg     PostMsg `json:"receiver_msg"`      //订单业务类型
}

type PostMsg struct {
	Name string  `json:"name"`
	Tel  string  `json:"tel"`
	Addr Address `json:"addr"`
}

// Address 收货地址
type Address struct {
	Street   unit.Relation `mapstructure:"street" json:"street"`
	Detail   string        `mapstructure:"detail" json:"detail"`
	Province unit.Relation `mapstructure:"province" json:"province"`
	Town     unit.Relation `mapstructure:"town" json:"town"`
}

/**
订单金额修改消息
卖家主动修改货款价格成功，推送此消息
卖家主动修改邮费成功，推送此消息
*/
type TradeAmountChanged struct {
	TradeBase
	TotalAmount int `mapstructure:"total_amount" json:"total_amount"`
	PostAmount  int `mapstructure:"post_amount" json:"post_amount"`
}

/**
订单部分发货消息
买家付款后，卖家对订单中的部分商品发货，且父订单状态为「部分发货」时，推送此消息
*/
type TradePartlySellerShip struct {
	TradeBase
	TotalAmount int `mapstructure:"total_amount" json:"total_amount"`
	PostAmount  int `mapstructure:"post_amount" json:"post_amount"`
}

/**
订单已支付待处理
拼团下单成功，买家完成支付后，但还未成团
普通订单成功，买家完成支付后，触发风控拦截（pending状态一般会持续两小时）
跨境订单支付成功后，等待运营上传身份证
*/
type TradePending struct {
	TradeBase
	TotalAmount int `mapstructure:"total_amount" json:"total_amount"`
	PostAmount  int `mapstructure:"post_amount" json:"post_amount"`
}

////////////////////////////////////////////////////
//以下为售后相关消息
////////////////////////////////////////////////////

//退款消息必定包含的通用信息
type RefundBase struct {
	Sid              uint64     `json:"s_id"`               // 子订单ID
	Pid              uint64     `json:"p_id"`               // 父订单ID
	ShopID           uint64     `json:"shop_id"`            //  店铺ID
	AfterSaleId      uint64     `json:"aftersale_id"`       //  售后单id
	AfterSaleStatus  uint8      `json:"aftersale_status"`   //售后状态
	AfterSaleType    uint8      `json:"aftersale_type"`     //售后类型
	RefundAmount     unit.Price `json:"refund_amount"`      //申请退款的金额（含运费）
	RefundPostAmount unit.Price `json:"refund_post_amount"` //申请退的运费金额
	RefundVoucherNum uint8      `json:"refund_voucher_num"` //申请退款的卡券的数量
	ReasonCode       uint8      `json:"reason_code"`        //申请售后原因码
	LatestRecord     string     `json:"latest_record"`      //最近一条操作记录
	ArbitrateId      string     `json:"arbitrate_id"`       //仲裁单ID
	ArbitrateStatus  int        `json:"arbitrate_status"`   //仲裁状态
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
	ApplyTime uint64 `json:"apply_time"` //售后申请时间
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
	AgreeTime uint64 `json:"agree_time"` //同意退款时间
}

/**
同意退货申请消息
已发货订单，买家申请退货，卖家同意或系统超时同意该退货申请时，推送此消息
*/
type ReturnApplyAgreed struct {
	RefundBase
	Addr      string `json:"addr"`       //退货地址
	AgreeTime uint64 `json:"agree_time"` //同意退款时间
}

/**
买家退货给卖家消息
已发货订单，买家退货申请被同意，买家提交退货物流成功时，推送此消息
*/
type BuyerReturnGoods struct {
	RefundBase
	Logistics  string `json:"logistics"`   //退货物流单号
	ReturnTime uint64 `json:"return_time"` //退货物流提交时间
}

/**
拒绝退款消息
买家在发货前申请整单退款，卖家拒绝
买家在发货后申请仅退款，卖家拒绝
买家在发货后申请申请退货，买家发货后，卖家拒绝收货
*/
type RefundRefused struct {
	RefundBase
	RefuseTime uint64 `json:"refuse_time"` //拒绝时间
}

/**
拒绝退货申请消息
当买家在发货后申请退货退款，卖家拒绝申请时，推送此消息
*/
type ReturnApplyRefused struct {
	RefundBase
	RefuseTime uint64 `json:"refuse_time"` //拒绝时间
}

/**
退款成功消息
当商家同意退款后，实际退款到账时，会推送此消息
*/
type RefundSuccess struct {
	RefundBase
	SuccessTime uint64 `json:"success_time"` //退款成功时间¬
}

/**
售后关闭消息
当买家取消申请或系统超时机制导致退款取消时，会推送此消息
*/
type RefundClosed struct {
	RefundBase
	CloseTime uint64 `json:"close_time"` //售后关闭时间
}

/**
买家发起客服仲裁消息
商家拒绝买家请求，买家可申请平台客服介入、发起客服仲裁，此时推送此消息
*/
type ArbitrateApplied struct {
	RefundBase
	ArbitrateId     string `json:"arbitrate_id"`     //仲裁单ID
	ArbitrateStatus int    `json:"arbitrate_status"` //仲裁状态
}

/**
客服仲裁结果消息
客服仲裁确定结果时，推送此消息
*/
type ArbitrateAudited struct {
	RefundBase
	ArbitrateConclusion int `json:"arbitrate_conclusion"` //仲裁结果： 1. 支持买家 2. 支持商家 3. 支持买家并立即退款
}

/**
买家取消仲裁消息
买家申请客服介入、发起客服仲裁后，而后又取消客服仲裁时，推送此消息
*/
type ArbitrateCancelled struct {
	RefundBase
}

/**
商家上传仲裁凭证消息
买家申请平台客服介入、发起客服仲裁后，客服要求商家上传凭证；商家上传仲裁凭证时，推送此消息
*/
type ArbitrateSubmited struct {
	RefundBase
}

/**
客服要求商家上传凭证消息
买家申请平台客服介入、发起客服仲裁后，客服处理，要求商家上传凭证时，推送此消息
*/
type ArbitrateSubmiting struct {
	RefundBase
}

/**
卖家收到买家换货包裹，确认换货并二次发货消息
已发货订单，买家上传换货物流，卖家确认收货并二次发货时，推送此消息
*/
type ExchangeComfirmed struct {
	RefundBase
}

/**
售后超时时长变更消息
售后超时时间发生变化时，会推送此消息
*/
type ExpirationChange struct {
	RefundBase
	StatusDeadline string `json:"status_deadline"` //变更后的超时时间戳
}

/**
买家修改售后申请消息
买家修改交易逆向申请时，推送此消息
*/
type RefundModified struct {
	RefundBase
	ModifyTime int64 `json:"modify_time"` //售后申请修改时间
}

//开发者在消息推送地址指向的本地服务中，通过解析消息体，可以获取每条消息的具体信息。
type PushMsgRequest struct {
	MsgType int32  `json:"msg_type"` // 表示的是消息类型，1为支付成功消息通知类型，2为卖家授权类型
	Msg     string `json:"msg"`
}

//应用订购支付成功相关信息
type AppOrderInfo struct {
	AppId     uint64        `json:"app_id"`
	OrderInfo PushOrderInfo `json:"order_info"`
}

type PushOrderInfo struct {
	OrderId          uint64      `json:"order_id"`           //订单唯一标识，可作幂等判断依据
	ShopId           uint64      `json:"shop_id"`            //店铺唯一标识
	ServiceStartTime int64       `json:"service_start_time"` //购买服务之后服务的开始生效时间，时间戳，单位：秒
	ServiceEndTime   int64       `json:"service_end_time"`   //购买服务之后服务的结束时间，时间戳，单位：秒
	Status           int32       `json:"status"`             //1：订单待付款，4：订单取消，5：已完成
	Phone            string      `json:"phone"`              //店铺所有者的手机号
	PayAmount        int64       `json:"pay_amount"`         //用户实际支付的金额，以分为单位
	PayTime          int64       `json:"pay_time"`           //支付时间
	OrderCreateTime  int64       `json:"order_create_time"`  //下单时间
	PayType          int32       `json:"pay_type"`           //支付方式
	PushSkuInfo      PushSkuInfo `json:"push_sku_info"`
}

type PushSkuInfo struct {
	Title        string `json:"title"`         //sku名称,所购买版本与时间的结合
	SpecType     int32  `json:"spec_type"`     //规格名称
	SpecValue    string `json:"spec_value"`    //规格名称
	Price        int64  `json:"price"`         //sku的价格，以分为单位
	Duration     int32  `json:"duration"`      //购买时间按照自然月，1个月，3个月，6个月，12个月。
	DurationUnit int32  `json:"duration_unit"` //时长单位, 0:天，1:月，2:年
}

//权相关信息
type AppAuthInfo struct {
	ActionType int32  `json:"action_type"` // action_type=1表示关闭授权
	AppId      uint64 `json:"app_id"`
	ShopId     uint64 `json:"shop_id"`
}
