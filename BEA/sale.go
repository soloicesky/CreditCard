package BEA

type SaleTransaction struct {
	entryMap map[EntryMode][]uint8
	BaseElement
	messageTypeId  string
	processingCode string
}

func NewSale(trans *TransactionData, config *Config) (*SaleTransaction, error) {
	fieldMap := map[EntryMode][]uint8{
		INSERT:   {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		SWIPE:    {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		WAVE:     {0, 2, 3, 4, 11, 14, 22, 23, 24, 25, 35, 41, 42, 55, 62},
		FALLBACK: {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MSD:      {0, 2, 3, 4, 11, 14, 22, 24, 25, 35, 41, 42, 62},
		MANUAL:   {0, 2, 3, 4, 11, 14, 22, 24, 25, 41, 42, 62},
	}

	return &SaleTransaction{
		entryMap: fieldMap,
		BaseElement: BaseElement{
			transData: trans,
			config:    config,
		},
		"0200",
		"000000",
	}, nil
}

func (sale *SaleTransaction) Valid() error {
	if err := sale.baseValid(); err != nil {
		return err
	}
	return validMatch(sale.transData.Pan,
		sale.transData.Amount,
		sale.transData.TransId,
		sale.transData.CardExpireDate,
		sale.transData.Track2,
	)
}

func (sale *SaleTransaction) SetFields() {
	sale.baseFieldSet()
	sale.set(3, param[sale.transData.TransType].processingCode)
	sale.set(24, param[sale.transData.TransType].nii)
	sale.set(25, param[sale.transData.TransType].posCondictionCode)
}

func (sale *SaleTransaction) Fields() []uint8 {
	return sale.entryMap[sale.transData.PosEntryMode]
}
