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
	Tax   float64 `json:"tax"`
	Error string  `json:"error"`
}

type Interface struct {
	usecase stock.UserCase
}

func (cli *Interface) Run(payload []RequestPayload) []map[string]interface{} {
	errorCounter := 0
	resp := make([]map[string]interface{}, 0)

	for _, data := range payload {
		if errorCounter == 3 {
			resp = append(resp, map[string]interface{}{"error": "Your account is blocked"})
			return resp
		}

		if data.Operation == "buy" {
			if err := cli.usecase.Buy(data.Quantity, int(data.UnitCost*100)); err != nil {
				errorCounter += 1
				resp = append(resp, map[string]interface{}{"error": err.Error()})
				continue
			}

			errorCounter = 0
			resp = append(resp, map[string]interface{}{"tax": 0})
			continue
		}

		if data.Operation == "sell" {
			tax, err := cli.usecase.Sell(data.Quantity, int(data.UnitCost*100))
			if err != nil {
				errorCounter += 1
				resp = append(resp, map[string]interface{}{"error": err.Error()})
				continue
			}

			errorCounter = 0
			resp = append(resp, map[string]interface{}{"tax": float64(tax / 100)})
		}
	}

	return resp
}

func NewInterface(usecase stock.UserCase) *Interface {
	return &Interface{usecase}
}
