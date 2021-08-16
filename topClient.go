package damai

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	MaxIdleConns        = 100
	MaxIdleConnsPerHost = 100
	IdleConnTimeout     = 90
	Timeout             = 30
)

// Client taobao open platform client
type Client struct {
	AppKey    string
	AppSecret string
	Param     Parameter
	IsHTTPS   bool
	IsSanBox  bool
}

// API 接口
type API interface {
	APIName() string
	GetParam() Parameter
	SetParam(k string, v interface{})
	CheckParam() (msg string, ok bool)
}

// Parameter 参数
type Parameter map[string]interface{}

// ErrorResponse 异常响应
type ErrorResponse struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
}

const (
	// OpenRouterHTTPAPI 正式版API路由地址 http 版本
	OpenRouterHTTPAPI = "http://gw.api.taobao.com/router/rest"
	// OpenRouterHTTPSAPI 正式版API路由地址 https 版本
	OpenRouterHTTPSAPI = "https://eco.taobao.com/router/rest"
	// OpenSanBoxRouterHTTPAPI 沙箱环境路由地址 http 版本
	OpenSanBoxRouterHTTPAPI = "http://gw.api.tbsandbox.com/router/rest"
	// OpenSanBoxRouterHTTPSAPI 沙箱环境路由地址 https 版本
	OpenSanBoxRouterHTTPSAPI = "https://gw.api.tbsandbox.com/router/rest"
)

var TopClient *Client


// NewClient new topClient
func NewClient(appKey, appSecret string) *Client {
	return &Client{
		AppKey:    appKey,
		AppSecret: appSecret,
		IsSanBox:  false,
		IsHTTPS:   true,
	}
}

// Exec execute
func (c *Client) Exec(api API) ([]byte, error) {
	if msg, ok := api.CheckParam(); !ok {
		return []byte(msg), nil
	}
	td, _ := time.ParseDuration("8h")
	timestamp := time.Now().UTC().Add(td).Format("2006-01-02 15:04:05")
	param := Parameter{
		"app_key":     c.AppKey,
		"sign_method": "md5",
		"format":      "json",
		"v":           "2.0",
		"timestamp":   timestamp,
		"simplify":    true,
	}

	param["method"] = api.APIName()
	for key, val := range api.GetParam() {
		param[key] = val
	}
	param["sign"] = sign(param, c.AppSecret)
	// fmt.Println(param)

	resp, err := request(c.getRouterAPI(), param)

	return resp, err
}

// getRouterAPI 选择 API router
func (c *Client) getRouterAPI() (router string) {
	if c.IsSanBox {
		if c.IsHTTPS {
			router = OpenSanBoxRouterHTTPSAPI
		} else {
			router = OpenSanBoxRouterHTTPAPI
		}
	} else {
		if c.IsHTTPS {
			router = OpenRouterHTTPSAPI
		} else {
			router = OpenRouterHTTPAPI
		}
	}
	return router
}

// sign 签名算法
//
// 文档地址： https://open.taobao.com/doc.htm?docId=101617&docType=1
func sign(param Parameter, appSecret string) string {
	var keys []string
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var buff bytes.Buffer
	buff.WriteString(appSecret)
	for _, key := range keys {
		buff.WriteString(fmt.Sprintf("%s%v", key, param[key]))
	}
	buff.WriteString(appSecret)

	s := md5.New()
	s.Write(buff.Bytes())

	return strings.ToUpper(hex.EncodeToString(s.Sum(nil)))
}

// request send request as POST method
func request(url string, param Parameter) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   Timeout * time.Second,
				KeepAlive: Timeout * time.Second,
			}).DialContext,
			MaxIdleConns:        MaxIdleConns,
			MaxIdleConnsPerHost: MaxIdleConnsPerHost,
			IdleConnTimeout:     IdleConnTimeout * time.Second,
		},
	}

	contentType := "application/x-www-form-urlencoded;charset=utf-8"
	req, err := http.NewRequest("POST", url, bytes.NewReader([]byte(param.getRequestData())))
	req.Header.Set("Content-Type", contentType)
	resp, err := client.Do(req)

	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("大麦接口响应解析失败" + err.Error() + param.getRequestData())
	}
	errorResponse := gjson.GetBytes(data, "error_response").Raw
	if errorResponse != "" {
		log.Println(errorResponse + param.getRequestData())
		var errResp ErrorResponse
		json.Unmarshal([]byte(errorResponse), &errResp)
		return []byte{}, errors.New(errResp.SubMsg)
	}
	return data, nil
}

// getRequestData get request data
func (p *Parameter) getRequestData() string {
	args := url.Values{}
	for k, v := range *p {
		args.Set(k, fmt.Sprintf("%v", v))
	}
	// fmt.Println(args.Encode())
	return args.Encode()
}

// interfaceToString interface convert to string
func interfaceToString(src interface{}) string {
	if src == nil {
		panic("")
	}
	switch src.(type) {
	case string:
		return src.(string)
	case int, int8, int32, int64:
	case uint8, uint16, uint32, uint64:
	case float32, float64:
		return fmt.Sprintf("%v", src)
	}
	data, err := json.Marshal(src)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// PhoneMasking 手机号脱敏
func PhoneMasking(phone string) string {
	if phone == "" {
		return ""
	}

	return phone[:3] + "0000" + phone[len(phone)-4:]
}
