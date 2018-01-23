package sdk

import (
	"fmt"
)

type TransactionType string

//交易类型常量
const (
	LOGON                          TransactionType = "LOGON"                          //签到
	SALE                           TransactionType = "SALE"                           //消费
	VOIDSALE                       TransactionType = "VOIDSALE"                       //消费撤销
	REFUND                         TransactionType = "REFUND"                         //退货
	VOIDREFUND                     TransactionType = "VOIDREFUND"                     //退货撤销
	OFFLINESALE                    TransactionType = "OFFLINESALE"                    //离线消费
	VOIDOFFLINESALE                TransactionType = "VOIDOFFLINESALE"                //离线消费撤销
	PREAUTH                        TransactionType = "PREAUTH"                        //预授权
	VOIDPREAUTH                    TransactionType = "VOIDPREAUTH"                    //预授权
	PREAUTHCOMPLETION              TransactionType = "PREAUTHCOMPLETION"              //预授权完成
	VOIDPREAUTHCOMPLETION          TransactionType = "VOIDPREAUTHCOMPLETION"          //预授权完成撤销
	SETTLEMENT                     TransactionType = "SETTLEMENT"                     //结算
	SETTLEMENTAFTERUPLOAD          TransactionType = "SETTLEMENTAFTERUPLOAD"          //批上送后结算
	BATCHUPLOAD                    TransactionType = "BATCHUPLOAD"                    //批上送
	BATCHUPLOADLAST                TransactionType = "BATCHUPLOADLAST"                //批上送
	REVERSAL                       TransactionType = "REVERSAL"                       //冲正
	OFFLINEUPLOAD                  TransactionType = "OFFLINEUPLOAD"                  //离线上送
	VOIDADJUSTSALE                 TransactionType = "VOIDADJUSTSALE"                 //撤销消费调整
	ADJUSTSALE                     TransactionType = "ADJUSTSALE"                     //消费调整
	ADJUSTOFFLINESALE              TransactionType = "ADJUSTOFFLINESALE"              //离线消费调整
	VOIDADJUSTOFFLINESALE          TransactionType = "VOIDADJUSTOFFLINESALE"          //离线消费调整撤销
	VOIDUPLOADADJUSTOFFLINESALE    TransactionType = "VOIDUPLOADADJUSTOFFLINESALE"    //上送离线消费调整撤销
	ADJUSTUPLOADOFFLINESALE        TransactionType = "ADJUSTUPLOADOFFLINESALE"        //上送离线消费调整撤销
	VOIDADJUSTUPLOADOFFLINESALE    TransactionType = "VOIDADJUSTUPLOADOFFLINESALE"    //上送离线消费调整撤销
	ADJUSTUPLOADADJUSTSALE         TransactionType = "ADJUSTUPLOADADJUSTSALE"         //上送离线消费调整撤销
	VOIDADJUSTUPLOADADJUSTSALE     TransactionType = "VOIDADJUSTUPLOADADJUSTSALE"     //上送离线消费调整撤销
	ADJUSTSALECOMPLETION           TransactionType = "ADJUSTSALECOMPLETION"           //上送离线消费调整撤销
	VOIDADJUSTSALECOMPLETION       TransactionType = "VOIDADJUSTSALECOMPLETION"       //上送离线消费调整撤销
	VOIDUPLOADADJUSTSALECOMPLETION TransactionType = "VOIDUPLOADADJUSTSALECOMPLETION" //上送离线消费调整撤销
	ADJUSTREFUND                   TransactionType = "ADJUSTREFUND"                   //调整退货
	VOIDADJUSTREFUND               TransactionType = "VOIDADJUSTREFUND"               //调整退货撤销
	VOIDUPLOADADJUSTREFUND         TransactionType = "VOIDUPLOADADJUSTREFUND"         //上送退货调整撤销
)

type EntryMode string

//刷卡方式
const (
	INSERT   EntryMode = "INSERT"   //插IC卡
	SWIPE    EntryMode = "SWIPE"    //刷磁条卡
	WAVE     EntryMode = "WAVE"     //挥非接卡
	FALLBACK EntryMode = "FALLBACK" //降级
	MSD      EntryMode = "MSD"      //非接卡MSD模式
	MANUAL   EntryMode = "MANUAL"   //手输卡号
)

//终端输入模式码
var posEntryMode = map[EntryMode]string{
	INSERT:   "05", //插卡
	SWIPE:    "90", //刷卡
	WAVE:     "07", //挥卡
	FALLBACK: "80", //降级
	MSD:      "91", //非接磁卡模式
}

//消息参数
type msgParam struct {
	id                string //消息类型
	processingCode    string //处理码
	nii               string //网络信息指示
	posCondictionCode string //POS条件码
}

