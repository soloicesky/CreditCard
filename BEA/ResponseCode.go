package BEA

const (
	APPROVED                       string = "00" //交易批准
	CALL_ISSUER                    string = "01" //请联系放卡行
	REFFERAL                       string = "02" //参考交易，请联系发卡行
	INVALID_MERCHANT               string = "03" //无效商户
	PICK_UP_CARD                   string = "04" //没收无效卡
	DO_NOT_HONOUR                  string = "05" //不允许，不接受
	INVALID_TRANSACTION            string = "12" //无效交易
	INVALID_AMOUNT                 string = "13" //无效金额
	INVALID_CARD_NUMBER            string = "14" //无效卡号
	REENTER_TRANSACTION            string = "19" //重复交易
	NO_TRANSACTIONS                string = "21" //找不到交易
	UNABLE_LOCATE_RECORD           string = "25" //
	FORMAT_ERROR                   string = "30"
	BANK_NORT_SUPPORTED            string = "31"
	LOST_CARD                      string = "41"
	STOLEN_CARD                    string = "43"
	INSUFFICIENT_FUNDS             string = "51"
	NO_CHEQUING_ACCOUNT            string = "52"
	NO_SAVINGS_ACCOUNT             string = "53"
	EXPIRED_CARD                   string = "54"
	INCORRECT_PIN                  string = "55"
	NO_CARD_RECORD                 string = "56"
	NOT_PERMITTED                  string = "58" //交易批准
	EXCEED_WITHDRAWAL_AMOUNT_LIMIT string = "61"
	SECURITY_VIOLATION             string = "63"
	PIN_TRIES_EXCEED               string = "75"
	INVALID_PRODUCT_CODE           string = "76"
	RECONCILE_ERROR                string = "77"
	TRANS_NOT_FOUND                string = "78"
	BATCH_ALREADY_OPEN             string = "79"
	BATCH_NUMBER_NOT_FOUND         string = "80"
	BATCH_NOT_FOUND                string = "85"
	BAT_TERMINAL_ID                string = "89"
	ISSUER_OR_SWITCH_INOPERATIVE   string = "91"
	DUPLICATE_TRANSMISSION         string = "94"
	BATCH_TRANSFER                 string = "95"
	SYSTEM_MALFUNCTION             string = "96"

	//Locally Generated Error Messages
	LE_LOST_CARRIER                  string = "LC"
	LE_COMM_ERROR                    string = "CE"
	LE_INVALID_DOWNLOADLINE_LOAD     string = "ID"
	LE_INVALID_AMOUNT                string = "IA"
	LE_INVALID_MESSAGE_TYPE          string = "IR"
	LE_INVALID_HOST_SEQUENCE_NUMBER  string = "IS"
	LE_INVALID_MAC                   string = "IM"
	LE_NO_ERPLY_TIMEOUT              string = "TO"
	LE_ADIVICE_REVERSAL_NOT_APPROVED string = "ND"

	//bindo generated error code
	BINDO_COMM_ERROR				string ="DE"
	BINDO_RECV_FAILED				string ="RF"
)

