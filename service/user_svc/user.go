package user_svc

import (
	"mlm/dto"
	"mlm/entity"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Delete(id uint) (bool, error)
}

type Service struct {
	repository UserRepository
}

func NewUserService(repo UserRepository) Service {
	return Service{
		repository: repo,
	}
}

func (s Service) Create(request dto.UserCreateRequest) (dto.UserCreateResponse, error) {

	user := entity.User{
		ID:          0,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		Role:        "marketer",
	}

	userCreated, err := s.repository.Create(user)
	if err != nil {
		return dto.UserCreateResponse{}, err
	}

	return dto.UserCreateResponse{
		ID: userCreated.ID,
	}, nil

}

func (s Service) Rollback(item uint) (bool, error) {

	b, err := s.repository.Delete(item)
	if err != nil {
		return false, err
	}
	return b, nil
}
