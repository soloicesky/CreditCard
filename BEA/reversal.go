package BEA

type ReversalTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
}

func NewAuthorize(trans *TransactionData, config *Config) *ReversalTransaction {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}
	return &ReversalTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
	}
}

func (reversal *ReversalTransaction) Valid() error {
	if err := reversal.baseValid(); err != nil {
		return err
	}
	return validMatch(reversal.transData.Pan,
		reversal.transData.Amount,
		reversal.transData.TransId,
		reversal.transData.CardExpireDate,
		reversal.transData.Track2,
	)
}

func (reversal *ReversalTransaction) SetFields() {
	reversal.baseFieldSet()
	reversal.set(3, param[reversal.transData.TransType].processingCode)
	reversal.set(24, param[reversal.transData.TransType].nii)
	reversal.set(25, param[reversal.transData.TransType].posCondictionCode)
}

func (reversal *ReversalTransaction) Fields() []uint8 {
	return reversal.entryMap[reversal.transData.PosEntryMode]
}
