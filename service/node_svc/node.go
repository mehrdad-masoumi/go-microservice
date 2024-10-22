package node_svc

import (
	"mlm/dto"
	"mlm/entity"
)

type NodeRepository interface {
	Create(node entity.Node) (entity.Node, error)
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

	// TODO- process node before insert, make Ancestry, Line, ....
	node := entity.Node{
		Id:          userCreated.ID,
		ParentId:    10,
		Ancestry:    "",
		Line:        "",
		LftReferral: "",
		RgtReferral: "",
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
