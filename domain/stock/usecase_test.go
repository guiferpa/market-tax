package stock

import (
	"errors"
	"testing"
)

func TestCalcTax(t *testing.T) {
	suite := []struct {
		Gain     int
		Expected int
	}{
		{100_000_00, 20_000_00},
		{0, 0},
		{-100, -20},
	}

	for _, s := range suite {
		if got, expected := calcTax(s.Gain), s.Expected; got != expected {
			t.Errorf("unexpected result for tax, got: %d, expected: %d", got, expected)
			return
		}
	}
}

func TestCalcWeightedAverage(t *testing.T) {
	suite := []struct {
		CurrentStockQuantity   int
		CurrentWeightedAverage int
		StockBoughtQuantity    int
		BoughtStockUnitCost    int
		Expected               int
	}{
		{5, 2000, 5, 1000, 1500},
		{0, 0, 5, 200, 200},
	}

	for _, s := range suite {
		params := calcWeightedAverageParams{
			CurrentStockQuantity:   s.CurrentStockQuantity,
			StockBoughtQuantity:    s.StockBoughtQuantity,
			CurrentWeightedAverage: s.CurrentWeightedAverage,
			BoughtStockUnitCost:    s.BoughtStockUnitCost,
		}
		if got, expected := calcWeightedAverage(params), s.Expected; got != expected {
			t.Errorf("unexpected result for weighted average, got: %d, expected: %d", got, expected)
			return
		}
	}
}

type StockRepositoryMock struct {
	NCalledSetWeightAverage int
	NCalledGetWeightAverage int
	NCalledSetStockQuantity int
	NCalledGetStockQuantity int
	StockQuantityResult     int
}

func (m *StockRepositoryMock) SetWeightAvarage(value int) {
	m.NCalledSetWeightAverage += 1
}

func (m *StockRepositoryMock) GetWeightAverage() int {
	m.NCalledGetWeightAverage += 1
	return 0
}

func (m *StockRepositoryMock) SetFinancialLoss(value int) {

}

func (m *StockRepositoryMock) GetFinancialLoss() int {
	return 0
}

func (m *StockRepositoryMock) SetStockQuantity(value int) {
	m.StockQuantityResult = value
	m.NCalledSetStockQuantity += 1
}

func (m *StockRepositoryMock) GetStockQuantity() int {
	m.NCalledGetStockQuantity += 1
	return m.StockQuantityResult
}

func TestBuyWithInvalidQuantity(t *testing.T) {
	suite := []struct {
		Quantity int
	}{
		{0},
		{-1000},
	}

	for _, s := range suite {
		mock := &StockRepositoryMock{}
		svc := &UseCaseService{repository: mock}

		if got, expected := svc.Buy(s.Quantity, 1), ErrStockBuyWithInvalidQuantity; !errors.Is(got, expected) {
			t.Errorf("unexpected error, got: %v, expected %v", got, expected)
		}
	}
}

func TestBuyWithInvalidCost(t *testing.T) {
	suite := []struct {
		Cost int
	}{
		{0},
		{-1000},
	}

	for _, s := range suite {
		mock := &StockRepositoryMock{}
		svc := &UseCaseService{repository: mock}

		if got, expected := svc.Buy(1, s.Cost), ErrStockBuyWithInvalidCost; !errors.Is(got, expected) {
			t.Errorf("unexpected error, got: %v, expected %v", got, expected)
		}
	}
}

func TestBuy(t *testing.T) {
	suite := []struct {
		MockCurrentStockQuantity        int
		Quantity                        int
		Cost                            int
		ExpectedTotalStockQuantity      int
		ExpectedNCalledSetWeightAverage int
		ExpectedNCalledGetWeightAverage int
		ExpectedNCalledSetStockQuantity int
		ExpectedNCalledGetStockQuantity int
	}{
		{10, 5, 20, 15, 1, 1, 1, 2},
	}

	for _, s := range suite {
		mock := &StockRepositoryMock{
			StockQuantityResult: s.MockCurrentStockQuantity,
		}
		svc := &UseCaseService{repository: mock}

		if err := svc.Buy(s.Quantity, s.Cost); err != nil {
			t.Error(err)
			return
		}

		if got, expected := mock.NCalledSetWeightAverage, s.ExpectedNCalledSetWeightAverage; got != expected {
			t.Errorf("unexpected N called SetWeightAverage, got; %v, expected: %v", got, expected)
			return
		}

		if got, expected := mock.NCalledGetWeightAverage, s.ExpectedNCalledGetWeightAverage; got != expected {
			t.Errorf("unexpected N called GetWeightAverage, got; %v, expected: %v", got, expected)
			return
		}

		if got, expected := mock.NCalledSetStockQuantity, s.ExpectedNCalledSetStockQuantity; got != expected {
			t.Errorf("unexpected N called SetStockQuantity, got; %v, expected: %v", got, expected)
			return
		}

		if got, expected := mock.NCalledGetStockQuantity, s.ExpectedNCalledGetStockQuantity; got != expected {
			t.Errorf("unexpected N called GetStockQuantity, got; %v, expected: %v", got, expected)
			return
		}

		if got, expected := mock.StockQuantityResult, s.ExpectedTotalStockQuantity; got != expected {
			t.Errorf("unexpected for total of stock quantity, got; %v, expected: %v", got, expected)
			return
		}
	}
}
