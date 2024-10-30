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

	nodeRepo := &mocks.NodeRepository{}

	nodeRequest := dto.NodeCreateRequest{
		UserID:   10,
		Referral: "L100@12135",
		WalletId: "TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE",
	}

	expectedNode := entity.Node{
		ID:          10,
		ParentId:    11, // TODO - node parent
		Ancestry:    "",
		Line:        "",
		LftReferral: "",
		RgtReferral: "",
	}

	nodeRepo.
		On("FindNodeByReferral", nodeRequest.Referral).
		Return(expectedNode, nil).
		Once()

	nodeRepo.
		On("Create", mock.MatchedBy(func(node entity.Node) bool {
			if node.ID == expectedNode.ID {
				return true
			}
			return false
		})).
		Return(expectedNode, nil).
		Once()

	service := NewNodeService(nodeRepo)
	response, err := service.Create(nodeRequest)

	assert.NoError(t, err)
	assert.Equal(t, dto.NodeCreateResponse{
		ID: expectedNode.ID,
	}, response)

	nodeRepo.AssertExpectations(t)

}

func TestNodeService_Register_failure(t *testing.T) {

	nodeRepo := &mocks.NodeRepository{}

	nodeRequest := dto.NodeCreateRequest{
		UserID:   10,
		Referral: "L100@12135",
		WalletId: "TZJjSqfXwTP5YgjdhLdRV1zqpDwDsgdZZE",
	}

	expectedError := errors.New("database error")

	service := NewNodeService(nodeRepo)
	_, err := service.Create(nodeRequest)

	assert.Error(t, err)
	assert.Equal(t, expectedError, err)

	nodeRepo.AssertExpectations(t)

}
