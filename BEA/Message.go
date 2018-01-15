package BEA

import (
	"ISO8583"
	"TLV"
	_ "crypto/cipher"
	"crypto/des"
	"fmt"
)

const (
	AUTHORIZATION string = "AUTHORIZATION"
	REVERSAL      string = "REVERSAL"
)

//交易类型常量
const (
	LOGON                 string = "LOGON"                 //签到
	SALE                  string = "SALE"                  //消费
	VOIDSALE              string = "VOIDSALE"              //消费撤销
	REFUND                string = "REFUND"                //退货
	VOIDREFUND            string = "VOIDREFUND"            //退货撤销
	OFFLINESALE           string = "OFFLINESALE"           //离线消费
	VOIDOFFLINESALE       string = "VOIDOFFLINESALE"       //离线消费撤销
	PREAUTH               string = "PREAUTH"               //预授权
	PREAUTHCOMPLETION     string = "PREAUTHCOMPLETION"     //预授权完成
	VOIDPREAUTHCOMPLETION string = "VOIDPREAUTHCOMPLETION" //预授权完成撤销
	SETTLEMENT            string = "SETTLEMENT"            //结算
	BATCHUPLOAD           string = "BATCHUPLOAD"           //批上送
)

//刷卡方式
const (
	INSERT   string = "INSERT"   //插IC卡
	SWIPE    string = "SWIPE"    //刷磁条卡
	WAVE     string = "WAVE"     //挥非接卡
	FALLBACK string = "FALLBACK" //降级
	MSD      string = "MSD"      //非接卡MSD模式
	MANUAL   string = "MANUAL"   //手输卡号
)

//消息参数
type msgParam struct {
	id                string //消息类型
	processingCode    string //处理码
	nii               string //网络信息指示
	posCondictionCode string //POS条件码
}

//消息参数表
var param = map[string]msgParam{
	LOGON:                 {"0800", "920000", "028", "00"}, //消费参数
	SALE:                  {"0200", "000000", "028", "00"}, //消费参数
	VOIDSALE:              {"0200", "020000", "028", "00"}, //消费撤销参数
	REFUND:                {"0200", "200000", "028", "00"}, //退货
	VOIDREFUND:            {"0200", "220000", "028", "00"}, //撤销退货
	OFFLINESALE:           {"0220", "000000", "028", "00"}, //离线消费
	PREAUTH:               {"0100", "000000", "028", "00"}, //预授权
	PREAUTHCOMPLETION:     {"0220", "000000", "028", "00"}, //预授权完成
	VOIDPREAUTHCOMPLETION: {"0220", "000000", "028", "00"}, //预授权完成
	REVERSAL:              {"0400", "000000", "028", "00"}, //预授权完成
}

//终端输入模式码
var posEntryMode = map[string]string{
	INSERT:   "05", //插卡
	SWIPE:    "90", //刷卡
	WAVE:     "07", //挥卡
	FALLBACK: "80", //降级
	MSD:      "91", //非接磁卡模式
}

//加密ISO8583消息
func encryptISO8583Message(msg []byte) []byte {
	keyV1 := ISO8583.Base16Decode("ABCDEF0123456789EEEEEEEEEEEEEEEE")
	keyV2 := ISO8583.Base16Decode("FFFFFFFFFFFFFFFF9876543210FEDCBA")
	var key []byte

	for i := 0; i < len(keyV1); i++ {
		key = append(key, keyV1[i]^keyV2[i])
	}

	var tripleDESKey []byte
	tripleDESKey = append(tripleDESKey, key[:16]...)
	tripleDESKey = append(tripleDESKey, key[:8]...)

	tdesCipher, err := des.NewTripleDESCipher(tripleDESKey)

	if err != nil {
		panic(err)
	}

	round := (len(msg) + 7) / 8
	dstLen := round * 8
	fmt.Printf("msg len:%d,round:%d,dstLen:%d\r\n", len(msg), round, dstLen)
	plainMsg := make([]byte, dstLen)
	copy(plainMsg, msg)
	encryptedMsg := make([]byte, 0)
	encBlock := make([]byte, 8)
	var plainBlock []byte

	for i := 0; i < round; i++ {
		plainBlock = plainMsg[i*8 : i*8+8]
		tdesCipher.Encrypt(encBlock, plainBlock)
		encryptedMsg = append(encryptedMsg, encBlock...)
	}

	return encryptedMsg
}

/**
	根据给定的交易数据和需要打包的位图信息生成ISO8583 报文

**/
func CreateIISO8583Message(transData TransactionData, fields []byte, config Config) ([]byte, error) {

	ISO8583.PushElement(0, param[transData.TransType].id)
	ISO8583.PushElement(2, transData.Pan)
	ISO8583.PushElement(3, param[transData.TransType].processingCode)
	ISO8583.PushElement(4, fmt.Sprintf("%012s", transData.Amount))
	ISO8583.PushElement(11, fmt.Sprintf("%06s", transData.TransId))
	ISO8583.PushElement(14, transData.CardExpireDate)

	if ISO8583.StringIsEmpty(transData.Pin) {
		ISO8583.PushElement(22, posEntryMode[transData.PosEntryMode]+"2")
	} else {
		ISO8583.PushElement(22, posEntryMode[transData.PosEntryMode]+"1")
	}

	if !ISO8583.StringIsEmpty(transData.PanSeqNo) {
		ISO8583.PushElement(23, fmt.Sprintf("%04s", transData.PanSeqNo))
	}

	switch transData.TransType {
	case REVERSAL:
		ISO8583.PushElement(24, param[transData.OriginalTransType].nii)
		ISO8583.PushElement(25, param[transData.OriginalTransType].posCondictionCode)
	default:
		ISO8583.PushElement(24, param[transData.TransType].nii)
		ISO8583.PushElement(25, param[transData.TransType].posCondictionCode)
	}

	if !ISO8583.StringIsEmpty(transData.AcquireTransID) {
		ISO8583.PushElement(37, transData.AcquireTransID)
	}

	if !ISO8583.StringIsEmpty(transData.TerminalId) {
		ISO8583.PushElement(41, transData.TerminalId)
	}

	if !ISO8583.StringIsEmpty(transData.MerchantId) {
		ISO8583.PushElement(42, transData.MerchantId)
	}

	DE55 := TLV.BuildConstructTLVMsg(transData.IccRelatedData)
	ISO8583.PushElement(55, ISO8583.Base16Encode(DE55))

	if !ISO8583.StringIsEmpty(transData.OriginalAmount) {
		ISO8583.PushElement(60, fmt.Sprintf("%012s", transData.OriginalAmount))
	}

	if !ISO8583.StringIsEmpty(transData.Invoice) {
		ISO8583.PushElement(62, transData.Invoice)
	}

	msg, err := ISO8583.PrepareISO8583Message(fields)

	if err != nil {
		return nil, err
	}

	// msg = encryptISO8583Message(msg)

	dstMsg := make([]byte, 0)
	dstMsg = append(dstMsg, 0x00, 0x00) // len
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.TPDU)...)
	dstMsg = append(dstMsg, ISO8583.Base16Decode(config.EDS)...)
	dstMsg = append(dstMsg, msg...)
	dstMsg[2+5+4] = byte(len(msg) >> 8)
	dstMsg[2+5+5] = byte(len(msg) & 0x00FF)

	dstMsg[0] = byte((len(dstMsg) - 2) >> 8)
	dstMsg[1] = byte((len(dstMsg) - 2) & 0x00FF)

	return dstMsg, err
}
