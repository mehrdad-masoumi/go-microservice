package dto

type LocalTransaction func() error
type CompensatingAction func() error

type SagaStep struct {
	Transaction LocalTransaction
	Compensate  CompensatingAction
}
