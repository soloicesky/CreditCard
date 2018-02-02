package BEA

type VoidTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewVoid(trans *TransactionData, config *Config) *VoidTransaction {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	trxn := &VoidTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0200",
		processingCode: "000000",
	}

	switch trans.OriginalTransType {
	case KindSale:
		trxn.processingCode = "020000"
	case KindRefund:
		trxn.processingCode = "220000"
	case KindPreAuthorize:
		trxn.processingCode = "000000"
	case KindPreAuthCompletion:
		trxn.processingCode = "000000"
	}
}

func (void *VoidTransaction) Valid() error {
	if err := auth.baseValid(); err != nil {
		return err
	}
	return validMatch(void.transData.Pan,
		void.transData.Amount,
		void.transData.TransId,
		void.transData.CardExpireDate,
		void.transData.Track2,
	)
}

func (void *VoidTransaction) SetFields() {
	void.baseFieldSet()
	void.set(3, param[void.transData.TransType].processingCode)
	void.set(24, param[void.transData.TransType].nii)
	void.set(25, param[void.transData.TransType].posCondictionCode)

}

func (void *VoidTransaction) Fields() []uint8 {
	return void.entryMap[auth.transData.PosEntryMode]
}

func (void *VoidTransaction) Name() string{
	return void.transData.TransType
}
