package order

type SkuOrder struct {
	ShopOrderDetail
	ParentOrderId        string `json:"parent_order_id" mapstructure:"parent_order_id"`
	AfterSaleID          uint64 `json:"aftersale_id" gorm:"column:aftersale_id"`
	SendPay              int    `json:"send_pay" mapstructure:"send_pay"`
	SendPayDesc          string `json:"send_pay_desc" mapstructure:"send_pay_desc"`
	AuthorId             int    `json:"author_id" mapstructure:"author_id"`
	AuthorName           string `json:"author_name" mapstructure:"author_name"`
	ThemeType            string `json:"theme_type" mapstructure:"theme_type"`
	ThemeTypeDesc        string `json:"theme_type_desc" mapstructure:"theme_type_desc"`
	RoomId               int    `json:"room_id" mapstructure:"room_id"`
	ContentId            string `json:"content_id" mapstructure:"content_id"`
	VideoId              string `json:"video_id" mapstructure:"video_id"`
	OriginId             string `json:"origin_id" mapstructure:"origin_id"`
	Cid                  int    `json:"cid" mapstructure:"cid"`
	CBiz                 int    `json:"c_biz" mapstructure:"c_biz"`
	CBizDesc             string `json:"c_biz_desc" mapstructure:"c_biz_desc"`
	PageId               int    `json:"page_id" mapstructure:"page_id"`
	Code                 string `json:"code" mapstructure:"code"`
	LogisticsReceiptTime int    `json:"logistics_receipt_time" mapstructure:"logistics_receipt_time"`
	ConfirmReceiptTime   int    `json:"confirm_receipt_time" mapstructure:"confirm_receipt_time"`
	GoodsType            int    `json:"goods_type" mapstructure:"goods_type"`
	ProductId            int    `json:"product_id" mapstructure:"product_id"`
	SkuId                int    `json:"sku_id" mapstructure:"sku_id"`
	FirstCid             int    `json:"first_cid" mapstructure:"first_cid"`
	SecondCid            int    `json:"second_cid" mapstructure:"second_cid"`
	ThirdCid             int    `json:"third_cid" mapstructure:"third_cid"`
	FourthCid            int    `json:"fourth_cid" mapstructure:"fourth_cid"`
	OutSkuId             string `json:"out_sku_id" mapstructure:"out_sku_id"`
	SupplierId           string `json:"supplier_id" mapstructure:"supplier_id"`
	OutProductId         string `json:"out_product_id" mapstructure:"out_product_id"`
	InventoryType        string `json:"inventory_type" mapstructure:"inventory_type"`
	InventoryTypeDesc    string `json:"inventory_type_desc" mapstructure:"inventory_type_desc"`
	ReduceStockType      int    `json:"reduce_stock_type" mapstructure:"reduce_stock_type"`
	ReduceStockTypeDesc  string `json:"reduce_stock_type_desc" mapstructure:"reduce_stock_type_desc"`
	OriginAmount         int    `json:"origin_amount" mapstructure:"origin_amount" `
	HasTax               bool   `json:"has_tax" mapstructure:"has_tax"`
	ItemNum              int    `json:"item_num" mapstructure:"item_num"`
	SumAmount            int    `json:"sum_amount" mapstructure:"sum_amount"`
	SourcePlatform       string `json:"source_platform" mapstructure:"source_platform"`
	ProductPic           string `json:"product_pic" mapstructure:"product_pic" `
	IsComment            int    `json:"is_comment" mapstructure:"is_comment"`
	ProductName          string `json:"product_name" mapstructure:"product_name"`

	WarehouseIds        []string        `json:"warehouse_ids" mapstructure:"warehouse_ids"`
	WarehouseIdsJson    string          `json:"warehouse_ids_json" mapstructure:"warehouse_ids"`
	OutWarehouseIds     []string        `json:"out_warehouse_ids" mapstructure:"out_warehouse_ids"`
	OutWarehouseIdsJson string          `json:"out_warehouse_ids_json" mapstructure:"out_warehouse_ids"`
	Spec                []Spec          `json:"spec" mapstructure:"spec"`
	SpecJson            string          `json:"spec_json"`
	InventoryList       []InventoryList `json:"inventory_list" mapstructure:"inventory_list"`
	InventoryListJson   string          `json:"inventory_list_json"`
	SkuOrderTagUi       []SkuOrderTagUi `json:"sku_order_tag_ui" mapstructure:"sku_order_tag_ui"`
	SkuOrderTagUiJson   string          `json:"sku_order_tag_ui_json"`
}

type Spec struct {
	Name  string `json:"name" mapstructure:"name"`
	Value string `json:"value" mapstructure:"value"`
}
type InventoryList struct {
	WarehouseId       string `json:"warehouse_id" mapstructure:"warehouse_id"`               //仓id
	OutWarehouseId    string `json:"out_warehouse_id" mapstructure:"out_warehouse_id"`       //外部仓id
	InventoryType     int    `json:"inventory_type" mapstructure:"inventory_type"`           //库存类型，普通库存/区域库存
	InventoryTypeDesc string `json:"inventory_type_desc" mapstructure:"inventory_type_desc"` //库存类型描述
}
type SkuOrderTagUi struct {
	Key       string `json:"key" mapstructure:"key"`               //标签key
	Text      string `json:"text" mapstructure:"text"`             //标签文案
	HoverText string `json:"hover_text" mapstructure:"hover_text"` //该商品需要送到质检中心，质检完成后发给用户    标签备注文案
	TagType   string `json:"tag_type" mapstructure:"tag_type"`     //标签类型，如颜色
	HelpDoc   string `json:"help_doc" mapstructure:"help_doc"`     //帮助文档
	Sort      int    `json:"sort" mapstructure:"sort"`             //排序
}
