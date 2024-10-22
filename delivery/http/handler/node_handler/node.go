package node_handler

import (
	"github.com/labstack/echo/v4"
	"mlm/dto"
	"mlm/service/node_svc"
	"net/http"
)

type NodeHandler struct {
	nodeSvc node_svc.NodeService
}

func NewNodeHandler(nodeSvc node_svc.NodeService) NodeHandler {
	return NodeHandler{
		nodeSvc: nodeSvc,
	}
}

// @Summary		node register
// @Accept		json
// @Produce		json
// @Param payload body dto.NodeCreateRequest true "payload"
// @Router		/auth/register [post]
func (h NodeHandler) register(c echo.Context) error {
	item := dto.NodeCreateRequest{}
	c.Bind(&item)
	return c.JSON(http.StatusOK, dto.UserInfo{
		Id:          1,
		Username:    "",
		Email:       item.Email,
		PhoneNumber: item.PhoneNumber,
	})
}

func (h NodeHandler) SetRouter(e *echo.Echo) {

	users := e.Group("/auth")

	users.POST("/register", h.register)

}
