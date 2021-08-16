package damai

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var projectID = 205092004

func TestMain(m *testing.M) {
	key := "27559889"
	secret := "799eb5c0a6f7003b12953d5e904461b3"
	TopClient = NewClient(key, secret)
	code := m.Run()
	os.Exit(code)
}

func TestAdmOrderCreate(t *testing.T) {
	var param OrderCreateParam
	api := NewAdmOrderDistCreate(param)
	resp, err := TopClient.Exec(api)
	// error {"result":{"code":8000101,"message":"参数错误","success":false},"request_id":"16gnccy12fcri"}
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmOrderConfirm(t *testing.T) {
	api := NewAdmOrderConfirm(10000098)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmOrderQuery(t *testing.T) {
	api := NewAdmOrderQuery(10000098)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}
func TestAdmOrderCancel(t *testing.T) {
	api := NewAdmOrderCancel(10000098)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmDirectrefund(t *testing.T) {
	api := NewAdmOrderDirectrefund(10000098, 2)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmProjectDistQuery(t *testing.T) {
	api := NewAdmProjectDistQuery(projectID)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}
func TestAdmProjectDistQuerybypage(t *testing.T) {
	api := NewAdmProjectDistQuerybypage()
	api.SetParam("param", Parameter{"page_no": 1, "page_size": 100})
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmEticketDistQuery(t *testing.T) {
	api := NewAdmEticketDistQuery()
	// sub_order_id 可不传
	api.SetParam("param", Parameter{"main_order_id": 2621347650669})
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestAdmProjectDistDetailQuery(t *testing.T) {
	api := NewAdmProjectDistDetailQuery(projectID)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestProjectStatusQuery(t *testing.T) {
	api := NewAdmGWProjectStatusQuery(205162033, false)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}

func TestPerformStatusQuery(t *testing.T) {
	api := NewAdmGWPerformStatusQuery(210469909, false)
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}
func TestTicketItemStatusQuery(t *testing.T) {
	api := NewAdmGWTicketItemStatusQuery([]int{1})
	resp, err := TopClient.Exec(api)
	fmt.Println(">>>>>>", string(resp), err)
}
func TestTime(t *testing.T) {
	td, _ := time.ParseDuration("8h")
	timestamp := time.Now().UTC().Add(td).Format("2006-01-02 15:04:05")
	t.Log(timestamp)
}
