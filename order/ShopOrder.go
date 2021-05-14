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
	ExpShipTime             int             `json:"exp_ship_time" mapstructure:"exp_ship_time"`
	ShipTime                int             `json:"ship_time" mapstructure:"ship_time"`
	SellerRemarkStars       int             `json:"seller_remark_stars" mapstructure:"seller_remark_stars"`
	PostAddr                PostAddr        `json:"post_addr" mapstructure:"post_addr"`
	SkuOrderList            []SkuOrder      `json:"sku_order_list" mapstructure:"sku_order_list"`
	LogisticsInfo           []LogisticsInfo `json:"logistics_info" mapstructure:"logistics_info"`
	PostAddrJson            string          `json:"post_addr_json"`
	SkuOrderListJson        string          `json:"sku_order_list_json"`
	LogisticsInfoJson       string          `json:"logistics_info_json"`
}

type ShopOrderDetail struct {
	ShopOrder
	CampaignInfo        []CampaignInfo  `json:"campaign_info" mapstructure:"campaign_info"`
	PromotionDetail     PromotionDetail `json:"promotion_detail" mapstructure:"promotion_detail"`
	CampaignInfoJson    string          `json:"campaign_info_json"`
	PromotionDetailJson string          `json:"promotion_detail_json"`
}

type PostAddr struct {
	Province    Region `json:"province"`
	City        Region `json:"city"`
	Town        Region `json:"town"`
	Street      Region `json:"street"`
	Detail      string `json:"detail"`
	ExpShipTime int    `json:"exp_ship_time"`
	ShipTime    int    `json:"ship_time"`
}
type Region struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//活动信息
type CampaignInfo struct {
	CampaignId   int `json:"campaign_id"`
	CampaignType int `json:"campaign_type"`
}

type PromotionDetail struct {
	ShopDiscountDetail     ShopDiscountDetail     `json:"shop_discount_detail" mapstructure:"shop_discount_detail"`         //店铺优惠信息
	PlatformDiscountDetail PlatformDiscountDetail `json:"platform_discount_detail" mapstructure:"platform_discount_detail"` //平台优惠信息
	KolDiscountDetail      KolDiscountDetail      `json:"kol_discount_detail" mapstructure:"kol_discount_detail"`           //达人优惠信息
}

type ShopDiscountDetail struct {
	TotalAmount        int `json:"total_amount" mapstructure:"total_amount"`                 //优惠总金额
	CouponAmount       int `json:"coupon_amount" mapstructure:"coupon_amount"`               //券优惠金额
	FullDiscountAmount int `json:"full_discount_amount" mapstructure:"full_discount_amount"` //满减金额
}

//todo PlatformDiscountDetail
type PlatformDiscountDetail struct {
}

//todo KolDiscountDetail
type KolDiscountDetail struct {
}
