package sdk

import (
	"bindolabs/gateway/services/bea/ISO8583"
	"bindolabs/gateway/services/bea/TLV"
	_ "crypto/cipher"
	"crypto/des"
	"fmt"
)

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
func createIISO8583Message(transData *TransactionData, fields []byte, config *Config) ([]byte, error) {
	ISO8583.SetElement(0, param[transData.TransType].id)
	ISO8583.SetElement(2, transData.Pan)
	if transData.TransType == KindReversal {
		ISO8583.SetElement(3, param[transData.OriginalTransType].processingCode)
		ISO8583.SetElement(24, param[transData.OriginalTransType].nii)
		ISO8583.SetElement(25, param[transData.OriginalTransType].posCondictionCode)
	} else {
		ISO8583.SetElement(3, param[transData.TransType].processingCode)
		ISO8583.SetElement(24, param[transData.TransType].nii)
		ISO8583.SetElement(25, param[transData.TransType].posCondictionCode)
	}

	ISO8583.SetElement(4, fmt.Sprintf("%012s", transData.Amount))
	ISO8583.SetElement(11, fmt.Sprintf("%06s", transData.TransId))
	ISO8583.SetElement(14, transData.CardExpireDate)

	if ISO8583.StringIsEmpty(transData.Pin) {
		ISO8583.SetElement(22, posEntryMode[transData.PosEntryMode]+"2")
	} else {
		ISO8583.SetElement(22, posEntryMode[transData.PosEntryMode]+"1")
	}

	if !ISO8583.StringIsEmpty(transData.PanSeqNo) {
		ISO8583.SetElement(23, fmt.Sprintf("%04s", transData.PanSeqNo))
	}

	if !ISO8583.StringIsEmpty(transData.AcquireTransID) {
		ISO8583.SetElement(37, transData.AcquireTransID)
	}

	if !ISO8583.StringIsEmpty(transData.TerminalId) {
		ISO8583.SetElement(41, transData.TerminalId)
	}

	if !ISO8583.StringIsEmpty(transData.MerchantId) {
		ISO8583.SetElement(42, transData.MerchantId)
	}

	DE55 := TLV.BuildConstructTLVMsg(transData.IccRelatedData)
	ISO8583.SetElement(55, ISO8583.Base16Encode(DE55))

	if !ISO8583.StringIsEmpty(transData.OriginalAmount) {
		ISO8583.SetElement(60, fmt.Sprintf("%012s", transData.OriginalAmount))
	}

	if !ISO8583.StringIsEmpty(transData.Invoice) {
		ISO8583.SetElement(62, transData.Invoice)
	}

	batchTotal := fmt.Sprintf("%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d%03d%012d",
		transData.Batchtotals.CapturedSalesCount, transData.Batchtotals.CapturedSalesAmount,
		transData.Batchtotals.CapturedRefundCount, transData.Batchtotals.CapturedRefundAmount,
		transData.Batchtotals.DebitSalesCount, transData.Batchtotals.DebitSalesAmount,
		transData.Batchtotals.DebitRefundCount, transData.Batchtotals.DebitRefundAmount,
		transData.Batchtotals.AuthorizeSalesCount, transData.Batchtotals.AuthorizeSalesAmount,
		transData.Batchtotals.AuthorizeRefundCount, transData.Batchtotals.AuthorizeRefundAmount)

	ISO8583.SetElement(63, batchTotal)
	msg, err := ISO8583.PrepareISO8583Message(fields)

	if err != nil {
		return nil, fmt.Errorf("ISO8583::PrepareISO8583Message error: %s", err.Error())
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

	return dstMsg, nil
}
