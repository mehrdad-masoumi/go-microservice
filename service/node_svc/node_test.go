package node_svc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"mlm/dto"
	"mlm/entity"
	"mlm/mocks"
	"testing"
)

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
