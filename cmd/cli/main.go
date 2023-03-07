package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github/guiferpa/market-tax/domain/stock"
	"github/guiferpa/market-tax/handler/interface/cli"
	"github/guiferpa/market-tax/infra/storage/memory"
)

func main() {
	input := make([][]cli.RequestPayload, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		var line []cli.RequestPayload
		if err := json.Unmarshal(scanner.Bytes(), &line); err != nil {
			log.Panicln(err)
			return
		}

		input = append(input, line)
	}

	for _, payload := range input {
		storage := memory.NewMemoryStorage()
		service := stock.NewUseCaseService(storage)
		cli := cli.NewInterface(service)

		resp, err := cli.Run(payload)
		if err != nil {
			log.Println(err)
		}

		if err = json.NewEncoder(os.Stdout).Encode(resp); err != nil {
			log.Println(err)
		}
	}
}
