package BEA

/*
func DoRequest(transData *TransactionData, config *Config) (*TransactionData, error) {
	err := getSupportEntryMode(transData.PosEntryMode)
	if err != nil {
		return nil, fmt.Errorf("GetSupportEntryMode error: %s", err.Error())
	}

	fields, err := getFields(transData.TransType, transData.PosEntryMode)
	if err != nil {
		return nil, fmt.Errorf("GetFields error: %s", err.Error())
	}
	replyData, err := communicateWithHost(transData, config, fields)
	if err != nil {
		switch err {
		case RECV_ERR:
			transData.ResponseCode = BINDO_RECV_ERR
		case CONN_ERR:
			transData.ResponseCode = BINDO_CONN_ERR
		case SEND_ERR:
			transData.ResponseCode = BINDO_SEND_ERR
		default:
		}
		return nil, fmt.Errorf("communicateWithHost error: %s", err.Error())
	}
	return replyData, nil
}*/

func DoRequest(transData *TransactionData, config *Config) (*TransactionData, error) {
	//生成NewAuthorize结构体
	//Valid()
	//Fields()
	//删除多余的field
	//正式开始调用create message，进行交易
	return nil, nil
}

func getTransactionInterface(transData *TransactionData, config *Config) BEAInterface {
	switch transData.TransType {
	case KindPreAuthorize:
		return NewAuthorize(transData, config)
	}
	return nil
}
