package product

import "github.com/z-vip/doudian_sdk/unit"

// ResponseCategory ProductCategory方法的响应结果
type ResponseCategory struct {
	ID   unit.ProductCID `mapstructure:"id"`
	Name string          `mapstructure:"name"`
}

type Category struct {
	PropertyId   int64      `json:"property_id"`
	PropertyName string     `json:"property_name"`
	Required     bool       `json:"required"`
	Options      []*Options `json:"options"`
}
type Options struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
