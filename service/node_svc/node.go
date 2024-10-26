package node_svc

import (
	"mlm/dto"
	"mlm/entity"
	"mlm/pkg/utils"
	"strconv"
)

type NodeRepository interface {
	Create(node entity.Node) (entity.Node, error)
	FindNodeByReferral(referral string) (entity.Node, error)
}

type UserService interface {
	Create(request dto.UserCreateRequest) (dto.UserCreateResponse, error)
}

type NodeService struct {
	repository NodeRepository
	userSvc    UserService
}

func NewNodeService(repo NodeRepository, userSvc UserService) NodeService {
	return NodeService{
		repository: repo,
		userSvc:    userSvc,
	}
}

func (s NodeService) Register(request dto.NodeCreateRequest) (dto.NodeCreateResponse, error) {

	UserCreateRequest := dto.UserCreateRequest{
		Username:        request.Username,
		Email:           request.Email,
		PhoneNumber:     request.PhoneNumber,
		Password:        request.Password,
		ConfirmPassword: request.ConfirmPassword,
	}
	// create new user in storage
	userCreated, err := s.userSvc.Create(UserCreateRequest)
	if err != nil {
		return dto.NodeCreateResponse{}, err
	}

	parent, err := s.repository.FindNodeByReferral(request.Referral)
	if err != nil {
		return dto.NodeCreateResponse{}, err
	}

	node := entity.Node{
		Id:          userCreated.ID,
		ParentId:    parent.Id,
		Ancestry:    makeAncestry(parent.Ancestry, userCreated.ID),
		Line:        makeLine(request.Referral),
		LftReferral: makeReferral("L", userCreated.ID),
		RgtReferral: makeReferral("R", userCreated.ID),
	}

	nodeCreated, err := s.repository.Create(node)
	if err != nil {
		return dto.NodeCreateResponse{}, err
	}

	// return created user
	return dto.NodeCreateResponse{
		NodeId: nodeCreated.Id,
	}, nil
}

func makeAncestry(ancestry string, userID uint) string {
	return ancestry + "/" + strconv.Itoa(int(userID))
}

func makeLine(referral string) entity.Line {
	return entity.Line(referral[:1])
}
func makeReferral(line string, userId uint) string {
	return line + strconv.Itoa(int(userId)) + "@" + strconv.Itoa(utils.RandRange(1000, 9999))
}
