package damai

import (
	"strings"
)

type (
	// AdmProjectDistQuerybypage 分销项目分页查询项目列表
	AdmProjectDistQuerybypage struct {
		param Parameter
	}
	// AdmProjectDistQuery 分销单个项目信息查询
	AdmProjectDistQuery struct {
		param Parameter
	}
	// AdmProjectDistDetailQuery 大麦分销项目内容详情查询
	AdmProjectDistDetailQuery struct {
		param Parameter
	}
	// AdmEticketDistQuery 分销电子票查询接口
	AdmEticketDistQuery struct {
		param Parameter
	}
)

// ProjectDistDetail 大麦分销项目内容详情 model
type ProjectDistDetail struct {
	ProjectID              int    `json:"project_id"`
	ProjectName            string `json:"project_name"`
	ClassifyCode           int    `json:"classify_code"` // 项目分类编码
	ClassifyName           string `json:"classify_name"` // 项目分类名称
	SubClassifyCode        int    `json:"sub_classify_code"`
	SubClassifyName        string `json:"sub_classify_name"`
	ShowPic                string `json:"show_pic"`
	ShowDetail             string `json:"show_detail"`     // 演出详情
	ShowStartTime          string `json:"show_start_time"` // 演出开售时间
	ShowEndTime            string `json:"show_end_time"`   // 演出销售结束时间
	LimitNotice            string `json:"limit_notice"`    // 限购说明
	ChoiceSeatNotice       string `json:"choice_seat_notice"`
	RealNameNotice         string `json:"real_name_notice"`
	ChildrenNotice         string `json:"children_notice"`
	PolicyOfReturn         string `json:"policy_of_return"`
	EntranceNotice         string `json:"entrance_notice"`
	EticketNotice          string `json:"eticket_notice"`
	SelfGetTicketNotice    string `json:"self_get_ticket_notice"`
	DespositInfo           string `json:"deposit_info"`
	ProhibitedItems        string `json:"prohibited_items"`
	Artists                string `json:"artists"`
	IPCard                 string `json:"ip_card"`
	ShowTime               string `json:"show_time"`
	DeliveryTypes          string `json:"delivery_types"`      // 取票类型1-无纸化,2-快递票,3-自助换票,4-门店自取。
	PostCity               string `json:"post_city"`           // 发货城市国标
	PickupAddressList      string `json:"pickup_address_list"` // 上门自取地址
	PerformTimeDetailList  string `json:"perform_time_detail_list"`
	PurchaseLimitationOnce int    `json:"purchase_limitation_once"` // 项目单次限购数量

}

// ProjectDistDetailResp 大麦分销项目内容详情
type ProjectDistDetailResp struct {
	OpenResult
	Model ProjectDistDetail `json:"model"`
}

// OpenResult project 响应返回参数结构
type OpenResult struct {
	Success   bool   `json:"success"`
	ErrorCode string `json:"error_code"`
	ErrorMsg  string `json:"meror_msg"`
}

// ProjectDistInfo 单个项目信息
type ProjectDistInfo struct {
	ProjectID       int           `json:"project_id"`
	ProjectName     string        `json:"project_name"`
	ProjectType     int           `json:"project_type"`   // 项目类型 0:普通项目 1:预售项目 2:先付先抢-先付项目 3:先付先抢-先抢项目 4:搭售项目 5:超级票-暂时没用
	ProjectStatus   int           `json:"project_status"` // 项目状态，0：创建中 10：已创建 20：待销售 30：销售中 40：项目取消 50：项目结束 60 定时开售
	Introduce       string        `json:"introduce"`
	IsHasSeat       int           `json:"is_has_seat"`       // 是否有座：0=无座 1=有座
	IsTest          bool          `json:"is_test"`           // 是否测试项目 0-正式项目 1-测试项目
	IsGeneralAgent  int           `json:"is_general_agent"`  // 是否总票代
	TicketAgencyFee string        `json:"ticket_agency_fee"` // 票务代理费（单位：百分比）
	TraderIDList    []int         `json:"trader_id_list"`
	TraderNameList  []string      `json:"trader_name_list"`
	City            city          `json:"city"`
	Province        city          `json:"province"`
	Country         city          `json:"country"`
	Venue           venue         `json:"venue"`
	PerformInfoList []PerformInfo `json:"perform_info_list"`
}

