package project

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/middlewares"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) CreateProjectArea(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateProjectAreaRequest{}

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.ProjectArea{
		Base: entities.Base{
			CreatedBy: userID,
		},
		ProjectID: projectID,
		Name:      req.Name,
	}

	err = h.Project.CreateArea(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.ProjectPhase{
			ID: data.ID,
		},
	})
}

func (h Handler) ReadProjectArea(c *gin.Context) {
	ctx := c.Request.Context()

	/*
		projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
			return
		}
	*/

	areaID, err := strconv.ParseInt(c.Param("area_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ReadArea(ctx, areaID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.ProjectArea{
			ID:   areaID,
			Name: data.Name,
			Project: &resources.Project{
				ID:   data.GetProject().GetID(),
				Name: data.GetProject().GetName(),
			},
		},
	})
}

func (h Handler) UpdateProjectArea(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateProjectAreaRequest{}

	/*
		projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
			return
		}
	*/

	areaID, err := strconv.ParseInt(c.Param("area_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.ProjectArea{
		ID: areaID,
		Base: entities.Base{
			UpdatedBy: userID,
		},
		Name: req.Name,
	}

	err = h.Project.UpdateArea(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}

func (h Handler) ListProjectArea(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ListArea(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectAreaResponse{}
	for _, val := range data {
		resp.Areas = append(resp.Areas, &resources.ProjectArea{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}

func (h Handler) RemoveProjectArea(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	areaID, err := strconv.ParseInt(c.Param("area_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, _, _ := middlewares.ParseToken(c)

	err = h.Project.DeleteArea(ctx, &entities.ProjectArea{
		ID:        areaID,
		ProjectID: projectID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}
