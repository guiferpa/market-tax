package stock

type StockRepository interface {
	GetWeightAverage() int
	SetWeightAvarage(value int)
	SetFinancialLoss(value int)
	GetFinancialLoss() int
	GetStockQuantity() int
	SetStockQuantity(value int)
}

type UserCase interface {
	Buy(quantity, cost int) error
	Sell(quantity, cost int) (tax int, err error)
}