//消息参数表
var param = map[TransactionType]msgParam{
	LOGON:                          {"0800", "920000", "028", "00"}, //消费参数
	SALE:                           {"0200", "000000", "028", "00"}, //消费参数
	VOIDSALE:                       {"0200", "020000", "028", "00"}, //消费撤销参数
	REFUND:                         {"0200", "200000", "028", "00"}, //退货
	VOIDREFUND:                     {"0200", "220000", "028", "00"}, //撤销退货
	OFFLINESALE:                    {"0220", "000000", "028", "00"}, //离线消费
	VOIDOFFLINESALE:                {"0220", "000000", "028", "00"}, //离线消费撤销
	PREAUTH:                        {"0100", "000000", "028", "00"}, //预授权
	VOIDPREAUTH:                    {"0120", "000000", "028", "00"}, //预授权撤销
	PREAUTHCOMPLETION:              {"0220", "000000", "028", "00"}, //预授权完成
	VOIDPREAUTHCOMPLETION:          {"0220", "000000", "028", "00"}, //预授权完成撤销
	REVERSAL:                       {"0400", "000000", "028", "00"}, //冲正
	SETTLEMENT:                     {"0500", "920000", "028", "00"}, //结算
	SETTLEMENTAFTERUPLOAD:          {"0500", "960000", "028", "00"}, //批上送后结算
	BATCHUPLOAD:                    {"0320", "000000", "028", "00"}, //批上送
	BATCHUPLOADLAST:                {"0320", "000000", "028", "00"}, //批上送结束
	VOIDADJUSTSALE:                 {"0200", "000000", "028", "00"}, //消费调整撤销
	ADJUSTSALE:                     {"0220", "000000", "028", "00"}, //消费调整
	ADJUSTOFFLINESALE:              {"0220", "000000", "028", "00"}, //离线消费调整
	VOIDADJUSTOFFLINESALE:          {"0220", "000000", "028", "00"}, //离线消费调整撤销
	VOIDUPLOADADJUSTOFFLINESALE:    {"0220", "000000", "028", "00"}, //离线消费调整上送撤销
	ADJUSTUPLOADOFFLINESALE:        {"0220", "000000", "028", "00"}, //离线消费上送调整
	VOIDADJUSTUPLOADOFFLINESALE:    {"0220", "000000", "028", "00"}, //离线消费上送调整撤销
	ADJUSTUPLOADADJUSTSALE:         {"0220", "020000", "028", "00"}, //离线调整上送调整
	VOIDADJUSTUPLOADADJUSTSALE:     {"0220", "020000", "028", "00"}, //消费调整上送调整撤销
	ADJUSTSALECOMPLETION:           {"0220", "000000", "028", "00"}, //预授权完成调整
	VOIDADJUSTSALECOMPLETION:       {"0220", "000000", "028", "00"}, //预授权完成调整撤销
	VOIDUPLOADADJUSTSALECOMPLETION: {"0220", "020000", "028", "00"}, //预授权完成调整上送撤销
	ADJUSTREFUND:                   {"0220", "220000", "028", "00"}, //退货调整
	VOIDADJUSTREFUND:               {"0220", "220000", "028", "00"}, //退货调整撤销
	VOIDUPLOADADJUSTREFUND:         {"0220", "220000", "028", "00"}, //退货调整上送撤销
}

func getAllEntryModes() []EntryMode {
	var modes []EntryMode
	modes = append(modes, INSERT)
	modes = append(modes, SWIPE)
	modes = append(modes, WAVE)
	modes = append(modes, FALLBACK)
	modes = append(modes, MSD)
	modes = append(modes, MANUAL)
	return modes
}

func getSupportEntryMode(mode EntryMode) error {
	modes := getAllEntryModes()
	for _, v := range modes {
		if mode == v {
			return nil
		}
	}
	return fmt.Errorf("not support pos_entry_mode: %s", mode)
}

//func GetAllSupportTransTypes() []TransactionType {
//	var types []TransactionType
//	types = append(types, LOGON)
//	types = append(types, SALE)
//	types = append(types, VOIDSALE)
//	types = append(types, REFUND)
//	types = append(types, VOIDREFUND)
//	types = append(types, OFFLINESALE)
//	types = append(types, VOIDOFFLINESALE)
//	types = append(types, PREAUTH)
//	types = append(types, VOIDPREAUTH)
//	types = append(types, PREAUTHCOMPLETION)
//	types = append(types, VOIDPREAUTHCOMPLETION)
//	types = append(types, SETTLEMENT)
//	types = append(types, BATCHUPLOAD)
//	types = append(types, REVERSAL)
//	return types
//}

//func getAllSupportTransTypes() []TransactionType {
//	var types []TransactionType
//	for k, _ := range fieldMap {
//		types = append(types, k)
//	}
//	return types
//}
