package wallet_transaction_svc

type Repository interface {
	Balance()
}

type Service struct {
	repository Repository
}

func (s Service) Creditor(userID uint, wallet string, amount float64) (bool, error) {
	return true, nil
}

func (s Service) Debtor(userID uint, wallet string, amount float64) (bool, error) {
	return true, nil
}

func (s Service) Locked(userID uint, wallet string, amount float64) (bool, error) {
	return true, nil
}