var rcDict = map[string]string{
	APPROVED:                       "Approved",
	CALL_ISSUER:                    "Refer to card issuer",
	REFFERAL:                       "Refer to card Issuer\"s special conditions", //参考交易，请联系发卡行
	INVALID_MERCHANT:               "Invalid merchant",                           //无效商户
	PICK_UP_CARD:                   "Pick-up",                                    //没收无效卡
	DO_NOT_HONOUR:                  "Do not honour",                              //不允许，不接受
	INVALID_TRANSACTION:            "Invalid transaction",                        //无效交易
	INVALID_AMOUNT:                 "Invalid amount",                             //无效金额
	INVALID_CARD_NUMBER:            "Invalid card number",                        //无效卡号
	REENTER_TRANSACTION:            "Re-enter transaction",                       //重复交易
	NO_TRANSACTIONS:                "No Transactions",                            //找不到交易
	UNABLE_LOCATE_RECORD:           "Unable to locate record on file",            //
	FORMAT_ERROR:                   "Format error",
	BANK_NORT_SUPPORTED:            "Bank not supported by switch",
	LOST_CARD:                      "Lost card",
	STOLEN_CARD:                    "Stolen card, pick up",
	INSUFFICIENT_FUNDS:             "Not sufficient funds",
	NO_CHEQUING_ACCOUNT:            "No chequing account",
	NO_SAVINGS_ACCOUNT:             "No savings account",
	EXPIRED_CARD:                   "Expired card",
	INCORRECT_PIN:                  "Incorrect PIN",
	NO_CARD_RECORD:                 "No card record",
	NOT_PERMITTED:                  "Transaction not permitted to terminal", //交易批准
	EXCEED_WITHDRAWAL_AMOUNT_LIMIT: "Exceeds withdrawal amount limit",
	SECURITY_VIOLATION:             "Security violation",
	PIN_TRIES_EXCEED:               "Allowable number of PIN tries exceeded",
	INVALID_PRODUCT_CODE:           "Invalid product code",
	RECONCILE_ERROR:                "Reconcile error (or host text if sent)",
	TRANS_NOT_FOUND:                "Trans. number not found",
	BATCH_ALREADY_OPEN:             "Batch already open",
	BATCH_NUMBER_NOT_FOUND:         "Batch number not found",
	BATCH_NOT_FOUND:                "Batch not found",
	BAT_TERMINAL_ID:                "Bad Terminal ID",
	ISSUER_OR_SWITCH_INOPERATIVE:   "Issuer or switch inoperative",
	DUPLICATE_TRANSMISSION:         "Duplicate transmission",
	BATCH_TRANSFER:                 "Reconcile error. Batch upload started",
	SYSTEM_MALFUNCTION:             "System malfunction",

	//LocallyGeneratedErrorMessages
	LE_LOST_CARRIER:                  "Lost carrier",
	LE_COMM_ERROR:                    "Communications error",
	LE_INVALID_DOWNLOADLINE_LOAD:     "Invalid downline load",
	LE_INVALID_AMOUNT:                "Invalid amount",
	LE_INVALID_MESSAGE_TYPE:          "Invalid message type",
	LE_INVALID_HOST_SEQUENCE_NUMBER:  "nvalid host sequence number",
	LE_INVALID_MAC:                   "Invalid MAC",
	LE_NO_ERPLY_TIMEOUT:              "No reply timeout",
	LE_ADIVICE_REVERSAL_NOT_APPROVED: "Advice / Reversal transactions not approved",

		//bindo generated error code
	BINDO_COMM_ERROR:				"fail to connect to acquirer",
	BINDO_RECV_FAILED:				"fail to receive response message",
}

var SuccessCodeList []string = []string{
	APPROVED,
}

var FailedCodeList []string = []string{
	CALL_ISSUER,
	REFFERAL,
	INVALID_MERCHANT,
	PICK_UP_CARD,
	DO_NOT_HONOUR,
	INVALID_TRANSACTION,
	INVALID_AMOUNT,
	INVALID_CARD_NUMBER,
	REENTER_TRANSACTION,
	NO_TRANSACTIONS,
	UNABLE_LOCATE_RECORD,
	FORMAT_ERROR,
	BANK_NORT_SUPPORTED,
	LOST_CARD,
	STOLEN_CARD,
	INSUFFICIENT_FUNDS,
	NO_CHEQUING_ACCOUNT,
	NO_SAVINGS_ACCOUNT,
	EXPIRED_CARD,
	INCORRECT_PIN,
	NO_CARD_RECORD,
	NOT_PERMITTED,
	EXCEED_WITHDRAWAL_AMOUNT_LIMIT,
	SECURITY_VIOLATION,
	PIN_TRIES_EXCEED,
	INVALID_PRODUCT_CODE,
	RECONCILE_ERROR,
	TRANS_NOT_FOUND,
	BATCH_ALREADY_OPEN,
	BATCH_NUMBER_NOT_FOUND,
	BATCH_NOT_FOUND,
	BAT_TERMINAL_ID,
	ISSUER_OR_SWITCH_INOPERATIVE,
	DUPLICATE_TRANSMISSION,
	BATCH_TRANSFER,
	SYSTEM_MALFUNCTION,
	LE_LOST_CARRIER,
	LE_COMM_ERROR,
	LE_INVALID_DOWNLOADLINE_LOAD,
	LE_INVALID_AMOUNT,
	LE_INVALID_MESSAGE_TYPE,
	LE_INVALID_HOST_SEQUENCE_NUMBER,
	LE_INVALID_MAC,
	LE_NO_ERPLY_TIMEOUT,
	LE_ADIVICE_REVERSAL_NOT_APPROVED,
	BINDO_COMM_ERROR,
}

var ReversalCodeList []string = []string{
	BINDO_RECV_FAILED,
}