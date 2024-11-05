package user_svc

import (
	"mlm/dto"
	"mlm/entity"
	"mlm/pkg/richerror"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Delete(id uint) (bool, error)
}

type UserValidator interface {
	Validate(request dto.UserCreateRequest) (map[string]string, error)
}

type Service struct {
	repository UserRepository
	validator  UserValidator
}

func NewUserService(repo UserRepository, validator UserValidator) Service {
	return Service{
		repository: repo,
		validator:  validator,
	}
}

func (s Service) Create(request dto.UserCreateRequest) (dto.UserCreateResponse, error) {

	const op = "userService.Create"
	fieldErrors, err := s.validator.Validate(request)
	if err != nil {
		return dto.UserCreateResponse{FieldErrors: fieldErrors}, richerror.New(op).WithErr(err)
	}

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

	const op = "userService.Rollback"
	b, err := s.repository.Delete(item)
	if err != nil {
		return false, richerror.New(op).WithErr(err)
	}
	return b, nil
}
