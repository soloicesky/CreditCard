package BEA

import (
	"BEA/ISO8583"
	"BEA/TLV"
	"fmt"
	"testing"
)

func TestBEA(t *testing.T) {
	var transData TransactionData
	// BEA.Init()

	transData.TransId = "000001"
	transData.Invoice = "000001"
	transData.TransDate = "1221"
	transData.TransTime = "164030"
	transData.TransType = SALE
	transData.CurrencyCode = "0344"
	transData.MerchantId = "000015204000099"
	transData.TerminalId = "63150001"
	transData.Pan = ""
	transData.PanSeqNo = ""
	transData.CardExpireDate = "2512"
	transData.Track1 = ""
	transData.Track2 = "5413330056003578D251210100062001602"
	transData.PosEntryMode = "INSERT"

	iccData := make(map[string]string, 0)
	TLV.ParseConstructTLVMsg(ISO8583.Base16Decode("5F2A020344820200008407A0000000041010950500000000009A031712139B0200009C01009F02060000000066369F03060000000000009F090200029F1A0203449F1E0831323334353637389F3303E0B8C89F3501229F360200019F370485A04EA89F4104000000209F6002C3DE9F6102391C9F62060000000000389F6401059F63060000000007C69F650200E09F6701059F6602071E9F6A04000000609F6B125413330056003578D251210100062001602F"), iccData)

	transData.IccRelatedData = iccData
	var config Config
	config.TPDU = "7000280000"
	config.EDS = "0003000A00F000"
	config.Host = "192.168.22.188:8081"
	fmt.Printf("%+v\n", transData)
	transData, _ = Sale(transData, config)
}
