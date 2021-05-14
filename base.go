package doudian_sdk

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/z-vip/doudian_sdk/unit"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

// ParamMap 用于包装请求数据
// 值虽然是interface，但最终在请求时都被转成了string
type ParamMap map[string]interface{}

// GatewayURL 抖音小店网关地址
const GatewayURL = "https://openapi-fxg.jinritemai.com"

// SortKeyList 公共参数排序后的字段列表，签名时用到
var SortKeyList = [5]string{
	"app_key",
	"method",
	"param_json",
	"timestamp",
	"v",
}

// BaseApp 应用的基础配置
type BaseApp struct {
	Key         string
	Secret      string
	accessToken *string
	gatewayURL  string
	RequestUrl  string
}

//NewApp 实例化应用 (appId, appSecret, accessToken, refreshToken)
func NewApp(appId, appSecret, accessToken, refreshToken string) *App {
	app := App{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	base := NewBaseApp(appId, appSecret)
	base.accessToken = &app.AccessToken
	app.base = base
	return &app
}

// NewBaseApp 实例化基础应用
func NewBaseApp(k, s string) *BaseApp {
	return &BaseApp{Key: k, Secret: s, gatewayURL: GatewayURL}
}

// SetAccessToken 重置抖音小店access_token
func (b *BaseApp) SetAccessToken(t string) *BaseApp {
	b.accessToken = &t
	return b
}

// SetGatewayURL 重置抖音小店网关地址
func (b *BaseApp) SetGatewayURL(u string) *BaseApp {
	b.gatewayURL = u
	return b
}

// NewAccessToken 获权AccessToken
// NewApp和NewAccessToken不是同一个对象的实例 该方法将创建新的app
// https://op.jinritemai.com/docs/guide-docs/9/21
func (b *BaseApp) NewAccessToken(t ...string) (*App, error) {
	app := App{}
	if len(t) == 0 {
		body := url.Values{}
		body.Add("app_id", b.Key)
		body.Add("app_secret", b.Secret)
		body.Add("grant_type", "authorization_self")
		resp, err := http.Get(GatewayURL + "/oauth2/access_token?" + body.Encode())
		if err != nil {
			return nil, err
		}
		var ret BaseResp
		if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
			return nil, err
		}
		if err := mapstructure.Decode(ret.Data, &app); err != nil {
			return nil, err
		}

		if ret.ErrNo != 0 {
			return nil, errors.New(ret.Message)
		}

		b.accessToken = &app.AccessToken
	} else {
		app.AccessToken = t[0]
		b.accessToken = &app.AccessToken
	}
	app.base = b
	return &app, nil
}

// NewAccessTokenMust 获权AccessToken
// 同NewAccessToken，只不过error信息存储至对象内的Error属性
// https://op.jinritemai.com/docs/guide-docs/9/21
func (b *BaseApp) NewAccessTokenMust(t ...string) *App {
	app, err := b.NewAccessToken(t...)
	if err != nil {
		return &App{Error: err}
	}
	return app
}

type BaseResp struct {
	Data    interface{} `json:"data"`
	ErrNo   int         `json:"err_no"`
	Message string      `json:"message"`
}

