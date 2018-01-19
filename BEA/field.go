package BEA

type beaFieldMap map[EntryMode][]byte

var (
	preAuthorizationFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	saleFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	offlineSaleFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 37, 38, 41, 42, 55, 60, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 37, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 37, 38, 41, 42, 55, 60, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 37, 38, 41, 42, 60, 60, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 37, 38, 41, 42, 60, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 60, 62},
	}
	preAuthCompletionFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 37, 38, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 37, 38, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 38, 41, 42, 62},
	}

	refundFieldMap = beaFieldMap{
		INSERT:   {0, 3, 4, 11, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 22, 24, 25, 35, 37, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 22, 23, 24, 25, 35, 37, 38, 41, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 22, 24, 25, 35, 37, 38, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 22, 24, 25, 35, 37, 38, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 37, 38, 41, 42, 62},
	}

	voidRefundFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 41, 42, 62},
	}

	voidSaleFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 37, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 37, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 37, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 37, 41, 42, 62},
	}

	voidOfflineSaleFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 38, 41, 42, 62},
	}

	voidPreAuthFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 38, 41, 42, 62},
	}

	voidPreAuthCompletionFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 12, 13, 14, 22, 23, 24, 25, 35, 38, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 35, 38, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 12, 13, 14, 22, 24, 25, 38, 41, 42, 62},
	}
	batchUploadFieldMap = beaFieldMap{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	reversalFieldMap = beaFieldMap{
		INSERT:   {0, 3, 4, 11, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 3, 4, 11, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 3, 4, 11, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
)

var (
	logonField      = []byte{0, 3, 11, 24, 41}
	settlementField = []byte{0, 3, 11, 24, 41}
)

var transactionFieldMap = map[TransactionType]beaFieldMap{
	SALE:                  saleFieldMap,                  //消费参数
	VOIDSALE:              voidSaleFieldMap,              //消费撤销参数
	REFUND:                refundFieldMap,                //退货
	VOIDREFUND:            voidRefundFieldMap,            //撤销退货
	OFFLINESALE:           offlineSaleFieldMap,           //离线消费
	VOIDOFFLINESALE:       voidOfflineSaleFieldMap,       //离线消费撤销
	PREAUTH:               preAuthorizationFieldMap,      //预授权
	VOIDPREAUTH:           voidPreAuthFieldMap,           //预授权撤销
	PREAUTHCOMPLETION:     preAuthCompletionFieldMap,     //预授权完成
	VOIDPREAUTHCOMPLETION: voidPreAuthCompletionFieldMap, //预授权完成撤销
	REVERSAL:              reversalFieldMap,              //冲正
}
var managerFieldMap = map[TransactionType][]byte{
	LOGON:      logonField,
	SETTLEMENT: settlementField,
}

func getFields(transType TransactionType, mode EntryMode) ([]byte, error) {
	var fields []byte
	for k, v := range managerFieldMap {
		if k == transType {
			fields = v
			break
		}
	}
	if len(fields) == 0 {
		fields = transactionFieldMap[transType][mode]
	}

	if len(fields) > 0 {
		return fields, nil
	} else {
		return nil, fmt.Errorf("get fields of %s ,type of %s error", transType, mode)
	}
}