// PerformInfo 场次信息
type PerformInfo struct {
	PerformID           int            `json:"perform_id"`
	PerformName         string         `json:"perform_name"`
	PerformType         int            `json:"perform_type"`   // 场次类型(1 单场次，2 多次通票，3 单次通票)-暂时没用,可以认为都是普通场次
	PerformStatus       int            `json:"perform_status"` // 场次状态(0：创建中 10：已创建 20：待销售 30：销售中 40：场次取消 50：场次结束)一般不会透出30之前的状态给渠道
	ChangeReason        string         `json:"change_reason"`
	StartTime           string         `json:"start_time"`
	EndTime             string         `json:"end_time"`
	IsChangePerformTime int            `json:"is_change_perform_time"`
	Remark              string         `json:"remark"`
	ReserveSeat         int            `json:"reserve_seat"` // 是否对号入座 0：不对号入座 1：对号入座 2：对区入座
	SubPerformList      string         `json:"-"`            // 子场次列表-暂时没用
	WeeklyList          []int          `json:"weekly_list"`  // 场次有效期规则 每周周几可以
	PerformSetting      performSetting `json:"perform_setting"`
	PriceInfoList       []priceInfo    `json:"price_info_list"`
}

// performSetting 场次设置
type performSetting struct {
	PerformID            int    `json:"perform_id"`
	SaleType             int    `json:"sale_type"` // 销售设置 0开票 1预售
	TakeTicketTypes      []int  `json:"take_ticket_types"`
	SeatSelectTypeList   []int  `json:"seat_select_type_list"`   // 选座类型0-立即购买 1-选座购买,当是有座项目并且是列表有1.就可以h5选座
	IssueTicketModesList []int  `json:"issue_ticket_modes_list"` // 出票方式 1纸质票 2静态二维码电子票 3动态二维码电子票 4身份证电子票 5 短信码电子票
	IssueEnterModesList  []int  `json:"issue_enter_modes_list"`  // 入场方式 1纸质票入场 2电子票入场
	IsOneOrderOneCard    int    `json:"is_one_order_one_card"`   // 一单一证 0：不是，1：是
	IsOneTicketOneCard   int    `json:"is_one_ticket_one_card"`  // 一票一证 0：不是，1：是
	IsRealNameEnter      int    `json:"is_real_name_enter"`      // 是否实名制入场 0：不是，1：是
	CardType             string `json:"card_type"`               // json 字符串
}

// priceInfo 价格
type priceInfo struct {
	PerformID int    `json:"perform_id"`
	MaxStock  int    `json:"max_stock"`
	PriceID   int    `json:"price_id"`
	Price     price  `json:"price"`
	PriceName string `json:"price_name"`
	PriceType int    `json:"price_type"` // 票品的类型 0普通单票 1固定套票 2 自由组合套票
	ProjectID int    `json:"project_id"`
	AbleSell  bool   `json:"able_sell"` // 是否可售 true可售 false不可售
}

type price struct {
	Cent int `json:"cent"` // 分
}
type city struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Ticket 电子票 model
type Ticket struct {
	BatchCode            int    `json:"batch_code"` // 票池批次号
	CertificateNo        string `json:"certificate_no"`
	CertificateType      string `json:"certificate_type"`
	CombineID            int    `json:"combine_id"`
	CustomerName         string `json:"customer_name"`
	EntryType            string `json:"entry_type"`             // 入场方式 1-电子票 2-纸质票 3-电子票+纸质票
	ExchangeCode         string `json:"exchange_code"`          // 换票码
	ExchangeTicketMethod string `json:"exchange_ticket_method"` // 换票方式：1-电子票；2-纸质票
	FloorID              int    `json:"floor_id"`
	FloorName            string `json:"floor_name"`
	FullSeatName         string `json:"full_seat_name"`
	MainOrderID          int    `json:"main_order_id"` // 主订单号
	OrderID              int    `json:"order_id"`      // 子订单号
	PerformID            int    `json:"perform_id"`
	PerformTime          string `json:"perform_time"`
	PhoneCountryCode     string `json:"phone_country_code"`
	PhoneNumber          string `json:"phone_number"`
	PrintStatus          string `json:"print_status"`   // 打印状态 ：1-未打印 2-已打印
	ProductSource        int    `json:"product_source"` // 项目来源: 1票务云；2直连第三方
	ProjectID            int    `json:"project_id"`
	ProjectName          string `json:"project_name"`

	QrCode      string `json:"qr_code"`
	QrCodeType  string `json:"qr_code_type"` // 类型 1-静态二维码 2-动态二维码
	ReserveSeat string `json:"reserve_seat"` // 入座方式0=不对号入座 1=对号入座 2=对区入座

	SeatCol string `json:"seat_col"` // 座位号
	SeatID  int    `json:"seat_id"`  // 座位 ID
	SeatRow string `json:"seat_row"` // 座位排

	ServicePhone string `json:"service_phone"` // 服务电话
	SourceSystem int    `json:"source_system"` // 来源系统1-B端，2-C端 默认C端

	SourceTicketItemPrice int    `json:"source_ticket_item_price"` // 原价格
	StandName             string `json:"stand_name"`               // 看台名称
	StandPortal           string `json:"stand_portal"`             // 看台入场口
	SupplierID            int    `json:"supplier_id"`              // 项目代理商 ID

	TicketItemName  string `json:"ticket_item_name"`  // 票品名称
	TicketItemPrice int    `json:"ticket_item_price"` // 普通票品价格
	TicketItemID    int    `json:"ticket_itemid"`     // 票品 ID
	ValidateStatus  string `json:"validate_status"`   // 验票状态 1:未验; 2:已验

	VenueID   int    `json:"venue_id"`   // 场馆 ID
	VenueName string `json:"venue_name"` // 场馆名称
	VoucherID int    `json:"voucher_id"` // 票单号
}