// ToParamMap 将任意struct转换为成ParamMap
// paramName "-" 这个字段将被忽略，须要注意的是如果字段是bool，那么将被转换成字符串
// 类型的"true"和"false"
func ToParamMap(data interface{}, ret ...*ParamMap) ParamMap {
	var (
		r ParamMap // 最终结果
		t reflect.Type
		v reflect.Value
	)
	if len(ret) == 0 {
		r = ParamMap{} // 非递归
	} else {
		r = *ret[0] // 递归时 将以指针式进行赋值
	}
	if val, ok := data.(reflect.Value); ok {
		// 递归
		t = val.Type()
		v = val
	} else {
		// 非递归
		t = reflect.TypeOf(data)
		v = reflect.ValueOf(data)
	}
	// HookConvertValue
	var Func1has bool
	var Func1 reflect.Value
	if _, has := t.MethodByName("HookConvertValue"); has {
		Func1has = has
		Func1 = v.MethodByName("HookConvertValue")
	}
	// HookSkipCheck
	var Func2has bool
	var Func2 reflect.Value
	if _, has := t.MethodByName("HookSkipCheck"); has {
		Func2has = has
		Func2 = v.MethodByName("HookSkipCheck")
	}
	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)       // 字段
		x := v.Field(i)       // 值
		fs := f.Type.String() // 字段类型 字符串
		ff := reflect.ValueOf(f)
		// tag 结构体后的标记 n标记名称 o标记参数
		tag := f.Tag.Get("paramName")
		if tag == "-" {
			continue
		}
		n := ""
		o := ""
		if strings.Index(tag, unit.SPE3) == -1 {
			n = tag
		} else {
			xx := strings.Split(tag, unit.SPE3)
			n = xx[0]
			o = xx[1]
		}
		if n == "" {
			n = strings.ToLower(f.Name)
		}
		val := ""
		switch x.Kind() {
		// int
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val = strconv.FormatInt(x.Int(), 10)
		// uint
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			val = strconv.FormatUint(x.Uint(), 10)
		// float
		case reflect.Float32, reflect.Float64:
			val = strconv.FormatFloat(x.Float(), 'f', -1, 64)
		// string
		case reflect.String:
			val = x.String()
		// struct
		case reflect.Struct:
			if f.Name == f.Type.Name() {
				ToParamMap(x, &r)
			} else {
				if Func1has {
					val = Func1.Call([]reflect.Value{ff, reflect.ValueOf(x)})[0].String()
				}
			}
		// bool
		case reflect.Bool:
			if x.Bool() {
				val = "true"
			} else {
				val = "false"
			}
		}
		if x.Kind() != reflect.Struct {
			if o == "optional" {
				if val != "" && val != "0" {
					r[n] = val
				}
			} else {
				if Func2has {
					arg := []reflect.Value{reflect.ValueOf(fs), reflect.ValueOf(n), reflect.ValueOf(val)}
					if ret := Func2.Call(arg); !ret[0].Bool() {
						r[n] = val
					}
				} else {
					r[n] = val
				}
			}
		} else {
			if Func2has {
				arg := []reflect.Value{reflect.ValueOf(fs), reflect.ValueOf(n), reflect.ValueOf(val)}
				if ret := Func2.Call(arg); !ret[0].Bool() {
					r[n] = val
				}
			}
		}
	}
	// 递归返回
	if len(ret) != 0 {
		*ret[0] = r
		return nil
	}
	// 最终返回
	return r
}

// NewRequest 执行请求
func (b *BaseApp) NewRequest(method string, postData interface{}, outData interface{}) error {
	//var ret BaseResp
	var ret BaseResp
	var dat = ParamMap{}
	if postData != nil {
		if values, ok := postData.(ParamMap); ok {
			if len(values) > 0 {
				for k, v := range values {
					dat[k] = fmt.Sprint(v)
				}
			}
		} else {
			dat = ToParamMap(postData)
		}
	}
	params := ParamMap{
		"method":       method,
		"app_key":      b.Key,
		"access_token": *b.accessToken,
		"param_json":   dat,
		"timestamp":    time.Now().Format(unit.TimeYmdHis),
		"v":            "2",
		"sign":         "",
	}
	params["sign"] = Sign(params, b.Secret)

	query := url.Values{}
	for k, v := range params {
		if s, ok := v.(string); ok {
			query.Add(k, s)
		}
	}
	//fmt.Println("####", query.Encode())
	//str, _ := url.QueryUnescape(queryStr)
	b.RequestUrl = b.gatewayURL + "/" + strings.ReplaceAll(method, ".", "/")

	resp, err := http.PostForm(b.RequestUrl, query)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//_ = json.Unmarshal(body, &ret)

	//数字默认是处理为float64类型的，这就导致了int64可能会丢失精度，这时候要将处理的数字转换成json.Number的形式
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	_ = decoder.Decode(&ret)

	if ret.ErrNo != 0 || ret.Message != "success" {
		return fmt.Errorf("response error %d %s", ret.ErrNo, ret.Message)
	}
	if outData == nil {
		return nil
	}
	if ret.Data == nil {
		return errors.New("response error data is nil")
	}
	if reflect.TypeOf(outData).Elem().Kind() == reflect.Interface {
		rd := reflect.ValueOf(ret.Data)
		reflect.ValueOf(outData).Elem().Set(rd)
		return nil
	}
	if reflect.TypeOf(outData).Elem().Field(0).Tag.Get("json") != "" {
		dataByte, _ := json.Marshal(ret.Data)
		return json.Unmarshal(dataByte, outData)
	}
	return mapstructure.Decode(ret.Data, outData)
}

// Sign 参数签名
// 该方法会将param_json转换为json
func Sign(param ParamMap, secret string) string {
	paramJSON := param["param_json"].(ParamMap)
	if len(paramJSON) == 0 {
		param["param_json"] = "{}"
	} else {
		var ks []string
		for k := range paramJSON {
			ks = append(ks, k)
		}
		sort.Strings(ks)
		for i, k := range ks {
			ks[i] = fmt.Sprintf(`"%v":"%v"`, k, paramJSON[k])
		}
		param["param_json"] = "{" + strings.Join(ks, unit.SPE3) + "}"
	}
	signStr := ""
	for _, k := range SortKeyList {
		if len(param[k].(string)) == 0 {
			continue
		}
		signStr += fmt.Sprintf("%v%v", k, param[k])
	}
	signStr = ReplaceSpecial(secret + signStr + secret)
	fmt.Println("====", signStr)
	h := md5.New()
	h.Write([]byte(signStr))
	return hex.EncodeToString(h.Sum(nil))
}

