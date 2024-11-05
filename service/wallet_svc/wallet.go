package wallet_svc

import "mlm/entity"

type Repository interface {
	Balance(userID uint, wallet string) (entity.Wallet, error)
	Create(userID uint, wallet string) (entity.Wallet, error)
}

type Service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return Service{
		repository: repo,
	}
}

func (s Service) Create(userID uint, wallet string) {

}

func (s Service) Balance(userID uint, wallet string) {

}
