package cli

import "testing"

type MockStockUseCase struct {
	NCalledBuy      int
	NCalledSell     int
	BuyErrorResult  error
	SellErrorResult error
}

func (msu *MockStockUseCase) Buy(quantity, cost int) error {
	msu.NCalledBuy += 1
	return msu.BuyErrorResult
}

func (msu *MockStockUseCase) Sell(quantity, cost int) (int, error) {
	msu.NCalledSell += 1
	return 0, msu.SellErrorResult
}

func TestCLIRun(t *testing.T) {
	suite := []struct {
		Payload             RequestPayload
		ExpectedNCalledBuy  int
		ExpectedNCalledSell int
	}{
		{RequestPayload{Operation: "sell"}, 0, 1},
		{RequestPayload{Operation: "buy"}, 1, 0},
	}

	for _, s := range suite {
		m := &MockStockUseCase{}
		h := &Interface{usecase: m}

		if _, err := h.Run([]RequestPayload{s.Payload}); err != nil {
			t.Error(err)
			return
		}

		if got, expected := m.NCalledBuy, s.ExpectedNCalledBuy; got != expected {
			t.Errorf("unexpected N called buy, got: %v, expected: %v", got, expected)
			return
		}

		if got, expected := m.NCalledSell, s.ExpectedNCalledSell; got != expected {
			t.Errorf("unexpected N called sell, got: %v, expected: %v", got, expected)
			return
		}
	}
}
