package stock

import "testing"

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
