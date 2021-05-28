package doudian_sdk

import (
	"errors"
	"github.com/z-vip/doudian_sdk/aftersale"
	"github.com/z-vip/doudian_sdk/logistics"
	"github.com/z-vip/doudian_sdk/order"
	"github.com/z-vip/doudian_sdk/orderCode"
	"github.com/z-vip/doudian_sdk/product"
	"github.com/z-vip/doudian_sdk/product/sku"
	"github.com/z-vip/doudian_sdk/product/spec"
	"github.com/z-vip/doudian_sdk/shop"
	"github.com/z-vip/doudian_sdk/trade"
	"github.com/z-vip/doudian_sdk/unit"
	"time"
)

type App struct {
	base         *BaseApp
	AccessToken  string `mapstructure:"access_token"`
	ExpiresIn    int64  `mapstructure:"expires_in"`
	CreatedAt    int64  `mapstructure:"-"`
	RefreshToken string `mapstructure:"refresh_token"`
	Scope        string `mapstructure:"scope"`
	ShopID       uint64 `mapstructure:"shop_id"`
	ShopName     string `mapstructure:"shop_name"`
	Error        error  `mapstructure:"-"`
}

// TokenExpired 判断access_token是否过期
// 由于网络的延时 和 分布式服务器之间的时间可能不是绝对同步, access_token 过期时间留了一个缓冲区
func (a *App) AccessTokenExpired() bool {
	bufTime := 0

	switch {
	case a.ExpiresIn > 60*60:
		bufTime = 60 * 20
	case a.ExpiresIn > 60*30:
		bufTime = 60 * 10
	case a.ExpiresIn > 60*15:
		bufTime = 60 * 5
	case a.ExpiresIn > 60*5:
		bufTime = 60
	case a.ExpiresIn > 60:
		bufTime -= 20
	}

	return time.Now().Unix() >= a.CreatedAt+a.ExpiresIn-int64(bufTime)
}

// RefreshAccessToken 通过 refreshToken 获取新的 access_token.
// https://op.jinritemai.com/docs/guide-docs/9/22
func (a *App) RefreshAccessToken() error {
	if a.RefreshToken == "" {
		return errors.New("RefreshToken 为空")
	}

	app, err := a.base.RefreshAccessToken(a.RefreshToken)
	if err != nil {
		return err
	}
	a.base.accessToken = &app.AccessToken
	a.AccessToken = app.AccessToken
	a.RefreshToken = app.RefreshToken
	a.ExpiresIn = app.ExpiresIn
	a.CreatedAt = app.CreatedAt
	a.Scope = app.Scope
	a.ShopID = app.ShopID
	a.ShopName = app.ShopName
	return nil
}

func (a *App) SetBase(base *BaseApp) {
	a.base = base
}

func (a *App) GetRequestUrl() (uri string) {
	if a.base != nil {
		uri = a.base.RequestUrl
	}
	return
}

// ShopBrandList 获取店铺的已授权品牌列表
// https://op.jinritemai.com/docs/api-docs/13/54
func (a *App) ShopBrandList() {
}

