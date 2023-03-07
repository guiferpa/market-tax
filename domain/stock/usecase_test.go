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
