package order

type ShopOrder struct {
	ShopId                  int64           `json:"shop_id" mapstructure:"shop_id"`
	ShopName                string          `json:"shop_name" mapstructure:"shop_name"`
	OpenId                  string          `json:"open_id" mapstructure:"open_id"`
	OrderId                 string          `json:"order_id" mapstructure:"order_id"`
	OrderLevel              int             `json:"order_level" mapstructure:"order_level"`
	Biz                     int             `json:"biz" mapstructure:"biz"`
	BizDesc                 string          `json:"biz_desc" mapstructure:"biz_desc"`
	OrderType               int             `json:"order_type" mapstructure:"order_type"`
	OrderTypeDesc           string          `json:"order_type_desc" mapstructure:"order_type_desc"`
	TradeType               int             `json:"trade_type" mapstructure:"trade_type"`
	TradeTypeDesc           string          `json:"trade_type_desc" mapstructure:"trade_type_desc"`
	OrderStatus             int             `json:"order_status" mapstructure:"order_status"`
	OrderStatusDesc         string          `json:"order_status_desc" mapstructure:"order_status_desc"`
	MainStatus              int             `json:"main_status" mapstructure:"main_status"`
	MainStatusDesc          string          `json:"main_status_desc" mapstructure:"main_status_desc"`
	PayTime                 int             `json:"pay_time" mapstructure:"pay_time"`
	OrderExpireTime         int             `json:"order_expire_time" mapstructure:"order_expire_time"`
	FinishTime              int             `json:"finish_time" mapstructure:"finish_time"`
	CreateTime              int             `json:"create_time" mapstructure:"create_time"`
	UpdateTime              int             `json:"update_time" mapstructure:"update_time"`
	CancelReason            string          `json:"cancel_reason" mapstructure:"cancel_reason"`
	BuyerWords              string          `json:"buyer_words" mapstructure:"buyer_words"`
	SellerWords             string          `json:"seller_words" mapstructure:"seller_words"`
	BType                   int             `json:"b_type" mapstructure:"b_type"`
	BTypeDesc               string          `json:"b_type_desc" mapstructure:"b_type_desc"`
	SubBType                int             `json:"sub_b_type" mapstructure:"sub_b_type"`
	SubBTypeDesc            string          `json:"sub_b_type_desc" mapstructure:"sub_b_type_desc"`
	AppId                   int             `json:"app_id" mapstructure:"app_id"`
	PayType                 int             `json:"pay_type" mapstructure:"pay_type"`
	ChannelPaymentNo        string          `json:"channel_payment_no" mapstructure:"channel_payment_no"`
	OrderAmount             int             `json:"order_amount" mapstructure:"order_amount"`
	PayAmount               int             `json:"pay_amount" mapstructure:"pay_amount"`
	PostAmount              int             `json:"post_amount" mapstructure:"post_amount"`
	PostInsuranceAmount     int             `json:"post_insurance_amount" mapstructure:"post_insurance_amount"`
	ModifyAmount            int             `json:"modify_amount" mapstructure:"modify_amount"`
	ModifyPostAmount        int             `json:"modify_post_amount" mapstructure:"modify_post_amount"`
	PromotionAmount         int             `json:"promotion_amount" mapstructure:"promotion_amount"`
	PromotionShopAmount     int             `json:"promotion_shop_amount" mapstructure:"promotion_shop_amount"`
	PromotionPlatformAmount int             `json:"promotion_platform_amount" mapstructure:"promotion_platform_amount"`
	ShopCostAmount          int             `json:"shop_cost_amount" mapstructure:"shop_cost_amount"`
	PlatformCostAmount      int             `json:"platform_cost_amount" mapstructure:"platform_cost_amount"`
	PromotionTalentAmount   int             `json:"promotion_talent_amount" mapstructure:"promotion_talent_amount"`
	PromotionPayAmount      int             `json:"promotion_pay_amount" mapstructure:"promotion_pay_amount"`
	PostTel                 string          `json:"post_tel" mapstructure:"post_tel"`
	PostReceiver            string          `json:"post_receiver" mapstructure:"post_receiver"`
	PostAddr                string          `json:"post_addr" mapstructure:"post_addr"`
	ExpShipTime             int             `json:"exp_ship_time" mapstructure:"exp_ship_time"`
	ShipTime                int             `json:"ship_time" mapstructure:"ship_time"`
	LogisticsInfo           []LogisticsInfo `json:"logistics_info" mapstructure:"logistics_info"`
	SkuOrderList            []SkuOrder      `json:"sku_order_list" mapstructure:"sku_order_list"`
	SellerRemarkStars       int             `json:"seller_remark_stars" mapstructure:"seller_remark_stars"`
}

type ShopOrderDetail struct {
	ShopOrder
	PromotionDetail string `json:"promotion_detail" mapstructure:"promotion_detail"`
}
