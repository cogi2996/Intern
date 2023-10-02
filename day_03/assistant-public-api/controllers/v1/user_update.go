package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateUserRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}
	err = h.UserService.Update(ctx, &entities.User{
		ID:        id,
		Code:      req.Code,
		Name:      req.Name,
		Phone:     req.Phone,
		Email:     req.Email,
		Address:   req.Address,
		ManagerID: req.ManagerID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}
