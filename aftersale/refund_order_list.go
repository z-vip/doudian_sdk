package aftersale

import (
	"github.com/z-vip/doudian_sdk/unit"
	"reflect"
	"time"
)

// ArgRefundOrderList RefundOrderList方法的参数
type ArgRefundOrderList struct {
	Type      RFD          `paramName:"type,optional"`
	StartTime time.Time    `paramName:"start_time,optional"`
	EndTime   time.Time    `paramName:"end_time,optional"`
	OrderBy   string       `paramName:"order_by"`
	IsDesc    unit.BoolInt `paramName:"is_desc,optional"`
	Page      uint8        `paramName:"page,optional"`
	Size      uint8        `paramName:"size,optional"`
}

func (a ArgRefundOrderList) HookConvertValue(f reflect.StructField, v reflect.Value) string {
	if f.Type.String() == "time.Time" {
		return v.Interface().(time.Time).Format(unit.TimeYmdHis)
	}
	return ""
}

type After struct {
	OrderId                  int64                  `json:"order_id" mapstructure:"order_id" description:"订单号"`
	AaftersaleId             int64                  `json:"aftersale_id" mapstructure:"aftersale_id" description:"售后单号"`
	ApplyTime                string                 `json:"apply_time" mapstructure:"apply_time" description:"申请售后时间"`
	AftersaleType            int                    `json:"aftersale_type" mapstructure:"aftersale_type" description:"售后类型"`
	StatusDeadline           string                 `json:"status_deadline" mapstructure:"status_deadline" description:"超时自动流转截止时间"`
	DeadlineType             int                    `json:"deadline_type" mapstructure:"deadline_type" description:"自动流转类型"`
	AftersaleProcessDesc     string                 `json:"aftersale_process_desc" mapstructure:"aftersale_process_desc" description:"售后流程描述：仅退款，退货退款等"`
	AftersaleStatusDesc      string                 `json:"aftersale_status_desc" mapstructure:"aftersale_status_desc" description:"售后状态描述：待卖家发货，待买家退货等"`
	ReturnStatusDesc         string                 `json:"return_status_desc" mapstructure:"return_status_desc" description:"退货物流状态描述：未发货，已发货等"`
	ReasonDesc               string                 `json:"reason_desc" mapstructure:"reason_desc" description:"售后原因描述：买错了等"`
	PartType                 int                    `json:"part_type" mapstructure:"part_type" description:"全额部分退款"`
	Pid                      int64                  `json:"pid" mapstructure:"pid" description:"店铺订单id（主订单id）"`
	AftersaleRefundTypeDesc  string                 `json:"aftersale_refund_type_desc" mapstructure:"aftersale_refund_type_desc" description:"退款类型描述：线下退款等"`
	RefundType               int                    `json:"refund_type" mapstructure:"refund_type" description:"售后退款类型"`
	RefundStatus             int                    `json:"refund_status" mapstructure:"refund_status" description:"退款状态"`
	AftersaleStatus          int                    `json:"aftersale_status" mapstructure:"aftersale_status" description:"售后状态"`
	PostReceiver             string                 `json:"post_receiver" mapstructure:"post_receiver" description:"买家收件人名"`
	ArbitrateStatus          int                    `json:"arbitrate_status" mapstructure:"arbitrate_status" description:"仲裁状态"`
	UrgeSmsCnt               int                    `json:"urge_sms_cnt" mapstructure:"urge_sms_cnt" description:"剩余的催发货短信次数"`
	AftersaleItems           []Aftersale_items      `json:"aftersale_items" mapstructure:"aftersale_items" description:"售后申请的子订单信息"`
	AftersaleItemsJson       string                 `json:"aftersale_items_json" mapstructure:"aftersale_items" description:"售后申请的子订单信息"`
	AftersaleRecordItems     []AftersaleRecordItems `json:"aftersale_record_items" mapstructure:"aftersale_record_items" description:"售后申请历史记录"`
	AftersaleRecordItemsJson string                 `json:"aftersale_record_items_json" mapstructure:"aftersale_record_items" description:"售后申请历史记录"`
	ProductName              string                 `json:"product_name" mapstructure:"product_name" description:"商品名称"`
	ProductId                int                    `json:"product_id" mapstructure:"product_id" description:"商品id"`
	ProductImg               string                 `json:"product_img" mapstructure:"product_img" description:"商品图片"`
	Num                      int                    `json:"num" mapstructure:"num" description:"商品数量"`
	PayAmount                int                    `json:"pay_amount" mapstructure:"pay_amount" description:"支付金额 单位分"`
	PostAmount               int                    `json:"post_amount" mapstructure:"post_amount" description:"邮费 单位分"`
	RefundAmount             int                    `json:"refund_amount" mapstructure:"refund_amount" description:"退款金额 单位分"`
	RefundPostAmount         int                    `json:"refund_post_amount" mapstructure:"refund_post_amount" description:"退货运费 单位分"`
	AftersaleService         []string               `json:"aftersale_service" mapstructure:"aftersale_service" description:"售后标签：七天无理由，极速退等"`
	AftersaleServiceJson     string                 `json:"aftersale_service_json" mapstructure:"aftersale_service_json" description:"售后标签：七天无理由，极速退等"`
	SkuSpec                  unit.PropertyOPTS      `json:"sku_spec" mapstructure:"sku_spec" description:"商品规格"`
	SkuSpecJson              string                 `json:"sku_spec_json" mapstructure:"sku_spec_json" description:"商品规格"`
	Role                     string                 `json:"role" mapstructure:"role" description:"操作角色，system：系统、service：平台客服、user：用户、shop：商家"`
	OpType                   string                 `json:"op_type" mapstructure:"op_type" description:"操作类型"`
	OpName                   string                 `json:"op_name" mapstructure:"op_name" description:"操作人名称"`
	Conclusion               int                    `json:"conclusion" mapstructure:"conclusion" description:"操作结果，1：同意2:拒绝3:立即退款"`
	OpTime                   string                 `json:"op_time" mapstructure:"op_time" description:"操作时间"`
	Comment                  string                 `json:"comment" mapstructure:"comment" description:"操作意见 原因"`
	Evidence                 []string               `json:"evidence" mapstructure:"evidence" description:"凭证"`
	EvidenceJson             string                 `json:"evidence_json" mapstructure:"evidence_json" description:"凭证"`
	Remark                   string                 `json:"remark" mapstructure:"remark" description:"说明/备注"`
	GotPkg                   int                    `json:"got_pkg" mapstructure:"got_pkg" description:"收到货物说明，0代表未收到货，1代表已收到货"`
	LogisticsCode            string                 `json:"logistics_code" mapstructure:"logistics_code" description:"物流单号"`
	CompanyCode              string                 `json:"company_code" mapstructure:"company_code" description:"物流公司编码"`
	CompanyName              string                 `json:"company_name" mapstructure:"company_name" description:"物流公司名称"`
	AftersaleOrderType       int                    `json:"aftersale_order_type" mapstructure:"aftersale_order_type" description:"售后单对应订单类型"`
}