// venue 场馆
type venue struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Lng       string `json:"lng"`
	Lat       string `json:"lat"`
	VenueAddr string `json:"venue_address"`
}

// --------------------------------------------//

// APIName api method name
func (adm *AdmProjectDistQuerybypage) APIName() string {
	return MaitixProjectDistQuerybypage
}

// SetParam set request param
func (adm *AdmProjectDistQuerybypage) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmProjectDistQuerybypage) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmProjectDistQuerybypage) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok && !strings.Contains(value, "page") {
		return "Missing required arguments:param.page_no && page size", false
	}

	return "", true
}

// NewAdmProjectDistQuerybypage  分销项目分页查询项目列表
//
// 文档地址：https://open.taobao.com/api.htm?docId=44335&docType=2
func NewAdmProjectDistQuerybypage() *AdmProjectDistQuerybypage {
	return &AdmProjectDistQuerybypage{
		param: Parameter{
			"param": Parameter{
				"page_no":   1,
				"page_size": 200, // 每页数据大小，可以稍微大一点
			},
		},
	}
}

// --------------------------------------------//

// APIName api method name
func (adm *AdmProjectDistQuery) APIName() string {
	return MaitixProjectDistQuery
}

// SetParam set request param
func (adm *AdmProjectDistQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmProjectDistQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmProjectDistQuery) CheckParam() (msg string, ok bool) {
	if _, ok = adm.param["project_id"]; !ok {
		return "Missing required arguments:project_id", false
	}

	return "", true
}

// NewAdmProjectDistQuery 分销单个项目信息查询
//
// 文档地址：https://open.taobao.com/api.htm?docId=43682&docType=2
func NewAdmProjectDistQuery(projectID int) *AdmProjectDistQuery {
	return &AdmProjectDistQuery{
		param: Parameter{
			"project_id": projectID,
		},
	}
}

// --------------------------------------------//

// APIName api method name
func (adm *AdmProjectDistDetailQuery) APIName() string {
	return MaitixProjectDistDetailQuery
}

// SetParam set request param
func (adm *AdmProjectDistDetailQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmProjectDistDetailQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmProjectDistDetailQuery) CheckParam() (msg string, ok bool) {
	if _, ok = adm.param["project_id"]; !ok {
		return "Missing required arguments:project_id", false
	}

	return "", true
}

// NewAdmProjectDistDetailQuery 销项目内容详情查询
//
// 文档地址：https://open.taobao.com/api.htm?docId=45916&docType=2
func NewAdmProjectDistDetailQuery(projectID int) *AdmProjectDistDetailQuery {
	return &AdmProjectDistDetailQuery{
		param: Parameter{
			"project_id": projectID,
		},
	}
}

// --------------------------------------------//

// APIName api method name
func (adm *AdmEticketDistQuery) APIName() string {
	return MaitixEticketDistQuery
}

// SetParam set request param
func (adm *AdmEticketDistQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmEticketDistQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmEticketDistQuery) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["param"]; !ok {
		return "Missing required arguments:param", false
	}
	if value, ok := param.(string); !ok && !strings.Contains(value, "main_order_id") {
		return "Missing required arguments:param.main_order_id", false
	}

	return "", true
}

// NewAdmEticketDistQuery 分销电子票查询接口
//
// 文档地址：https://open.taobao.com/api.htm?docId=44328&docType=2
func NewAdmEticketDistQuery() *AdmEticketDistQuery {
	return &AdmEticketDistQuery{
		param: Parameter{},
	}
}
