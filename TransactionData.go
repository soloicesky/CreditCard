package BEA

type TransactionData struct {
	TransId           string            //流水号--------11
	AcquireTransID    string            //收单行交易号---37
	TransDate         string            //交易日期------13
	TransTime         string            //交易时间------12
	Amount            string            //授权金额------04
	Pin               string            //联机PINBLOCK--52
	MerchantId        string            //商户号--------42
	TerminalId        string            //终端号--------41
	Pan               string            //主账号--------02
	PanSeqNo          string            //卡片序列号-----
	CardExpireDate    string            //有效期--------14
	Track1            string            //磁道一--------
	Track2            string            //磁道二--------35
	PosEntryMode      string            //刷卡方式------22
	IccRelatedData    map[string]string //IC卡相关数据--
	AuthCode          string            //授权码-------38
	ResponseCode      string            //响应码-------39
	Invoice           string            //发票号-------62
	OriginalAmount    string            //原交易金额----
	OriginalTransType string            //原交易类型----
	TransType         string            //交易类型------
	CurrencyCode      string            //货币代码------
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
