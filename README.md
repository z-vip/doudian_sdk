


# 安装使用

```sh
$ go get -u github.com/lcxking/doudian_sdk
```

## 快速上手
```go
package main

import (
	"fmt"
	"github.com/z-vip/doudian_sdk"
	"github.com/z-vip/doudian_sdk/product/spec"
)

func main() {

	const (
		key    = "6870386******410567"
		secret = "25dd8e74-******a90d3bd"
	)

	var app = doudian_sdk.NewBaseApp(key, secret).NewAccessTokenMust()

	// 按商品模块取得方法集
	product := doudian_sdk.GetProduct(app)

	// 构建参数
	arg := spec.SpecAddArg{Name: "规格参数一"}
	arg.NewSpecBox("容量").Add("100ml").Add("300ml").Add("500ml").Add("1000ml").Done()
	arg.NewSpecBox("颜色").Add("红色").Add("白色").Add("黄色").Add("绿色").Done()
	arg, err := arg.GetArgs()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("请求参数：%+v\n\n", arg)

	// 发起请求
	ret, err := product.SpecAdd(arg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", ret)
}

```