package stock

type UseCaseService struct{}

func calcTax(gain int) int {
	return (gain * 20) / 100
}

func NewUseCaseService() *UseCaseService {
	return &UseCaseService{}
}
