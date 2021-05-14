package order

type SkuOrder struct {
	OrderId                 string          `json:"order_id" mapstructure:"order_id"`
	ParentOrderId           string          `json:"parent_order_id" mapstructure:"parent_order_id"`
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
	BType                   int             `json:"b_type" mapstructure:"b_type"`
	BTypeDesc               string          `json:"b_type_desc" mapstructure:"b_type_desc"`
	SubBType                int             `json:"sub_b_type" mapstructure:"sub_b_type"`
	SubBTypeDesc            string          `json:"sub_b_type_desc" mapstructure:"sub_b_type_desc"`
	SendPay                 int             `json:"send_pay" mapstructure:"send_pay"`
	SendPayDesc             string          `json:"send_pay_desc" mapstructure:"send_pay_desc"`
	AuthorId                int             `json:"author_id" mapstructure:"author_id"`
	AuthorName              string          `json:"author_name" mapstructure:"author_name"`
	ThemeType               string          `json:"theme_type" mapstructure:"theme_type"`
	ThemeTypeDesc           string          `json:"theme_type_desc" mapstructure:"theme_type_desc"`
	AppId                   int             `json:"app_id" mapstructure:"app_id"`
	RoomId                  int             `json:"room_id" mapstructure:"room_id"`
	ContentId               string          `json:"content_id" mapstructure:"content_id"`
	VideoId                 string          `json:"video_id" mapstructure:"video_id"`
	OriginId                string          `json:"origin_id" mapstructure:"origin_id"`
	Cid                     int             `json:"cid" mapstructure:"cid"`
	CBiz                    int             `json:"c_biz" mapstructure:"c_biz"`
	CBizDesc                string          `json:"c_biz_desc" mapstructure:"c_biz_desc"`
	PageId                  int             `json:"page_id" mapstructure:"page_id"`
	PayType                 int             `json:"pay_type" mapstructure:"pay_type"`
	ChannelPaymentNo        string          `json:"channel_payment_no" mapstructure:"channel_payment_no"`
	OrderAmount             int             `json:"order_amount" mapstructure:"order_amount"`
	PayAmount               int             `json:"pay_amount" mapstructure:"pay_amount"`
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
	Code                    string          `json:"code" mapstructure:"code"`
	PostTel                 string          `json:"post_tel" mapstructure:"post_tel"`
	PostReceiver            string          `json:"post_receiver" mapstructure:"post_receiver"`
	PostAddr                PostAddr        `json:"post_addr" mapstructure:"post_addr"`
	ExpShipTime             int             `json:"exp_ship_time" mapstructure:"exp_ship_time"`
	ShipTime                int             `json:"ship_time" mapstructure:"ship_time"`
	LogisticsReceiptTime    int             `json:"logistics_receipt_time" mapstructure:"logistics_receipt_time"`
	ConfirmReceiptTime      int             `json:"confirm_receipt_time" mapstructure:"confirm_receipt_time"`
	GoodsType               int             `json:"goods_type" mapstructure:"goods_type"`
	ProductId               int             `json:"product_id" mapstructure:"product_id"`
	SkuId                   int             `json:"sku_id" mapstructure:"sku_id"`
	Spec                    []Spec          `json:"spec" mapstructure:"spec"`
	FirstCid                int             `json:"first_cid" mapstructure:"first_cid"`
	SecondCid               int             `json:"second_cid" mapstructure:"second_cid"`
	ThirdCid                int             `json:"third_cid" mapstructure:"third_cid"`
	FourthCid               int             `json:"fourth_cid" mapstructure:"fourth_cid"`
	OutSkuId                string          `json:"out_sku_id" mapstructure:"out_sku_id"`
	SupplierId              string          `json:"supplier_id" mapstructure:"supplier_id"`
	OutProductId            string          `json:"out_product_id" mapstructure:"out_product_id"`
	WarehouseIds            []string        `json:"warehouse_ids" mapstructure:"warehouse_ids"`
	OutWarehouseIds         []string        `json:"out_warehouse_ids" mapstructure:"out_warehouse_ids"`
	InventoryType           string          `json:"inventory_type" mapstructure:"inventory_type"`
	InventoryTypeDesc       string          `json:"inventory_type_desc" mapstructure:"inventory_type_desc"`
	ReduceStockType         int             `json:"reduce_stock_type" mapstructure:"reduce_stock_type"`
	ReduceStockTypeDesc     string          `json:"reduce_stock_type_desc" mapstructure:"reduce_stock_type_desc"`
	OriginAmount            int             `json:"origin_amount" mapstructure:"origin_amount"`
	HasTax                  bool            `json:"has_tax" mapstructure:"has_tax"`
	ItemNum                 int             `json:"item_num" mapstructure:"item_num"`
	SumAmount               int             `json:"sum_amount" mapstructure:"sum_amount"`
	SourcePlatform          string          `json:"source_platform" mapstructure:"source_platform"`
	ProductPic              string          `json:"product_pic" mapstructure:"product_pic"`
	IsComment               int             `json:"is_comment" mapstructure:"is_comment"`
	ProductName             string          `json:"product_name" mapstructure:"product_name"`
	InventoryList           []InventoryList `json:"inventory_list" mapstructure:"inventory_list"`
	PromotionDetail         string          `json:"promotion_detail" mapstructure:"promotion_detail"`
	CampaignInfo            string          `json:"campaign_info" mapstructure:"campaign_info"`
	SkuOrderTagUi           string          `json:"sku_order_tag_ui" mapstructure:"sku_order_tag_ui"`
}

type InventoryList struct {
	WarehouseId       string `json:"warehouse_id" mapstructure:"warehouse_id"`               //仓id
	OutWarehouseId    string `json:"out_warehouse_id" mapstructure:"out_warehouse_id"`       //外部仓id
	InventoryType     int    `json:"inventory_type" mapstructure:"inventory_type"`           //库存类型，普通库存/区域库存
	InventoryTypeDesc string `json:"inventory_type_desc" mapstructure:"inventory_type_desc"` //库存类型描述
}

type Spec struct {
	Name  string `json:"name" mapstructure:"name"`
	Value string `json:"value" mapstructure:"value"`
}
