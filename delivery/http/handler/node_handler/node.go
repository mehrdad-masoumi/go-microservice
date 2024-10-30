package node_handler

import (
	"github.com/labstack/echo/v4"
	"mlm/dto"
	"mlm/pkg/http_msg"
	"mlm/pkg/saga"
	"mlm/service/node_svc"
	"mlm/service/user_svc"
	"mlm/validator/node_validator"
	"net/http"
)

type Workflow interface {
	AddStep(name string, step dto.SagaStep)
	Execute() error
}

type NodeHandler struct {
	nodeSvc       node_svc.NodeService
	nodeValidator node_validator.Validator
	userSvc       user_svc.Service
}

func NewNodeHandler(nodeSvc node_svc.NodeService, nodeValidator node_validator.Validator, userSvc user_svc.Service) NodeHandler {
	return NodeHandler{
		nodeSvc:       nodeSvc,
		nodeValidator: nodeValidator,
		userSvc:       userSvc,
	}
}

// @Summary		node register
// @Accept		json
// @Produce		json
// @Param payload body dto.RegisterRequest true "payload"
// @Router		/auth/register [post]
func (h NodeHandler) register(c echo.Context) error {
	item := dto.RegisterRequest{}
	c.Bind(&item)

	fieldErrors, err := h.nodeValidator.ValidateNodeCreateRequest(item)
	if err != nil {

		msg, code := http_msg.Error(err)

		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	workflow := saga.New()

	userRequest := dto.UserCreateRequest{
		Username:        item.Username,
		Email:           item.Email,
		PhoneNumber:     item.PhoneNumber,
		Password:        item.Password,
		ConfirmPassword: item.ConfirmPassword,
	}

	var userID uint

	workflow.AddStep("1", dto.SagaStep{
		Transaction: func() error {
			resp, err := h.userSvc.Create(userRequest)
			if err != nil {
				// TODO - HANDLE ERROR
				return err
			}
			userID = resp.ID
			return nil
		},
		Compensate: func() error {
			_, err := h.userSvc.Rollback(userID)
			if err != nil {
				// TODO - HANDLE ERROR
				return err
			}
			return nil
		},
	})

	workflow.AddStep("2", dto.SagaStep{
		Transaction: func() error {

			nodeRequest := dto.NodeCreateRequest{
				UserID:   userID,
				Referral: item.Referral,
				WalletId: item.WalletId,
			}

			resp, err := h.nodeSvc.Create(nodeRequest)
			if err != nil {
				// TODO - HANDLE ERROR
				return err
			}
			userID = resp.ID
			return nil
		},
		Compensate: func() error {
			_, err := h.nodeSvc.Rollback(userID)
			if err != nil {
				// TODO - HANDLE ERROR
			}
			return nil
		},
	})

	err = workflow.Execute()

	if err != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, true)
}

func (h NodeHandler) SetRouter(e *echo.Echo) {

	users := e.Group("/auth")

	users.POST("/register", h.register)

}
