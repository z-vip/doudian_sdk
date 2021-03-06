package spec

import "github.com/z-vip/doudian_sdk/unit"

// ResponseList SpecList方法的响应结果
type ResponseList struct {
	ID   unit.SpecID `mapstructure:"id"`   // 项id
	Name string      `mapstructure:"name"` // 项名称
}
