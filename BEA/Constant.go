package BEA

import (
	"errors"
)

var (
	CONN_ERR error = errors.New("fail to connect to host")
	SEND_ERR error = errors.New("fail to send data to host")
	RECV_ERR error = errors.New("fail to recv from host")
)
