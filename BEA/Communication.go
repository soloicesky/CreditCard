package BEA

import (
	"CreditCard/ISO8583"
	"CreditCard/TLV"
	"errors"
	"fmt"
	"net"
	"time"
)

//和后台通信发送授权请求报文并接收授权响应报文
func communicateWithHost(reqMsg []byte, config Config, timeoutS int) ([]byte, error) {
	var count int
	rspMsg := make([]byte, 0)
	fmt.Println("connet host: ", config.Host)
	conn, err := net.Dial("tcp", config.Host)

	if err != nil {
		fmt.Println("fail to connect to", config.Host)
		return rspMsg, CONN_ERR
	}

	conn.SetReadDeadline(time.Now().Add(time.Duration(timeoutS) * time.Second))
	count, err = conn.Write(reqMsg)

	if err != nil {
		return rspMsg, SEND_ERR
	}

	totalLen := 0 //保存数据长度
	buf := make([]byte, 128)

	for {
		count, err = conn.Read(buf)

		if err != nil {
			fmt.Printf("read error:%s\r\n", err)
			return rspMsg, RECV_ERR
		}

		rspMsg = append(rspMsg, buf[0:count]...)

		if len(rspMsg) >= 2 {
			totalLen = 2 + int(rspMsg[0]<<8) + int(rspMsg[1])
		}

		if totalLen > 0 && len(rspMsg) >= totalLen {
			break
		}
	}

	return rspMsg, nil
}

/**
	保存数据
	fieldId 域标识
	value  值
	storage 存储位置
**/
func SaveData(fieldId int, value string, storage interface{}) error {
	transData, OK := storage.(TransactionData)

	if !OK {
		return errors.New("interface is not a type of TransactionData")
	}

	fmt.Printf("id:%d value:%s\r\n", fieldId, value)

	switch fieldId {
	case 37:
		transData.AcquireTransID = value
	case 38:
		transData.AuthCode = value
	case 39:
		transData.ResponseCode = value
	case 55:
		de55 := ISO8583.Base16Decode(value)
		TLV.ParseConstructTLVMsg(de55, transData.IccRelatedData)
	}

	return nil
}

/**
	创建并发送一个授权报文
	transData 交易数据
	config 配置参数
	fields 域集合
**/
func CommunicationHost(transData TransactionData, config Config, fields []byte) (TransactionData, error) {
	msg, err := CreateIISO8583Message(transData, fields, config)

	if err != nil {
		fmt.Printf("err:%v\r\n", err)
		return transData, err
	}

	fmt.Printf("Final Msg:%s\r\n", ISO8583.Base16Encode(msg))
	msg, err = communicateWithHost(msg, config, 30)

	if err != nil {
		switch err{
			case RECV_ERR:
				transData.ResponseCode = BINDO_RECV_FAILED
			case CONN_ERR:
				fallthrough
			case SEND_ERR:
				transData.ResponseCode = BINDO_COMM_ERROR
		default:
		}
		return transData, err
	}

	fmt.Printf("reponse ISO8583:%s\r\n", ISO8583.Base16Encode(msg))
	err = ISO8583.DecodeISO8583Message(msg[2+5+7:], SaveData, transData)

	if err!=nil {
		transData.ResponseCode = BINDO_RECV_FAILED
	}

	return transData, err
}
