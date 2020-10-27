package product

import "github.com/lcxking/doudian_sdk/unit"

// ResponseCategory ProductCategory方法的响应结果
type ResponseCategory struct {
	ID   unit.ProductCID `mapstructure:"id"`
	Name string          `mapstructure:"name"`
}
