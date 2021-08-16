# damai

> 大麦票务云分销API go 语言实现

## 使用方法

使用前请查看大麦票务云分销API文档 [地址](https://open.taobao.com/api.htm?docId=38040&docType=2)

参考以下方法传入大麦 `key` 和  `secret`，之后参考示例 `topApi.go`

```golang
var TopClient *Client

func init() {
    key := "xxxxxxxx"
    secret := "799eb5c0a6f7003bxxxxxxxxxxxxxxxx"
    TopClient = NewClient(key, secret)
}

```

参考 

### 已实现的接口

* [x] 大麦-查询分销单
* [x] 大麦-出票
* [x] 大麦-库存释放
* [x] 大麦-新分销下单
* [x] 分销单个项目信息查询
* [x] 分销电子票查询接口
* [x] 分销项目分页查询项目列表服务
* [x] 大麦分销项目内容详情查询
* [x] 分销项目状态查询
* [x] 分销票品状态查询
* [x] 分销场次状态查询
* [ ] 分销商查询座位信息
* [ ] 分销商选座获取qtoken
* [ ] 计算渠道用户下单快递费
* [ ] 查询分销物流单
* [ ] 分销查询取票点接口
* [ ] 加密招商一网能支付入参
* [ ] 查询招行支付状态api


---

