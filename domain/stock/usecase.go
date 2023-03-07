package stock

type UseCaseService struct{}

func calcTax(gain int) int {
	return (gain * 20) / 100
}

type calcWeightedAverageParams struct {
	CurrentStockQuantity   int
	StockBoughtQuantity    int
	CurrentWeightedAverage int
	BoughtStockUnitCost    int
}

func calcWeightedAverage(p calcWeightedAverageParams) int {
	return ((p.CurrentStockQuantity * p.CurrentWeightedAverage) + (p.StockBoughtQuantity * p.BoughtStockUnitCost)) / (p.CurrentStockQuantity + p.StockBoughtQuantity)
}

func NewUseCaseService() *UseCaseService {
	return &UseCaseService{}
}
