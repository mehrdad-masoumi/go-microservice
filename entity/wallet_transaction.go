package entity

type WalletTransaction struct {
	ID       uint
	UserID   uint
	Title    string
	Creditor uint
	Debtor   uint
	Type     string
	// TODO - ADD FOR
}
