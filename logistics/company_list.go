package logistics

import "github.com/z-vip/doudian_sdk/unit"

// ResponseLogisticsCompanyList OrderLogisticsCompanyList方法的响应结果
type ResponseLogisticsCompanyList []Company

// LogisticsCompany 快递公司列表
type Company struct {
	ID   unit.CompanyID `mapstructure:"id"`
	Name string         `mapstructure:"name"`
}
