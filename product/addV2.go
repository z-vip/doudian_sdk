package product

// ArgAdd ProductAdd方法的参数
// 不建议直接使用这个结构体，因为体内的字段全部是string，无法预知其格式
// 请通过NewArgAdd方法来进行添加参数
type ArgAddV2 struct {
	AssocIds         string    `json:"assoc_ids"`
	BrandId          string    `json:"brand_id"`
	CategoryLeafId   string    `json:"category_leaf_id"`
	ClassQuality     string    `json:"class_quality"`
	Commit           string    `json:"commit"`
	DeliveryDelayDay string    `json:"delivery_delay_day"`
	Description      string    `json:"description"`
	DiscountPrice    string    `json:"discount_price"`
	FirstCid         string    `json:"first_cid"`
	FreightId        string    `json:"freight_id"`
	MarketPrice      string    `json:"market_price"`
	Mobile           string    `json:"mobile"`
	Name             string    `json:"name"`
	OutProductId     string    `json:"out_product_id"`
	OuterProductId   string    `json:"outer_product_id"`
	PayType          string    `json:"pay_type"`
	Pic              string    `json:"pic"`
	PresellDelay     string    `json:"presell_delay"`
	PresellEndTime   string    `json:"presell_end_time"`
	PresellType      string    `json:"presell_type"`
	ProductFormat    string    `json:"product_format"`
	ProductType      string    `json:"product_type"`
	QualityList      []Quality `json:"quality_list"`
	QualityReport    string    `json:"quality_report"`
	RecommendRemark  string    `json:"recommend_remark"`
	ReduceType       string    `json:"reduce_type"`
	Remark           string    `json:"remark"`
	SecondCid        string    `json:"second_cid"`
	SpecName         string    `json:"spec_name"`
	SpecPrices       string    `json:"spec_prices"`
	Specs            string    `json:"specs"`
	Supply7dayReturn string    `json:"supply_7day_return"`
	ThirdCid         string    `json:"third_cid"`
	Weight           string    `json:"weight"`
	WeightUnit       string    `json:"weight_unit"`
}

// 请通过NewArgAdd方法来进行添加参数
type Quality struct {
	QualityKey         string             `json:"quality_key"`
	QualityName        string             `json:"quality_name"`
	QualityAttachments QualityAttachments `json:"quality_attachments"`
}
type QualityAttachments struct {
	MediaType int    `json:"media_type"`
	Url       string `json:"url"`
}

//返回参数
type AddV2 struct {
	ProductId      int64  `json:"product_id"`
	OutProductId   int64  `json:"out_product_id"`
	OuterProductId string `json:"outer_product_id"`
	CreateTime     string `json:"create_time"`
	Sku            []*Sku `json:"sku"`
}

type Sku struct {
	SkuId         int64  `json:"sku_id"`
	OutSkuId      int64  `json:"out_sku_id"`
	OuterSkuId    string `json:"outer_sku_id"`
	Code          string `json:"create_time"`
	SpecDetailId1 int64  `json:"spec_detail_id1"`
	SpecDetailId2 int64  `json:"spec_detail_id2"`
	SpecDetailId3 int64  `json:"spec_detail_id3"`
}
