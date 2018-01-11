package BEA

const(
	CONN_ERR  error = error.new("fail to connect to host")
	SEND_ERR  error = error.new("fail to send data to host")
	RECV_ERR  error = error.new("fail to recv from host")
)