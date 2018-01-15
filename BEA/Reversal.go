package BEA

/**
	根据提供的现有交易数据执行一笔消费交易
	@param transData 交易数据
	@param config 后台配置参数
	如有出错则返回错误信息 err不为nil，否则err为nil 交易数据为当前交易数据
**/
func Reversal(transData TransactionData, config Config) (TransactionData, error) {
	var fields []byte
	switch transData.PosEntryMode {
	case INSERT:
		fields = []byte{0, 3, 4, 11, 22, 23, 24, 25, 35, 41, 42, 55, 62}
	case SWIPE:
		fields = []byte{0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62}
	case WAVE:
		fields = []byte{0, 3, 4, 11, 22, 23, 24, 25, 35, 41, 42, 55, 62}
	case FALLBACK:
		fields = []byte{0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62}
	case MSD:
		fields = []byte{0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62}
	case MANUAL:
		fields = []byte{0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62}
	}

	return CommunicationHost(transData, config, fields)
}