package damai

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// AdmProjectDistDetail 销项目内容详情查询
func AdmProjectDistDetail(projectID int) (detail ProjectDistDetail, err error) {

	api := NewAdmProjectDistDetailQuery(projectID)
	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		log.Println()
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	err = json.Unmarshal([]byte(model), &detail)

	return

}

// AdmProjectDistInfo 分销单个项目信息查询
func AdmProjectDistInfo(projectID int) (info []ProjectDistInfo, err error) {

	api := NewAdmProjectDistQuery(projectID)
	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model_list").Raw

	err = json.Unmarshal([]byte(model), &info)

	return
}

// AdmPerformInfoByDate 根据日期获取场次信息
func AdmPerformInfoByDate(projectID int, date string) (info PerformInfo, err error) {
	projects, err := AdmProjectDistInfo(projectID)
	if err != nil {
		return
	}
	var projectInfo ProjectDistInfo
	if len(projects) == 1 {
		projectInfo = projects[0]
	}

	performInfoList := projectInfo.PerformInfoList

	for _, list := range performInfoList {
		if strings.Contains(list.StartTime, date) {
			info = list
			return
		}

	}
	return
}

// OrderCreate 创建订单
func OrderCreate(param OrderCreateParam) (detail OrderCreateResp, err error) {

	api := NewAdmOrderDistCreate(param)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 订单创建失败
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	err = json.Unmarshal([]byte(model), &detail)
	return
}

// OrderConfirm 确认出票
func OrderConfirm(orderID int) (detail OrderConfirmResp, err error) {
	api := NewAdmOrderConfirm(orderID)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 确认出票失败，锁座有效期内幂等，重复查询三次以上
		log.Println(string(resp) + "order_id:" + strconv.Itoa(orderID))
		if gjson.GetBytes(resp, "result.code").Int() == 8000208 {
			err = errors.New("ERR_DAMAI_ORDER_STATUS_FAILED")
		}
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &detail)

	return
}

// OrderCancel 取消订单
func OrderCancel(orderID int) (detail OrderCancelResp, err error) {
	api := NewAdmOrderCancel(orderID)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 订单取消失败，锁座有效期内幂等，重复查询三次以上
		log.Println(string(resp) + "order_id:" + strconv.Itoa(orderID))
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &detail)

	return
}

// OrderDetail 查询订单
func OrderDetail(orderID int) (detail QueryOrder, err error) {
	api := NewAdmOrderQuery(orderID)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 订单查询失败
		log.Println(string(resp) + "order_id:" + strconv.Itoa(orderID))
		if gjson.GetBytes(resp, "result.code").Int() == -700008 {
			err = errors.New("ERR_DAMAI_ORDER_DETAIL_QUERY_FAILED")
		}
		return
	}

	model1 := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model1)

	var model2 string
	err = json.Unmarshal([]byte(model1), &model2)
	err = json.Unmarshal([]byte(model2), &detail)

	return
}

// OrderDirectRefund 直接退票
//
// orderID 大麦订单号
// refundType 退票原因：0=出错票，1=客人退票，2=其它原因
func OrderDirectRefund(orderID, refundType int) (detail bool, err error) {
	api := NewAdmOrderDirectrefund(orderID, refundType)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 退票失败
		log.Println(string(resp) + "order_id:" + strconv.Itoa(orderID))
		if gjson.GetBytes(resp, "result.code").Int() == 8000101 {
			err = errors.New("ERR_DAMAI_TICKET_REFUND_FAILED")
		}
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &detail)

	return
}

// GetEticketInfo 分销电子票查询
//
// orderID 大麦订单号
func GetEticketInfo(orderID int) (detail []Ticket, err error) {
	api := NewAdmEticketDistQuery()

	api.SetParam("param", Parameter{
		"main_order_id": orderID,
	})
	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()
	fmt.Println(success)
	if !success {
		// TODO: 退票失败
		log.Println(string(resp) + "order_id:" + strconv.Itoa(orderID))
		// TODO: 异常码待补充
		if gjson.GetBytes(resp, "result.code").Int() == 8000101 {
			err = errors.New("ERR_DAMAI_TICKET_REFUND_FAILED")
		}
		return
	}

	model := gjson.GetBytes(resp, "result.model_arr_list").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &detail)

	return
}

// GetProjectStatus 查询分销项目状态
func GetProjectStatus(projectID int, performStatus bool) (status ProjectStatus, err error) {
	api := NewAdmGWProjectStatusQuery(projectID, performStatus)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 项目状态失败
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &status)
	return
}

// GetPerformStatus 查询分销场次状态
func GetPerformStatus(performID int, ticketStatus bool) (status PerformStatus, err error) {
	api := NewAdmGWPerformStatusQuery(performID, ticketStatus)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()

	if !success {
		// TODO: 场次状态失败
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model").Raw

	fmt.Println(model)

	err = json.Unmarshal([]byte(model), &status)
	return
}

// GetTicketItemStatus 查询分销票品状态
func GetTicketItemStatus(ticketID []int) (status []TicketItemStatus, err error) {
	api := NewAdmGWTicketItemStatusQuery(ticketID)

	resp, err := TopClient.Exec(api)

	if err != nil {
		return
	}
	success := gjson.GetBytes(resp, "result.success").Bool()
	if !success {
		// TODO: 场次状态失败
		log.Println(string(resp))
		return
	}

	model := gjson.GetBytes(resp, "result.model_list").Raw

	// fmt.Println(model)

	err = json.Unmarshal([]byte(model), &status)
	return
}
