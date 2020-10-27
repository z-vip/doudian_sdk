package product

import "github.com/lcxking/doudian_sdk/unit"

// ResponseCategory ProductCateProperty方法的响应结果
type ResponseCateProperty struct {
	PropertyID   uint64          `mapstructure:"property_id"`
	PropertyName string          `mapstructure:"property_name"`
	Required     bool            `mapstructure:"required"`
	Options      CatePropertyOPT `mapstructure:"options"`
}

// CatePropertyOPT 属性项
type CatePropertyOPT unit.PropertyOPTS
