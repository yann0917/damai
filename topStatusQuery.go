package damai

import (
	"encoding/json"
	"strings"
)

type (

	// AdmGWProjectStatusQuery 分销项目状态查询
	AdmGWProjectStatusQuery struct {
		param Parameter
	}

	// AdmGWPerformStatusQuery 分销场次状态查询
	AdmGWPerformStatusQuery struct {
		param Parameter
	}

	// AdmGWTicketItemStatusQuery 分销票品状态查询
	AdmGWTicketItemStatusQuery struct {
		param Parameter
	}

	// ProjectStatus 项目状态
	ProjectStatus struct {
		ProjectID         int             `json:"project_id"`
		Status            int             `json:"status"` // 1-可售,0-不可售
		PerformStatusList []PerformStatus `json:"dis_perform_status_d_t_o_list,omitempty"`
	}

	// PerformStatus 场次状态
	PerformStatus struct {
		ProjectID            int                `json:"project_id"`
		PerformID            int                `json:"perform_id"`
		Status               int                `json:"status"`
		TicketItemStatusList []TicketItemStatus `json:"dis_ticket_item_status_d_t_o_list,omitempty"`
	}

	// TicketItemStatus 票品状态
	TicketItemStatus struct {
		ProjectID    int `json:"project_id"`
		PerformID    int `json:"perform_id"`
		TicketItemID int `json:"ticket_item_id"`
		Status       int `json:"status"`
		SubStatus    int `json:"sub_status"`
	}
)

// --------------------------------------------//

// APIName api method name
func (adm *AdmGWProjectStatusQuery) APIName() string {
	return MaitixGWProjectStatusQuery
}

// SetParam set request param
func (adm *AdmGWProjectStatusQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmGWProjectStatusQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmGWProjectStatusQuery) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["dis_project_status_query_param"]; !ok {
		return "Missing required arguments:dis_project_status_query_param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "project_id") {
		return "Missing required arguments:param.project_id", false
	}

	return "", true
}

// NewAdmGWProjectStatusQuery 分销项目状态查询
//
// 文档地址：https://open.taobao.com/api.htm?docId=46024&docType=2
func NewAdmGWProjectStatusQuery(projectID int, queryPerformStatus bool) *AdmGWProjectStatusQuery {
	param, err := json.Marshal(Parameter{
		"project_id":           projectID,
		"query_perform_status": queryPerformStatus})
	if err != nil {
		// log.Err(err)
	}
	return &AdmGWProjectStatusQuery{
		param: Parameter{
			"dis_project_status_query_param": string(param),
		},
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmGWPerformStatusQuery) APIName() string {
	return MaitixGWPerformStatusQuery
}

// SetParam set request param
func (adm *AdmGWPerformStatusQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmGWPerformStatusQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmGWPerformStatusQuery) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["dis_perform_status_query_param"]; !ok {
		return "Missing required arguments:dis_perform_status_query_param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "perform_id") {
		return "Missing required arguments:param.perform_id", false
	}

	return "", true
}

// NewAdmGWPerformStatusQuery 分销场次状态查询
//
// 文档地址：https://open.taobao.com/api.htm?docId=46026&docType=2
func NewAdmGWPerformStatusQuery(performID int, queryTicketItemStatus bool) *AdmGWPerformStatusQuery {
	param, err := json.Marshal(Parameter{
		"perform_id":               performID,
		"query_ticket_item_status": queryTicketItemStatus})
	if err != nil {
		// log.Err(err)
	}
	return &AdmGWPerformStatusQuery{
		param: Parameter{
			"dis_perform_status_query_param": string(param),
		},
	}

}

// --------------------------------------------//

// APIName api method name
func (adm *AdmGWTicketItemStatusQuery) APIName() string {
	return MaitixGWTicketItemStatusQuery
}

// SetParam set request param
func (adm *AdmGWTicketItemStatusQuery) SetParam(k string, v interface{}) {
	adm.param[k] = interfaceToString(v)
}

// GetParam get request param
func (adm *AdmGWTicketItemStatusQuery) GetParam() Parameter {
	return adm.param
}

// CheckParam check request param
func (adm *AdmGWTicketItemStatusQuery) CheckParam() (msg string, ok bool) {
	var param interface{}
	if param, ok = adm.param["dis_ticket_item_status_query_param"]; !ok {
		return "Missing required arguments:dis_ticket_item_status_query_param", false
	}
	if value, ok := param.(string); !ok || !strings.Contains(value, "ticket_item_id_list") {
		return "Missing required arguments:param.ticket_item_id_list", false
	}

	return "", true
}

// NewAdmGWTicketItemStatusQuery 分销票品状态查询
//
// 文档地址：https://open.taobao.com/api.htm?docId=46025&docType=2
func NewAdmGWTicketItemStatusQuery(ticketID []int) *AdmGWTicketItemStatusQuery {
	param, err := json.Marshal(Parameter{"ticket_item_id_list": ticketID})
	if err != nil {
		// log.Err(err)
	}
	return &AdmGWTicketItemStatusQuery{
		param: Parameter{
			"dis_ticket_item_status_query_param": string(param),
		},
	}

}
