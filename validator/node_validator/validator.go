package node_validator

import (
	"mlm/dto"
	"mlm/entity"
)

const (
	referralRegex = "\\w{1}\\d+@\\d{4}$"
)

type Repository interface {
	FindNodeByReferral(referral string) (entity.Node, error)
}

type Validator struct {
	repo Repository
}

func NewNodeValidator(repo Repository) Validator {
	return Validator{repo: repo}
}

func (v Validator) ValidateNodeRequest(req dto.NodeCreateRequest) (map[string]string, error) {

	return nil, nil
}
