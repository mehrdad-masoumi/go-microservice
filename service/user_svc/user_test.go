package user_svc_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mlm/dto"
	"mlm/entity"
	"mlm/mocks"
	"mlm/service/user_svc"
	"testing"
)

func TestUserService_Register_Success(t *testing.T) {
	// Create a mock repository
	userRepo := &mocks.UserRepository{}

	// Define the input request
	request := dto.UserCreateRequest{
		Username:        "mehrdad",
		Email:           "mehrdad@gmail.com",
		PhoneNumber:     "0912024617",
		Password:        "123456789",
		ConfirmPassword: "123456789",
	}

	// Simulate a user that will be returned by the repository when Register is called
	expectedUser := entity.User{
		ID:          1,
		Email:       request.Email,
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		Role:        "marketer",
	}

	// Set up the mock expectation
	userRepo.
		On("Create", mock.MatchedBy(func(user entity.User) bool {
			return user.PhoneNumber == request.PhoneNumber && user.Email == request.Email
		})).
		Return(expectedUser, nil).
		Once()

	// Create the service with the mock repository
	service := user_svc.NewUserService(userRepo)

	// Call the Register method
	response, err := service.Create(request)

	// Assert no errors and correct response
	assert.NoError(t, err)
	assert.Equal(t, dto.UserCreateResponse{
		ID: response.ID,
	}, response)

	// Assert that all mock expectations were met
	userRepo.AssertExpectations(t)
}

func TestUserService_Register_Failure(t *testing.T) {
	// Create a mock repository
	userRepo := &mocks.UserRepository{}

	// Define the input request
	request := dto.UserCreateRequest{
		PhoneNumber: "0912024617",
	}

	// Simulate an error when the repository's Register method is called
	expectedError := errors.New("database error")

	// Set up the mock expectation
	userRepo.On("Create", mock.Anything).Return(entity.User{}, expectedError).Once()

	// Create the service with the mock repository
	service := user_svc.NewUserService(userRepo)

	// Call the Register method
	_, err := service.Create(request)

	// Assert that the error is returned
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	// Assert that all mock expectations were met
	userRepo.AssertExpectations(t)
}
