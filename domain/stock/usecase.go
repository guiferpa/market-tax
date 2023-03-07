package stock

import "fmt"

type StockInvalidBuyError struct {
	Quantity int
	Cost     int
}

func (err *StockInvalidBuyError) Error() string {
	return fmt.Sprintf("invalid buy with quantity equals %d and cost %d", err.Quantity, err.Cost)
}

var (
	ErrStockBuyWithInvalidQuantity = &StockInvalidBuyError{}
	ErrStockBuyWithInvalidCost     = &StockInvalidBuyError{}
)

type UseCaseService struct {
	repository StockRepository
}

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

func (s *UseCaseService) Buy(quantity, cost int) error {
	if quantity < 1 {
		ErrStockBuyWithInvalidQuantity.Quantity = quantity
		return ErrStockBuyWithInvalidQuantity
	}

	if cost < 1 {
		ErrStockBuyWithInvalidCost.Cost = cost
		return ErrStockBuyWithInvalidCost
	}

	params := calcWeightedAverageParams{
		CurrentStockQuantity:   s.repository.GetStockQuantity(),
		StockBoughtQuantity:    quantity,
		CurrentWeightedAverage: s.repository.GetWeightAverage(),
		BoughtStockUnitCost:    cost,
	}

	average := calcWeightedAverage(params)

	s.repository.SetWeightAvarage(average)
	s.repository.SetStockQuantity(s.repository.GetStockQuantity() + quantity)

	return nil
}

func NewUseCaseService() *UseCaseService {
	return &UseCaseService{}
}