func ReplaceSpecial(param string) string {
	param = strings.ReplaceAll(param, "&", "\\u0026")
	param = strings.ReplaceAll(param, "<", "\\u003c")
	param = strings.ReplaceAll(param, ">", "\\u00ce")
	return param
}

// ExchangeAccessToken 通过 授权码 换取 access_token.
// https://op.jinritemai.com/docs/guide-docs/9/22
func (b *BaseApp) ExchangeAccessToken(code string) (*App, error) {
	app := App{}
	body := url.Values{}
	body.Add("app_id", b.Key)
	body.Add("app_secret", b.Secret)
	body.Add("code", code)
	body.Add("grant_type", "authorization_code")
	resp, err := http.Get(GatewayURL + "/oauth2/access_token?" + body.Encode())
	if err != nil {
		return nil, err
	}
	var ret BaseResp
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(ret.Data, &app); err != nil {
		return nil, err
	}

	if ret.ErrNo != 0 {
		return nil, errors.New(ret.Message)
	}

	b.accessToken = &app.AccessToken
	app.base = b
	app.CreatedAt = time.Now().Unix()
	return &app, nil
}

// RefreshAccessToken 通过 refreshToken 获取新的 access_token.
// https://op.jinritemai.com/docs/guide-docs/9/22
func (b *BaseApp) RefreshAccessToken(refreshToken string) (*App, error) {
	app := App{}
	body := url.Values{}
	body.Add("app_id", b.Key)
	body.Add("app_secret", b.Secret)
	body.Add("refresh_token", refreshToken)
	body.Add("grant_type", "refresh_token")
	resp, err := http.Get(GatewayURL + "/oauth2/refresh_token?" + body.Encode())
	if err != nil {
		return nil, err
	}
	var ret BaseResp
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(ret.Data, &app); err != nil {
		return nil, err
	}

	if ret.ErrNo != 0 {
		return nil, errors.New(ret.Message)
	}

	b.accessToken = &app.AccessToken
	app.base = b
	app.CreatedAt = time.Now().Unix()

	return &app, nil
}

//请求api
func (b *BaseApp) RequestApi(method string, input interface{}, output interface{}) error {
	var ret BaseResp
	data := b.ParseParams(method, input)
	//fmt.Println("api request:", data.Encode())

	b.RequestUrl = b.gatewayURL + "/" + strings.ReplaceAll(method, ".", "/")
	resp, err := http.PostForm(b.RequestUrl, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println("api response:", string(body))

	//数字默认是处理为float64类型的，这就导致了int64可能会丢失精度，这时候要将处理的数字转换成json.Number的形式
	decoder := json.NewDecoder(resp.Body)
	decoder.UseNumber()
	_ = decoder.Decode(&ret)
	if ret.ErrNo != 0 || ret.Message != "success" {
		return fmt.Errorf("response error %d %s", ret.ErrNo, ret.Message)
	}
	if output == nil {
		return nil
	}
	if ret.Data == nil {
		return errors.New("response error data is nil")
	}
	//fmt.Println("##data##", ret.Data)
	//直接使用json反解析data到数据
	dataByte, _ := json.Marshal(ret.Data)
	return json.Unmarshal(dataByte, output)
}

func (b *BaseApp) ParseParams(method string, input interface{}) url.Values {
	var uri = url.Values{}
	uri.Add("app_key", b.Key)
	uri.Add("access_token", *b.accessToken)

	uri.Add("method", method)
	paramByte, _ := json.Marshal(input)
	paramJson := string(paramByte)
	uri.Add("param_json", paramJson)
	timestamp := time.Now().Format(unit.TimeYmdHis)
	uri.Add("timestamp", timestamp)
	v := "2"
	uri.Add("v", v)
	//签名
	sign := b.CreateSign(method, paramJson, timestamp, v)
	uri.Add("sign", sign)

	return uri
}
func (b *BaseApp) CreateSign(method, paramJson, timestamp, v string) string {
	//var str string = secret + "app_key6844048284663924231" + "methodproduct.list" + "param_json{\"page\":\"0\",\"size\":\"20\"}" + "timestamp2020-07-05 22:33:59" + "v2" + secret
	str := fmt.Sprintf("%sapp_key%smethod%sparam_json%stimestamp%sv%s%s", b.Secret, b.Key, method, paramJson, timestamp, v, b.Secret)
	return getMd5String(str)
}
func getMd5String(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		return ""
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}
