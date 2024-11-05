package entity

type Wallet struct {
	UserID       uint
	Type         string
	Deposit      uint
	Withdraw     uint
	LockedAmount uint
	Balance      uint
}
