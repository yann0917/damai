package damai

import (
	"encoding/json"
	"strings"
)

type (
	// AdmOrderQuery 大麦-查询分销单
	AdmOrderQuery struct {
		param Parameter
	}
	// AdmOrderConfirm 大麦-出票
	AdmOrderConfirm struct {
		param Parameter
	}
	// AdmOrderDistCreate 大麦-新分销下单
	AdmOrderDistCreate struct {
		Param OrderCreateParam `json:"param"`
	}
	// AdmOrderCancel 大麦-库存释放
	AdmOrderCancel struct {
		param Parameter
	}

	// AdmOrderDirectrefund 大麦-直接退票
	AdmOrderDirectrefund struct {
		param Parameter
	}
)

const (
	// MaitixOrderQuery 大麦-查询分销单
	MaitixOrderQuery = "alibaba.damai.maitix.order.query"
	// MaitixOrderConfirm 大麦-出票
	MaitixOrderConfirm = "alibaba.damai.maitix.order.confirm"
	// MaitixOrderCancel 大麦-库存释放
	MaitixOrderCancel = "alibaba.damai.maitix.order.cancel"
	// MaitixOrderDirectrefund 大麦-直接退票
	MaitixOrderDirectrefund = "alibaba.damai.maitix.order.directrefund"
	// MaitixOrderDistCreate 大麦-新分销下单
	MaitixOrderDistCreate = "alibaba.damai.maitix.order.distribution.create"
	// MaitixProjectDistQuery 分销单个项目信息查询
	MaitixProjectDistQuery = "alibaba.damai.maitix.project.distribution.query"
	// MaitixEticketDistQuery 分销电子票查询接口
	MaitixEticketDistQuery = "alibaba.damai.maitix.eticket.distribution.query"
	// MaitixProjectDistQuerybypage 分销项目分页查询项目列表服务
	MaitixProjectDistQuerybypage = "alibaba.damai.maitix.project.distribution.querybypage"
	// MaitixProjectDistDetailQuery 大麦分销项目内容详情查询
	MaitixProjectDistDetailQuery = "alibaba.damai.maitix.project.distribution.detail.query"
	// MaitixGWProjectStatusQuery 分销项目状态查询
	MaitixGWProjectStatusQuery = "alibaba.damai.maitix.opengateway.project.status.query"
	// MaitixGWPerformStatusQuery 分销场次状态查询
	MaitixGWPerformStatusQuery = "alibaba.damai.maitix.opengateway.perform.status.query"
	// MaitixGWTicketItemStatusQuery 分销票品状态查询
	MaitixGWTicketItemStatusQuery = "alibaba.damai.maitix.opengateway.ticketItem.status.query"
	// MaitixSeatInfoQuery 分销商查询座位信息
	MaitixSeatInfoQuery = "alibaba.damai.maitix.seat.info.query"
	// MaitixSetTokenQuery 分销商选座获取qtoken
	MaitixSetTokenQuery = "alibaba.damai.maitix.seat.token.query"
	// MaitixDeliveryCalculate 计算渠道用户下单快递费
	MaitixDeliveryCalculate = "alibaba.damai.maitix.distribution.delivery.calculate"
	// MaitixDeliveryQuery 查询分销物流单
	MaitixDeliveryQuery = "alibaba.damai.maitix.distribution.delivery.query"
	// MaitixExchangePointQuery 分销查询取票点接口
	MaitixExchangePointQuery = "alibaba.damai.maitix.distribution.exchangepoint.query"
	// MaitixCMBParamEncrypt 加密招商一网能支付入参
	MaitixCMBParamEncrypt = "alibaba.damai.maitix.distribution.cmb.paramencrypt"
	// MaitixCMBQueryPay 查询招行支付状态api
	MaitixCMBQueryPay = "alibaba.damai.maitix.distribution.cmb.querypayresult"


)

// MxResult order 响应参数结构
type MxResult struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

// OrderCreateParam 创建订单接收参数
type OrderCreateParam struct {
	ProjectID            int          `json:"project_id"`
	PerformID            int          `json:"perform_id"`
	ThirdOrderID         string       `json:"third_order_no"`
	TotalPrice           int          `json:"total_price"`   // 订单总价
	TicketMode           int          `json:"ticket_mode"`   // 出票方式 1=纸质票 2=身份证电子票 3=二维码电子票 4=短信电子票
	DeliveryType         int          `json:"delivery_type"` // 取票方式 1-无纸化；2-快递票；3-自助换票；4-门店自取。1和3为电子票，2和4为纸质票。
	RealTicketBuyerPhone string       `json:"real_ticket_buyer_phone"`
	SeatProps            []SeatProp   `json:"seat_props"`
	TicketItems          []TicketItem `json:"ticket_items"`
}

