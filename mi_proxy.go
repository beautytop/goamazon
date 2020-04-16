package goamazon

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hunterhug/marmot/miner"
)

type mvpResult struct {
	Code          interface{}       `json:"code"`
	CodeMsg       string            `json:"code_msg"`
	TodayFetchNum interface{}       `json:"today_fetch_num"`
	TodayTotalNum interface{}       `json:"today_total_num"`
	CostTime      string            `json:"cost_time"`
	ResultCount   int64             `json:"result_count"`
	ResponseTime  string            `json:"dtime"`
	Result        []mvpResultDetail `json:"result"`
}

type mvpResultDetail struct {
	Ip           string `json:"ip:port"`   //"ip:port": "67.207.95.138:8080",
	Type         string `json:"http_type"` //"http_type": "HTTPS",
	An           string `json:"anonymous"` //"anonymous": "高匿",
	Isp          string `json:"isp"`       //"isp": "null",
	Country      string `json:"country"`   //"country": "美国"
	TransferTime int64  `json:"transfer_time"`
	PingTime     int64  `json:"ping_time"`
}

func MiProxy(account string, num int) (ips []string, err error) {
	url := "http://proxy.mimvp.com/api/fetch.php?orderid=%s&num=%d&result_format=json&anonymous=5&result_fields=1,2,3,4,5&http_type=1,2,5&ping_time=5&transfer_time=5"
	worker := miner.NewAPI()
	worker.Url = fmt.Sprintf(url, account, num)
	data, err := worker.Get()
	if err != nil {
		return
	}
	r := new(mvpResult)
	err = json.Unmarshal(data, r)
	if err != nil {
		return
	}
	if fmt.Sprintf("%v", r.Code) != "0" {
		return nil, errors.New(r.CodeMsg)
	}
	for _, v := range r.Result {
		if v.Type == "Socks5" {
			v.Ip = "socks5://" + v.Ip
		} else if v.Type == "HTTPS" {
			v.Ip = "https://" + v.Ip
		} else {
			v.Ip = "http://" + v.Ip
		}
		ips = append(ips, v.Ip)
	}
	return
}