type Aftersale_items struct {
	ProductName      string            `json:"product_name" mapstructure:"product_name"`
	ProductId        int64             `json:"product_id" mapstructure:"product_id"`
	ProductImg       string            `json:"product_img" mapstructure:"product_img"`
	Num              int               `json:"num" mapstructure:"num"`
	PayAmount        int               `json:"pay_amount" mapstructure:"pay_amount"`
	PostAmount       int               `json:"post_amount" mapstructure:"post_amount"`
	RefundAmount     int               `json:"refund_amount" mapstructure:"refund_amount"`
	RefundPostAmount int               `json:"refund_post_amount" mapstructure:"refund_post_amount"`
	AftersaleService []string          `json:"aftersale_service" mapstructure:"aftersale_service"`
	SkuSpec          unit.PropertyOPTS `json:"sku_spec" mapstructure:"sku_spec"`
	OrderId          string            `json:"order_id" mapstructure:"order_id"`
	CreateTime       string            `json:"create_time" mapstructure:"create_time"`
	PartType         int               `json:"part_type" mapstructure:"part_type"`
}

type AftersaleRecordItems struct {
	AftersaleId        int64    `json:"aftersale_id" mapstructure:"aftersale_id"`
	OrderId            int64    `json:"order_id" mapstructure:"order_id"`
	Role               string   `json:"role" mapstructure:"role"`
	OpType             string   `json:"op_type" mapstructure:"op_type"`
	OpName             string   `json:"op_name" mapstructure:"op_name"`
	Conclusion         int64    `json:"conclusion" mapstructure:"conclusion"`
	OpTime             string   `json:"op_time" mapstructure:"op_time"`
	Comment            string   `json:"comment" mapstructure:"comment"`
	Evidence           []string `json:"evidence" mapstructure:"evidence"`
	Remark             string   `json:"remark" mapstructure:"remark"`
	GotPkg             int64    `json:"got_pkg" mapstructure:"got_pkg"`
	LogisticsCode      string   `json:"logistics_code" mapstructure:"logistics_code"`
	CompanyCode        string   `json:"company_code" mapstructure:"company_code"`
	CompanyName        string   `json:"company_name" mapstructure:"company_name"`
	RefundType         int      `json:"refund_type" mapstructure:"refund_type"`
	AftersaleOrderType int      `json:"aftersale_order_type" mapstructure:"aftersale_order_type"`
}

//结果
type AftersaleInfo struct {
	Page          int     `json:"page" mapstructure:"page"`
	Total         int     `json:"total" mapstructure:"total"`
	Size          int     `json:"size" mapstructure:"size"`
	AftersaleList []After `json:"aftersale_list" mapstructure:"aftersale_list"`
}
