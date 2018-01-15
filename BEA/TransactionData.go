package BEA

import (
	"encoding/json"
	"fmt"
)

type TransactionData struct {
	TransId           string            `json:"trans_id"`             //流水号--------11
	AcquireTransID    string            `json:"acq_trans_id"`         //收单行交易号---37
	TransDate         string            `json:"trans_date"`           //交易日期------13
	TransTime         string            `json:"trans_time"`           //交易时间------12
	Amount            string            `json:"amount"`               //授权金额------04
	Pin               string            `json:"pin"`                  //联机PINBLOCK--52
	MerchantId        string            `json:"merchant_id"`          //商户号--------42
	TerminalId        string            `json:"terminal_id"`          //终端号--------41
	Pan               string            `json:"pan"`                  //主账号--------02
	PanSeqNo          string            `json:"pan_seq_no,omitempty"` //卡片序列号-----
	CardExpireDate    string            `json:"card_exp_date"`        //有效期--------14
	Track1            string            `json:"track1,omitempty"`     //磁道一--------
	Track2            string            `json:"track2"`               //磁道二--------35
	PosEntryMode      string            `json:"pos_entry_mode"`       //刷卡方式------22
	IccRelatedData    map[string]string //IC卡相关数据--
	AuthCode          string            `json:"auth_code"`                   //授权码-------38
	ResponseCode      string            `json:"response_code"`               //响应码-------39
	Invoice           string            `json:"invoice,omitempty"`           //发票号-------62
	OriginalAmount    string            `json:"origin_amount,omitempty"`     //原交易金额----
	OriginalTransType string            `json:"origin_trans_type,omitempty"` //原交易类型----
	TransType         string            `json:"trans_type,omitempty"`        //交易类型------
	CurrencyCode      string            `json:"curreny,omitempty"`           //货币代码------
}

func (t TransactionData) FormJson() string {
	data, err := json.Marshal(t)
	if err != nil {
		return fmt.Sprintf("json marshal error: %s", err.Error())
	} else {
		return string(data)
	}
}

type ReconciliationTotals struct {
	CapturedSalesCount  string
	CapturedSalesAmount string
	RefundCount         string
	RefundAmount        string
}

//后台配置参数
type Config struct {
	Host string //后台地址
	TPDU string //tpdu
	EDS  string //eds
}