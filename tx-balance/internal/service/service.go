package service

type BalanceService struct {
	repo Balance
}

func NewBalanceService(repo Balance) *BalanceService {
	return &BalanceService{
		repo: repo,
	}
}