// SeatProp 座位参数
type SeatProp struct {
	TicketItemID              int    `json:"ticket_item_id"`
	RealTicketOwnerIDCardNo   string `json:"real_ticket_owner_id_card_no"`
	RealTicketOwnerIDCardType int    `json:"real_ticket_owner_id_card_type"`
	RealTicketOwnerName       string `json:"real_ticket_owner_name"`
	RealTicketOwnerPhone      string `json:"real_ticket_owner_phone"`
}

// TicketItem 下单票品信息
type TicketItem struct {
	IsPackage    int `json:"is_package"`
	Price        int `json:"price"`    // 票价，分
	Quantity     int `json:"quantity"` // 数量， 如果是套票,则是套票的套数
	TicketItemID int `json:"ticket_item_id"`
}

// OrderCreateResp 创建订单响应数据
type OrderCreateResp struct {
	OrderID      string     `json:"order_id"` // 大麦订单号
	TotalAmount  int        `json:"total_amount"`
	ExpressFee   int        `json:"express_fee"`
	SubOrderDtos []SubOrder `json:"sub_order_dtos"`
}

// OrderConfirmResp 确认订单响应数据
type OrderConfirmResp struct {
	OrderID   string `json:"order_id"`   // 订单id
	PayStatus int    `json:"pay_status"` // 支付状态：0:失败,1:成功
}

// OrderCancelResp 取消订单响应数据
type OrderCancelResp struct {
	OrderID int `json:"order_id"` // 订单id
}

// SubOrder 子订单
type SubOrder struct {
	SubOrderID         int          `json:"sub_order_id"` // 大麦子订单号
	ExternalSubOrderNo string       `json:"external_sub_order_no"`
	OriginPrice        int          `json:"origin_price"` // 原价，分
	RealPrice          int          `json:"real_price"`   // 实际价，分
	VoucherID          int          `json:"voucher_id"`   // 票单 ID
	SubOrderSeatDto    SubOrderSeat `json:"sub_order_seat_dto"`
}

// SubOrderSeat 子订单座位
type SubOrderSeat struct {
	ProjectID     int    `json:"project_id"`
	ProjectName   string `json:"project_name"`
	PerformID     int    `json:"perform_id"`
	PerformName   string `json:"perform_name"`
	PriceID       int    `json:"price_id"`
	PriceName     string `json:"price_name"`
	Entry         string `json:"entry"`
	StandID       int    `json:"stand_id"`
	StandName     string `json:"stand_name"`
	SeatFloorID   int    `json:"seat_floor_id"`
	SeatFloorName string `json:"seat_floor_name"`
	SeatAreaID    int    `json:"seat_area_id"`
	SeatAreaName  string `json:"seat_area_name"`
	SeatGroup     int    `json:"seat_group"` // 座位分组，0:无座 1:有座
	CombinID      string `json:"combin_id"`
	SeatID        int    `json:"seat_id"`
	SeatName      string `json:"seat_name"`
	SeatRowID     int    `json:"seat_row_id"`
	SeatRowName   string `json:"seat_row_name"`
	SeatType      byte   `json:"seat_type"`
}

// QueryOrder 查询分销单
type QueryOrder struct {
	OutOID       string          `json:"outOid"`
	GmtCreate    int             `json:"gmtCreate"`
	GmtModified  int             `json:"gmtModified"`
	OrderID      int             `json:"orderId"`
	TicketMode   int             `json:"ticketMode"`
	BatchCode    int             `json:"batchCode"`
	RefundStatus int             `json:"refundStatus"`
	TotalAmount  int             `json:"totalAmount"`
	SellerID     int             `json:"sellerId"`
	TradeStatus  int             `json:"tradeStatus"`
	PayStatus    int             `json:"payStatus"`
	SettleType   int             `json:"settleType"`
	SubOrders    []QuerySubOrder `json:"subOrders"`
}

// QuerySubOrder 查询子订单
type QuerySubOrder struct {
	SubOrderID       int    `json:"subOrderId"`
	HasLocked        bool   `json:"hasLocked"`
	HasValidated     bool   `json:"hasValidated"`
	RefundStatus     int    `json:"refundStatus"`
	SeatType         int    `json:"seatType"`
	Price            int    `json:"price"`
	VoucherID        int    `json:"voucherId"`
	BatchCode        int    `json:"batchCode"`
	TicketType       int    `json:"ticketType"`
	CertificateNo    string `json:"certificateNo"`
	CustomerName     string `json:"customerName"`
	PerformID        int    `json:"performId"`
	HasPrinted       bool   `json:"hasPrinted"`
	PhoneNumber      string `json:"phoneNumber"`
	PhoneCountryCode string `json:"phoneCountryCode"`
	PerformName      string `json:"performName"`
	TradeStatus      int    `json:"tradeStatus"`
	PriceName        string `json:"priceName"`
	PriceID          int    `json:"priceId"`
	ProjectName      string `json:"projectName"`
	ProjectID        int    `json:"projectId"`
	PayStatus        int    `json:"payStatus"`
	CertificateType  int    `json:"certificateType"`
}

