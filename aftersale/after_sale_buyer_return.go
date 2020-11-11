package aftersale

import "github.com/z-vip/doudian_sdk/unit"

type ArgAfterSaleBuyerReturn struct {
	OrderID   unit.OrderID `paramName:"order_id"`
	Type      RSR          `paramName:"type"`
	SmsID     unit.BoolStr `paramName:"sms_id"`
	Comment   Comm         `paramName:"comment,optional"`    // type = 2 时需要选择拒绝原因
	Evidence  string       `paramName:"evidence,optional"`   // type = 2 时需要上传图片凭证
	AddressID string       `paramName:"address_id,optional"` // 商家退货物流收货地址id,不传则使用默认售后收货地址

	//注意：type=1时，可以不入参address_id，直接入参退货地址的详细文本信息,如下：
	ReceiverName     string `paramName:"receiver_name,optional"`     // 退回商品的收货人姓名
	ReceiverTel      string `paramName:"receiver_tel,optional"`      // 退回商品的收货手机号
	ReceiverProvince string `paramName:"receiver_province,optional"` // 退货地址的省（直辖市也必须填，比如北京市）
	ReceiverCity     string `paramName:"receiver_city,optional"`     // 退货地址的市
	ReceiverDistrict string `paramName:"receiver_district,optional"` // 退货地址的区
	ReceiverAddress  string `paramName:"receiver_address,optional"`  // 退货详细地址
}
