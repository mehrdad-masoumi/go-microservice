package node_svc

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mlm/dto"
	"mlm/entity"
	"mlm/mocks"
	"regexp"
	"strconv"
	"testing"
)

func TestMakeAncestry(t *testing.T) {
	ancestry := "/1/2/"
	userID := uint(3)

	expected := ancestry + "/" + strconv.Itoa(int(userID))
	result := makeAncestry(ancestry, userID)

	assert.Equal(t, expected, result, "makeAncestry should concatenate ancestry and userID correctly")
}

func TestMakeLine(t *testing.T) {

	referral := "L100@1234"

	expected := entity.Line(referral[:1])

	result := makeLine(referral)

	assert.Equal(t, expected, result)

}

func TestMakeReferral(t *testing.T) {
	line := "L"
	userId := uint(123)

	// Call makeReferral function
	referral := makeReferral(line, userId)

	// Check that the referral starts with the line and userId
	expectedPrefix := line + strconv.Itoa(int(userId)) + "@"

	// Use regex to verify that the referral follows the format L123@<random_number>
	assert.True(t, regexp.MustCompile(`^`+expectedPrefix+`\d{4}$`).MatchString(referral), "Referral should follow the format L123@<4-digit_number>")

	// Optional: Parse the random number part and ensure it's within the expected range (1000-9999)
	randomPart := referral[len(expectedPrefix):]
	randomNumber, err := strconv.Atoi(randomPart)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, randomNumber, 1000)
	assert.LessOrEqual(t, randomNumber, 9999)
}

func TestNodeService_Register_success(t *testing.T) {

	userSvc := &mocks.UserService{}
	nodeRepo := &mocks.NodeRepository{}

	nodeRequest := dto.NodeCreateRequest{
		Username:        "mehrdad",
		Email:           "mehrdad@gmail.com",
		PhoneNumber:     "09120246217",
		Referral:        "L100@12135",
		WalletId:        "TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE",
		Password:        "123456789",
		ConfirmPassword: "123456789",
	}

	userRequest := dto.UserCreateRequest{
		Username:        nodeRequest.Username,
		Email:           nodeRequest.Email,
		PhoneNumber:     nodeRequest.PhoneNumber,
		Password:        nodeRequest.Password,
		ConfirmPassword: nodeRequest.ConfirmPassword,
	}

	expectedUser := dto.UserCreateResponse{
		ID: 1,
	}

	expectedNode := entity.Node{
		Id:          expectedUser.ID,
		ParentId:    10, // TODO - node parent
		Ancestry:    "",
		Line:        "",
		LftReferral: "",
		RgtReferral: "",
	}

	userSvc.
		On("Create", mock.MatchedBy(func(req dto.UserCreateRequest) bool {
			if req.PhoneNumber == userRequest.PhoneNumber && req.Email == userRequest.Email {
				return true
			}
			return false
		})).
		Return(expectedUser, nil).
		Once()

	nodeRepo.
		On("FindNodeByReferral", nodeRequest.Referral).
		Return(expectedNode, nil).
		Once()

	nodeRepo.
		On("Create", mock.MatchedBy(func(node entity.Node) bool {
			if node.Id == expectedNode.Id {
				return true
			}
			return false
		})).
		Return(expectedNode, nil).
		Once()

	service := NewNodeService(nodeRepo, userSvc)
	response, err := service.Register(nodeRequest)

	assert.NoError(t, err)
	assert.Equal(t, dto.NodeCreateResponse{
		NodeId: expectedNode.Id,
	}, response)

	userSvc.AssertExpectations(t)
	nodeRepo.AssertExpectations(t)

}

func TestNodeService_Register_failure(t *testing.T) {

	userSvc := &mocks.UserService{}
	nodeRepo := &mocks.NodeRepository{}

	nodeRequest := dto.NodeCreateRequest{
		Username:        "mehrdad",
		Email:           "mehrdad@gmail.com",
		PhoneNumber:     "09120246217",
		Referral:        "L100@12135",
		WalletId:        "TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE",
		Password:        "123456789",
		ConfirmPassword: "123456789",
	}

	expectedError := errors.New("database error")

	userSvc.
		On("Create", mock.Anything).
		Return(dto.UserCreateResponse{}, expectedError).
		Once()

	service := NewNodeService(nodeRepo, userSvc)
	_, err := service.Register(nodeRequest)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	userSvc.AssertExpectations(t)
	nodeRepo.AssertExpectations(t)

}