// --------------------------------------------//

// APIName api method name
func (adm *AdmOrderConfirm) APIName() string {
	return MaitixOrderConfirm
}

// SetParam set request param
func (adm *AdmOrderConfirm) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmOrderConfirm) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmOrderConfirm) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "order_id") {
		return "Missing required arguments:param.order_id", false
	}

	return "", true
}

// NewAdmOrderConfirm 确认订单
//
// 文档地址：https://open.taobao.com/api.htm?docId=38104&docType=2
func NewAdmOrderConfirm(orderID int) *AdmOrderConfirm {
	param, err := json.Marshal(Parameter{"order_id": orderID})
	if err != nil {
		// log.Err(err)
	}
	return &AdmOrderConfirm{
		param: Parameter{
			"param": string(param),
		},
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmOrderDistCreate) APIName() string {
	return MaitixOrderDistCreate
}

// SetParam set request param
func (adm *AdmOrderDistCreate) SetParam(k string, v interface{}) {
	// adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmOrderDistCreate) GetParam() (p Parameter) {
	// struct2map
	pa := make(Parameter)
	param, _ := json.Marshal(adm.Param)
	json.Unmarshal(param, &pa)

	p = make(Parameter)
	p["param"] = interfaceToString(pa)
	return
}

// CheckParam check request param
func (adm *AdmOrderDistCreate) CheckParam() (msg string, ok bool) {
	if adm.Param.ProjectID == 0 {
		return "Missing required arguments:param.project_id", false
	}
	if adm.Param.PerformID == 0 {
		return "Missing required arguments:param.perform_id", false
	}

	return "", true
}

// NewAdmOrderDistCreate 创建订单
//
// 文档地址：https://open.taobao.com/api.htm?docId=43662&docType=2
func NewAdmOrderDistCreate(p OrderCreateParam) *AdmOrderDistCreate {

	return &AdmOrderDistCreate{
		Param: p,
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmOrderQuery) APIName() string {
	return MaitixOrderQuery
}

// SetParam set request param
func (adm *AdmOrderQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmOrderQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmOrderQuery) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "order_id") {
		return "Missing required arguments:param.order_id", false
	}

	return "", true
}

// NewAdmOrderQuery 大麦-查询分销单
//
// 文档地址：https://open.taobao.com/api.htm?docId=38040&docType=2
func NewAdmOrderQuery(orderID int) *AdmOrderQuery {
	param, err := json.Marshal(Parameter{
		"order_id":                  orderID,
		"exclude_useless_sub_order": false})
	if err != nil {
		// log.Err(err)
	}
	return &AdmOrderQuery{
		param: Parameter{
			"param": string(param),
		},
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmOrderCancel) APIName() string {
	return MaitixOrderCancel
}

// SetParam set request param
func (adm *AdmOrderCancel) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmOrderCancel) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmOrderCancel) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "order_id") {
		return "Missing required arguments:param.order_id", false
	}

	return "", true
}

// NewAdmOrderCancel 取消订单
//
// 文档地址：https://open.taobao.com/api.htm?docId=38142&docType=2
func NewAdmOrderCancel(orderID int) *AdmOrderCancel {
	param, err := json.Marshal(Parameter{"order_id": orderID})
	if err != nil {
		// log.Err(err)
	}
	return &AdmOrderCancel{
		param: Parameter{
			"param": string(param),
		},
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmOrderDirectrefund) APIName() string {
	return MaitixOrderDirectrefund
}

// SetParam set request param
func (adm *AdmOrderDirectrefund) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmOrderDirectrefund) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmOrderDirectrefund) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok &&
		!strings.Contains(value, "order_id") &&
		!strings.Contains(value, "refund_type") {
		return "Missing required arguments:param.order_id", false
	}

	return "", true
}

// NewAdmOrderDirectrefund 大麦-直接退票
//
// 文档地址：https://open.taobao.com/api.htm?docId=43627&docType=2
//
// orderID 大麦订单号
// refundType 退票原因：0=出错票，1=客人退票，2=其它原因
func NewAdmOrderDirectrefund(orderID int, refundType int) *AdmOrderDirectrefund {
	param, err := json.Marshal(Parameter{
		"order_id":    orderID,
		"refund_type": refundType})
	if err != nil {
		// log.Err(err)
	}
	return &AdmOrderDirectrefund{
		param: Parameter{
			"param": string(param),
		},
	}

}
