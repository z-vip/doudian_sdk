package order

import (
	"github.com/z-vip/doudian_sdk/unit"
	"reflect"
	"time"
)

// ArgList OrderList方法的参数
type ArgList struct {
	Status    SS           `paramName:"order_status,optional"`
	StartTime string       `paramName:"start_time"`
	EndTime   string       `paramName:"end_time"`
	OrderBy   string       `paramName:"order_by"`
	IsDesc    unit.BoolInt `paramName:"is_desc,optional"`
	Page      uint8        `paramName:"page,optional"`
	Size      uint8        `paramName:"size,optional"`
}

// ArgSearchList OrderList方法的参数
type ArgSearchList struct {
	CombineStatus   string       `paramName:"combine_status,optional"`
	CreateTimeStart int64        `paramName:"create_time_start"`
	CreateTimeEnd   int64        `paramName:"create_time_end"`
	UpdateTimeStart int64        `paramName:"update_time_start"`
	UpdateTimeEnd   int64        `paramName:"update_time_end"`
	OrderBy         string       `paramName:"order_by"`
	IsDesc          unit.BoolInt `paramName:"is_desc,optional"`
	Page            int          `paramName:"page,optional"`
	Size            int          `paramName:"size,optional"`
}

type CombineStatus struct {
	OrderStatus string `paramName:"combine_status,optional"`
	MainStatus  string `paramName:"main_status,optional"`
}

func (a ArgList) HookConvertValue(f reflect.StructField, v reflect.Value) string {
	if f.Type.String() == "time.Time" {
		return v.Interface().(time.Time).Format(unit.TimeYmdHis)
	}
	return ""
}

// ResponseList OrderList方法的响应结果
type ResponseList struct {
	Count uint32   `mapstructure:"count"`
	Total uint32   `mapstructure:"total"`
	List  []Detail `mapstructure:"list"`
}

// ResponseSearchList OrderList方法的响应结果
type ResponseSearchList struct {
	Count         int               `mapstructure:"count"`
	Page          int               `mapstructure:"page"`
	Total         int               `mapstructure:"total"`
	ShopOrderList []ShopOrderDetail `mapstructure:"shop_order_list"`
}
