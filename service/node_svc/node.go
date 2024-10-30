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
	Delete(id uint) (bool, error)
}

type NodeService struct {
	repository NodeRepository
}

func NewNodeService(repo NodeRepository) NodeService {
	return NodeService{
		repository: repo,
	}
}

func (s NodeService) Create(request dto.NodeCreateRequest) (dto.NodeCreateResponse, error) {

	const target = "node_srv.Register"

	parent, err := s.repository.FindNodeByReferral(request.Referral)
	if err != nil {
		return dto.NodeCreateResponse{}, err
	}

	node := entity.Node{
		ID:          request.UserID,
		ParentId:    parent.ID,
		Ancestry:    makeAncestry(parent.Ancestry, request.UserID),
		Line:        makeLine(request.Referral),
		LftReferral: makeReferral("L", request.UserID),
		RgtReferral: makeReferral("R", request.UserID),
	}

	nodeCreated, err := s.repository.Create(node)
	if err != nil {
		return dto.NodeCreateResponse{}, err
	}

	// return created user
	return dto.NodeCreateResponse{
		ID: nodeCreated.ID,
	}, nil
}

func (s NodeService) Rollback(item uint) (any, error) {

	b, err := s.repository.Delete(item)
	if err != nil {
		return false, err
	}
	return b, nil
}

func makeAncestry(ancestry string, userID uint) string {
	return ancestry + strconv.Itoa(int(userID)) + "/"
}

func makeLine(referral string) string {
	return referral[:1]
}

func makeReferral(line string, userId uint) string {
	return line + strconv.Itoa(int(userId)) + "@" + strconv.Itoa(utils.RandRange(1000, 9999))
}
