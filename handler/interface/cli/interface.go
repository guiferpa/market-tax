package cli

import (
	"github/guiferpa/market-tax/domain/stock"
)

type RequestPayload struct {
	Operation string  `json:"operation"`
	UnitCost  float64 `json:"unit-cost"`
	Quantity  int     `json:"quantity"`
}

type ResponsePayload struct {
	Tax float64 `json:"tax"`
}

type Interface struct {
	usecase stock.UserCase
}

func (cli *Interface) Run(payload []RequestPayload) ([]ResponsePayload, error) {
	resp := make([]ResponsePayload, 0)

	for _, data := range payload {
		if data.Operation == "buy" {
			if err := cli.usecase.Buy(data.Quantity, int(data.UnitCost*100)); err != nil {
				return nil, err
			}

			resp = append(resp, ResponsePayload{Tax: .0})
			continue
		}

		if data.Operation == "sell" {
			tax, err := cli.usecase.Sell(data.Quantity, int(data.UnitCost*100))
			if err != nil {
				return nil, err
			}

			resp = append(resp, ResponsePayload{Tax: float64(tax / 100)})
		}
	}

	return resp, nil
}

func NewInterface(usecase stock.UserCase) *Interface {
	return &Interface{usecase}
}
