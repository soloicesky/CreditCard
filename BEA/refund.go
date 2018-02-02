package BEA

type RefundTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewRefund(trans *TransactionData, config *Config) *RefundTransaction {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	return &RefundTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		messageTypeId:  "0200",
		processingCode: "200000",
	}
}

func (refund *AuthorizeTransaction) Valid() error {
	if err := refund.baseValid(); err != nil {
		return err
	}
	return validMatch(refund.transData.Pan,
		refund.transData.Amount,
		refund.transData.TransId,
		refund.transData.CardExpireDate,
		refund.transData.Track2,
	)
}

func (refund *RefundTransaction) SetFields() {
	refund.baseFieldSet()
	refund.set(3, param[refund.transData.TransType].processingCode)
	refund.set(24, param[refund.transData.TransType].nii)
	refund.set(25, param[refund.transData.TransType].posCondictionCode)
}

func (refund *RefundTransaction) Fields() []uint8 {
	return refund.entryMap[auth.transData.PosEntryMode]
}

func (refund *RefundTransaction) Name() string {
	return refund.transData.TransType
}
