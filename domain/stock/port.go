package stock

type StockRepository interface {
	GetAverage() int
	SetAvarage(value int)
	SetFinancialLoss(value int)
	GetFinancialLoss() int
	GetStockQuantity()
	SetStockQuantity(value int)
}

type UserCase interface {
	Buy(quantity, cost int) error
	Sell(quantity, cost int) (tax int, err error)
}
