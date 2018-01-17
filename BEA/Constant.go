package BEA

import (
	"errors"
)

var (
	CONN_ERR     error = errors.New("fail to connect to host")   //连接失败
	SEND_ERR     error = errors.New("fail to send data to host") //发送失败
	RECV_ERR     error = errors.New("fail to recv from host")    //接收失败
	INVALID_DATA error = errors.New("invalid data")              // 数据非法
)
