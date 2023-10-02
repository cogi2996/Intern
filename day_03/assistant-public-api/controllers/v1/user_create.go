package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateUserRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.User{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Code:      req.Code,
		Name:      req.Name,
		Phone:     req.Phone,
		Email:     req.Email,
		Address:   req.Address,
		ManagerID: req.ManagerID,
	}
	id, err := h.UserService.Create(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.CreateResponse{
			ID:   id,
			UUID: data.UUID,
		},
	})
}