// ProductList 获取商品列表
// https://op.jinritemai.com/docs/api-docs/14/57
func (a *App) ProductList(arg product.ArgList) (product.ResponseList, error) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	if arg.CheckStatus == 0 {
		arg.CheckStatus = product.CheckPass
	}
	var body product.ResponseList
	if err := a.base.NewRequest("product.list", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductDetail 获取商品详细信息
// strID 字符串格式的商品id
// draft 是否从草稿加载 true从草稿加载 false不从草稿加载 默认false
// https://op.jinritemai.com/docs/api-docs/14/56
func (a *App) ProductDetail(ProductStrID unit.ProductID, draft ...bool) (product.ResponseDetail, error) {
	dt := "false"
	if len(draft) == 1 {
		if draft[0] {
			dt = "true"
		}
	}
	var body product.ResponseDetail
	err := a.base.NewRequest("product.detail", ParamMap{"product_id": ProductStrID, "show_draft": dt}, &body)
	if err != nil {
		return body, err
	}
	return body, nil
}

// ProductCategory 获取商品分类列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/58
func (a *App) ProductCategory(id ...unit.ProductCID) ([]product.ResponseCategory, error) {
	cid := unit.CidTOP
	if len(id) == 1 {
		cid = id[0]
	}
	var body []product.ResponseCategory
	if err := a.base.NewRequest("product.getGoodsCategory", ParamMap{"cid": cid}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductCateProperty 根据商品分类获取对应的属性列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/58
func (a *App) ProductCateProperty(cid1, cid2, cid3 unit.ProductCID) ([]product.ResponseCateProperty, error) {
	dat := ParamMap{"first_cid": cid1, "second_cid": cid2, "third_cid": cid3}
	var body []product.ResponseCateProperty
	if err := a.base.NewRequest("product.getCateProperty", dat, &body); err != nil {
		return body, err
	}
	return body, nil
}

// ProductAdd 添加商品
// https://op.jinritemai.com/docs/api-docs/14/59
func (a *App) ProductAdd(arg product.ArgAdd) (product.Product, error) {
	var body product.ResponseDetail
	if err := a.base.NewRequest("product.add", arg, &body); err != nil {
		return product.Product{}, err
	}
	return body.Product, nil
}

// ProductEdit 编辑商品
// 编辑商品的参数虽与ProductAdd共用，但须要使用product.NewArgEdit方法进行实例
// https://op.jinritemai.com/docs/api-docs/14/60
func (a *App) ProductEdit(arg product.ArgAdd) error {
	var body interface{}
	if err := a.base.NewRequest("product.edit", arg, &body); err != nil {
		return err
	}
	if ret, ok := body.(bool); ok && ret {
		return nil
	}
	return errors.New("edit fail")
}

// ProductDel 删除商品
// https://op.jinritemai.com/docs/api-docs/14/61
func (a *App) ProductDel(id unit.ProductID) error {
	if err := a.base.NewRequest("product.del", ParamMap{"product_id": id}, nil); err != nil {
		return err
	}
	return nil
}

// SpecAdd 添加选项规格
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SpecAdd(arg spec.ArgAdd) (spec.ResponseAdd, error) {
	var body spec.ResponseAdd
	if err := a.base.NewRequest("spec.add", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecList 获取选项规格列表
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SpecList() ([]spec.ResponseList, error) {
	var body []spec.ResponseList
	if err := a.base.NewRequest("spec.list", nil, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecDetail 获取选项规格详情
// https://op.jinritemai.com/docs/api-docs/14/63
func (a *App) SpecDetail(id unit.SpecID) (spec.ResponseDetail, error) {
	var body spec.ResponseDetail
	if err := a.base.NewRequest("spec.specDetail", ParamMap{"id": id}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SpecDel 删除选项规格
// https://op.jinritemai.com/docs/api-docs/14/65
func (a *App) SpecDel(id unit.SpecID) error {
	return a.base.NewRequest("spec.del", ParamMap{"id": id}, nil)
}

// SkuAdd 添加SKU
// https://op.jinritemai.com/docs/api-docs/14/64
func (a *App) SkuAdd(arg sku.ArgAdd) (map[uint64]unit.SkuID, error) {
	var body interface{}
	if err := a.base.NewRequest("sku.addAll", arg, &body); err != nil {
		return nil, err
	}
	ret, ok := sku.ResponseAdd{R: body}.Result()
	if ok {
		return ret, nil
	}
	return nil, errors.New("response data unable to parse")
}

// SkuList 获取商品sku列表
// id 分类id，如果不指则获取最顶级
// https://op.jinritemai.com/docs/api-docs/14/82
func (a *App) SkuList(ProductStrID unit.ProductID) ([]sku.ResponseList, error) {
	var body []sku.ResponseList
	if err := a.base.NewRequest("sku.list", ParamMap{"product_id": ProductStrID}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// SkuDetail 获取商品sku详情
// https://op.jinritemai.com/docs/api-docs/14/104
// todo 官方的这个方法有问题，404。可以通过SkuList方法暂时代替
func (a *App) SkuDetail(sid unit.SkuID) (sku.ResponseDetail, error) {
	var body sku.ResponseDetail
	if err := a.base.NewRequest("sku.detail", ParamMap{"sku_id": sid}, &body); err != nil {
		return body, err
	}
	return body, nil
}

// OrderList 订单列表
// https://op.jinritemai.com/docs/api-docs/15/55
func (a *App) OrderList(arg order.ArgList) (order.ResponseList, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("order.list", arg, &body); err != nil {
		return order.ResponseList{}, err
	}
	return body, nil
}

// SearchList 新订单列表
// https://op.jinritemai.com/docs/api-docs/15/464
func (a *App) SearchList(arg order.ArgSearchList) (order.ResponseSearchList, error) {
	var body order.ResponseSearchList
	if err := a.base.NewRequest("order.searchList", arg, &body); err != nil {
		return order.ResponseSearchList{}, err
	}
	return body, nil
}

// OrderDetail 订单详情
// https://op.jinritemai.com/docs/api-docs/15/68
func (a *App) OrderDetail(o unit.Order) (order.Detail, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("order.detail", ParamMap{"order_id": o.GetParentID()}, &body); err != nil {
		return order.Detail{}, err
	}
	//if body.Total != 1 {
	//	return order.Detail{}, errors.New("order total not is 1")
	//}
	if len(body.List) == 0 {
		return order.Detail{}, errors.New("订单不存在")
	}
	return body.List[0], nil
}

// OrderStockUp 确认货到付款订单
// https://op.jinritemai.com/docs/api-docs/15/69
func (a *App) OrderStockUp(o unit.Order) error {
	return a.base.NewRequest("order.stockUp", ParamMap{"order_id": o.GetParentID()}, nil)
}

// OrderCancel 取消货到付款订单
// https://op.jinritemai.com/docs/api-docs/15/72
func (a *App) OrderCancel(o unit.Order, reason string) error {
	return a.base.NewRequest("order.cancel", ParamMap{"order_id": o.GetParentID(), "reason": reason}, nil)
}

// OrderServiceList 获取客服向店铺发起的服务请求列表
// https://op.jinritemai.com/docs/api-docs/15/74
func (a *App) OrderServiceList(arg order.ArgServiceList) (order.ResponseServiceList, error) {
	var body order.ResponseServiceList
	if err := a.base.NewRequest("order.serviceList", arg, &body); err != nil {
		return order.ResponseServiceList{}, err
	}
	return body, nil
}

// OrderReplyService 回复服务请求
// https://op.jinritemai.com/docs/api-docs/15/75
func (a *App) OrderReplyService(id unit.ServiceID, reply string) error {
	return a.base.NewRequest("order.replyService", ParamMap{"id": id, "reply": reply}, nil)
}

// OrderAddOrderRemark 添加订单备注，给订单加旗标
// https://op.jinritemai.com/docs/api-docs/15/141
func (a *App) OrderAddOrderRemark(id unit.OrderID, remark string, IsAddStar string, star string) error {
	arg := ParamMap{"order_id": id, "remark": remark}
	if IsAddStar != "" {
		arg["is_add_star"] = IsAddStar
	}
	if star != "" {
		arg["star"] = star
	}
	return a.base.NewRequest("order.addOrderRemark", arg, nil)
}

// OrderLogisticsAdd 订单发货
// https://op.jinritemai.com/docs/api-docs/16/77
func (a *App) OrderLogisticsAdd(arg order.ArgLogisticsAdd) error {
	return a.base.NewRequest("order.logisticsAdd", arg, nil)
}

// OrderLogisticsEdit 修改发货物流
// https://op.jinritemai.com/docs/api-docs/16/79
func (a *App) OrderLogisticsEdit(arg order.ArgLogisticsAdd) error {
	return a.base.NewRequest("order.logisticsEdit", arg, nil)
}

// AddressProvinceList 获取省列表
// https://op.jinritemai.com/docs/api-docs/16/101
func (a *App) AddressProvinceList() ([]logistics.Province, error) {
	var body []logistics.Province
	if err := a.base.NewRequest("address.provinceList", nil, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// AddressCityList 获取市列表
// https://op.jinritemai.com/docs/api-docs/16/102
func (a *App) AddressCityList(provID uint32) ([]logistics.City, error) {
	var body []logistics.City
	if err := a.base.NewRequest("address.cityList", ParamMap{"province_id": provID}, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// AddressCityList 获取市列表
// https://op.jinritemai.com/docs/api-docs/16/103
func (a *App) AddressAreaList(cityID uint32) ([]logistics.Area, error) {
	var body []logistics.Area
	if err := a.base.NewRequest("address.areaList", ParamMap{"city_id": cityID}, &body); err != nil {
		return nil, err
	}
	return body, nil
}

// RefundOrderList 获取备货中有退款的订单列表
// 在订单发货前，用户能申请退款，但此时只能申请整单退。
// https://op.jinritemai.com/docs/api-docs/17/80
func (a *App) RefundOrderList(arg aftersale.ArgRefundOrderList) (order.ResponseList, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("refund.orderList", arg, &body); err != nil {
		return order.ResponseList{}, err
	}
	return body, nil
}

// RefundShopRefund 商家处理备货中退款申请
// 原始订单状态码				同意退款后			拒绝退款申请后
// 16（退款中-用户申请）		17（退款中-商家同意）	3（已发货）
// https://op.jinritemai.com/docs/api-docs/17/87
func (a *App) RefundShopRefund(arg aftersale.ArgRefundShopRefund) error {
	if arg.Type == aftersale.RSR02 {
		if arg.LogisticsID == 0 || arg.LogisticsCode == "" {
			return errors.New("logistics id & code required")
		}
	}
	return a.base.NewRequest("refund.shopRefund", arg, nil)
}

// RefundOrderList 获取已发货且有售后的订单列表
// https://op.jinritemai.com/docs/api-docs/17/88
func (a *App) AfterSaleOrderList(arg aftersale.ArgAfterSaleOrderList) (order.ResponseList, error) {
	var body order.ResponseList
	if err := a.base.NewRequest("refund.orderList", arg, &body); err != nil {
		return order.ResponseList{}, err
	}
	return body, nil
}

// AfterSaleBuyerReturn 商家处理退货申请
// https://op.jinritemai.com/docs/api-docs/17/90
func (a *App) AfterSaleBuyerReturn(arg aftersale.ArgAfterSaleBuyerReturn) error {
	if arg.Type == aftersale.RSR02 {
		if arg.Comment == "0" || arg.Evidence == "" {
			return errors.New("comment & evidence required")
		}
	}
	return a.base.NewRequest("afterSale.buyerReturn", arg, nil)
}

// AfterSaleFirmReceive 商家确认是否收到退货
// 用户填写退货物流后，商家处理，确认是否收到退货
// 注：货到付款订单，调此接口确认收货，必须上传退款凭证图片。售后状态会变为24（退货成功-商户已退款）
// https://op.jinritemai.com/docs/api-docs/17/91
func (a *App) AfterSaleFirmReceive(arg aftersale.ArgAfterSaleFirmReceive) error {
	if arg.Type == aftersale.RSR02 {
		if arg.Comment == "0" {
			return errors.New("comment required")
		}
	}
	arg.Register = "1"
	arg.Package = "1"
	arg.Facade = "1"
	arg.Function = "1"
	return a.base.NewRequest("afterSale.firmReceive", arg, nil)
}

// AfterSaleUploadCompensation 货到付款订单上传退款凭证
// 货到付款订单，用户申请退货，商家确认收到退货时（final_status=12 or 14时）。如果需要上传退款凭证，需要调此接口！
// https://op.jinritemai.com/docs/api-docs/17/92
func (a *App) AfterSaleUploadCompensation(arg aftersale.ArgAfterSaleUploadCompensation) error {
	return a.base.NewRequest("afterSale.firmReceive", arg, nil)
}

// AfterSaleAddOrderRemark 商家为订单添加售后备注
// 此接口可用于给备货中退款的订单、已发货退货/仅退款的订单，添加售后备注
// https://op.jinritemai.com/docs/api-docs/17/93
func (a *App) AfterSaleAddOrderRemark(id unit.OrderID, remark string) error {
	arg := ParamMap{"order_id": id, "remark": remark}
	return a.base.NewRequest("afterSale.addOrderRemark", arg, nil)
}

// AfterSaleRefundProcessDetail 获取退款详情
// 通过该接口获取退款详细信息
// 1、订单未发货，买家申请整单退款
// 2、订单已发货，买家申请发货后仅退款
// 3、订单已发货，买家申请发货后退货
// https://op.jinritemai.com/docs/api-docs/17/96
func (a *App) AfterSaleRefundProcessDetail(id unit.OrderID) (aftersale.ResponseAfterSaleRefundProcessDetail, error) {
	var body aftersale.ResponseAfterSaleRefundProcessDetail
	err := a.base.NewRequest("afterSale.refundProcessDetail", ParamMap{"order_id": id}, &body)
	return body, err
}

// AfterSaleBuyerRefund 商家发货后仅退款申请
// 订单已发货，用户申请售后仅退款，商家处理
// https://op.jinritemai.com/docs/api-docs/17/131
func (a *App) AfterSaleBuyerRefund(arg aftersale.ArgAfterSaleBuyerRefund) error {
	if arg.Type == aftersale.RSR02 {
		if arg.Comment == "" {
			return errors.New("arg.Comment required")
		}
		if aftersale.CommXEvidence[arg.Comment] && arg.Evidence == "" {
			return errors.New("arg.Evidence required")
		}
	}
	return a.base.NewRequest("afterSale.buyerRefund", arg, nil)
}

// OrderSettle 查询订单账单明细
// https://op.jinritemai.com/docs/api-docs/46/204
func (a *App) OrderSettle(arg order.ArgSettle) (order.ResponseSettle, error) {
	if arg.Size == 0 {
		arg.Size = 10
	}
	var body order.ResponseSettle
	if err := a.base.NewRequest("order.settle", arg, &body); err != nil {
		return body, err
	}
	return body, nil
}

/*
==订单API==
*/
// 订单搜索列表
// https://op.jinritemai.com/docs/api-docs/15/464
func (a *App) OrderSearchList(arg interface{}) (body order.SearchListInfo, err error) {
	err = a.base.RequestApi("order.searchList", arg, &body)
	return
}

// 订单详情
// https://op.jinritemai.com/docs/api-docs/15/465
func (a *App) OrderOrderDetail(orderId string) (body order.OrderDetailInfo, err error) {
	arg := ParamMap{
		"shop_order_id": orderId,
	}
	err = a.base.RequestApi("order.orderDetail", arg, &body)
	return
}

// OrderCodeDownloadOrderCodeByShop 下载bic订单码
// https://op.jinritemai.com/docs/api-docs/51/479
func (a *App) OrderCodeDownloadOrderCodeByShop(arg interface{}) (body orderCode.DownloadOrderCodeByShopInfo, err error) {
	err = a.base.RequestApi("orderCode.downloadOrderCodeByShop", arg, &body)
	return
}

// OrderCodeErpShopBindOrderCode bic流程订单物流发货接口
// https://op.jinritemai.com/docs/api-docs/51/480
func (a *App) OrderCodeErpShopBindOrderCode(arg interface{}) (body string, err error) {
	err = a.base.RequestApi("orderCode.erpShopBindOrderCode", arg, &body)
	return
}

// 售后单列表查询
// https://op.jinritemai.com/docs/api-docs/15/464
func (a *App) RefundListSearch(arg interface{}) (body aftersale.AftersaleInfo, err error) {
	err = a.base.RequestApi("trade.refundListSearch", arg, &body)
	return
}

/*
==店铺API==
*/
//获取店铺后台供商家发布商品的类目
///shop/getShopCategory2 	https://op.jinritemai.com/docs/api-docs/13/234
func (a *App) ShopGetShopCategory(cid string) (body []shop.Category, err error) {
	arg := ParamMap{
		"cid": cid,
	}
	err = a.base.RequestApi("shop.getShopCategory", arg, &body)
	return
}

/**
==订单接口==order==
*/

//获取快递公司列表
//order/logisticsAdd	https://op.jinritemai.com/docs/api-docs/16/389
func (a *App) OrderLogisticsCompanyList() (body []order.LogisticsCompany, err error) {
	err = a.base.RequestApi("order.logisticsCompanyList", nil, &body)
	return
}

//订单地址
//order/AddressAppliedSwitch
func (a *App) OrderAddressAppliedSwitch(arg interface{}) (body string, err error) {
	err = a.base.RequestApi("order.AddressAppliedSwitch", arg, nil)
	return
}

/*
==商品API==sku==
*/

// sku/editCode 	https://op.jinritemai.com/docs/api-docs/14/86
func (a *App) SkuSyncStock(arg interface{}) (err error) {
	err = a.base.RequestApi("sku.syncStock", arg, nil)
	return
}

//修改sku编码
// sku/editCode 	https://op.jinritemai.com/docs/api-docs/14/86
func (a *App) SkuEditCode(arg interface{}) (err error) {
	err = a.base.RequestApi("sku.editCode", arg, nil)
	return
}

///sku/editPrice https://op.jinritemai.com/docs/api-docs/14/84
func (a *App) SkuEditPrice(arg interface{}) (err error) {
	err = a.base.RequestApi("sku.editPrice", arg, nil)
	return
}

/*
==商品API==product==
*/
//添加商品
// /product/addV2 	https://op.jinritemai.com/docs/api-docs/14/249
func (a *App) ProductAddV2(arg interface{}) (body product.AddV2, err error) {
	err = a.base.RequestApi("product.addV2", arg, &body)
	return
}

//修改商品
// /product/editV2 	https://op.jinritemai.com/docs/api-docs/14/250
func (a *App) ProductEditV2(arg interface{}) (err error) {
	err = a.base.RequestApi("product.editV2", arg, nil)
	return
}

//支持针对指定商品上架处理
func (a *App) ProductSetOnline(productId string) (err error) {
	arg := ParamMap{
		"product_id": productId,
	}
	err = a.base.RequestApi("product.setOnline", arg, nil)
	return
}

//支持针对指定商品下架处理
func (a *App) ProductSetOffline(productId string) (err error) {
	arg := ParamMap{
		"product_id": productId,
	}
	err = a.base.RequestApi("product.setOffline", arg, nil)
	return
}

//根据商品分类获取对应的属性列表
//https://op.jinritemai.com/docs/api-docs/14/94		/product/getCateProperty
func (a *App) ProductGetCateProperty(categoryLeafId string) (body []product.Category, err error) {
	arg := ParamMap{
		"category_leaf_id": categoryLeafId,
	}
	err = a.base.RequestApi("product.getCateProperty", arg, &body)
	return
}

/**
trade 售后接口
*/
//支持卖家按照售后单维度查询数据。推荐开发者使用此接口替代/refund/orderList和/afterSale/orderList
//https://op.jinritemai.com/docs/api-docs/17/254	/trade/refundListSearch
func (a *App) TradeRefundListSearch(arg interface{}) (body trade.RefundListSearchInfo, err error) {
	err = a.base.RequestApi("trade.refundListSearch", arg, &body)
	return
}
