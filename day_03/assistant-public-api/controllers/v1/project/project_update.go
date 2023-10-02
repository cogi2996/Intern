package project

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) UpdateProject(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateProjectRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data := &entities.Project{
		ID:         projectID,
		Code:       req.Code,
		ManagerID:  req.ManagerID,
		Name:       req.Name,
		Address:    req.Address,
		RegionID:   req.RegionID,
		CategoryID: req.CategoryID,
	}
	data.StartDate, _ = services.StringToTimestamp(req.StartDate)
	data.EndDate, _ = services.StringToTimestamp(req.EndDate)

	err = h.Project.Update(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}
