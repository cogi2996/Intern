package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) CreateProject(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateProjectRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.Project{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Name:       req.Name,
		Code:       req.Code,
		ManagerID:  req.ManagerID,
		Address:    req.Address,
		RegionID:   req.RegionID,
		CategoryID: req.CategoryID,
	}
	data.StartDate, _ = services.StringToTimestamp(req.StartDate)
	data.EndDate, _ = services.StringToTimestamp(req.EndDate)

	id, err := h.Project.Create(ctx, data)
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
