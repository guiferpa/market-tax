package stock

type StockRepository interface {
	SetWeightAvarage(value int)
	GetWeightAverage() int
	SetFinancialLoss(value int)
	GetFinancialLoss() int
	SetStockQuantity(value int)
	GetStockQuantity() int
}

type UserCase interface {
	Buy(quantity, cost int) error
	Sell(quantity, cost int) (tax int, err error)
}
