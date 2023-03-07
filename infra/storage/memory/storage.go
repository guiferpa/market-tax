package memory

import "github/guiferpa/market-tax/pkg/storage"

const (
	WEIGHT_AVERAGE_KEY = "WEIGHT_AVERAGE_KEY"
	FINANCIAL_LOSS_KEY = "FINANCIAL_LOSS_KEY"
	STOCK_QUANTITY_KEY = "STOCK_QUANTITY_KEY"
)

type MemoryStorage struct {
	storage *storage.Storage
}

func (ms *MemoryStorage) SetWeightAvarage(value int) {
	ms.storage.Set(WEIGHT_AVERAGE_KEY, value)
}

func (ms *MemoryStorage) GetWeightAverage() int {
	return ms.storage.Get(WEIGHT_AVERAGE_KEY)
}

func (ms *MemoryStorage) SetFinancialLoss(value int) {
	ms.storage.Set(FINANCIAL_LOSS_KEY, value)
}

func (ms *MemoryStorage) GetFinancialLoss() int {
	return ms.storage.Get(FINANCIAL_LOSS_KEY)
}

func (ms *MemoryStorage) SetStockQuantity(value int) {
	ms.storage.Set(STOCK_QUANTITY_KEY, value)
}

func (ms *MemoryStorage) GetStockQuantity() int {
	return ms.storage.Get(STOCK_QUANTITY_KEY)
}

func NewMemoryStorage() *MemoryStorage {
	s := storage.NewStorage()
	return &MemoryStorage{storage: s}
}
