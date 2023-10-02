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
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) CreateProjectPhase(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateProjectPhaseRequest{}

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

	data := &entities.ProjectPhase{
		Base: entities.Base{
			CreatedBy: userID,
		},
		ProjectID: projectID,
		Name:      req.Name,
	}
	data.StartDate, _ = services.StringToTimestamp(req.StartDate)
	data.EndDate, _ = services.StringToTimestamp(req.EndDate)

	err = h.Project.CreatePhase(ctx, data)
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

func (h Handler) ReadProjectPhase(c *gin.Context) {
	ctx := c.Request.Context()

	/*
		projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
			return
		}
	*/

	phaseID, err := strconv.ParseInt(c.Param("phase_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ReadPhase(ctx, phaseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.ProjectPhase{
			ID:        phaseID,
			Name:      data.Name,
			StartDate: services.FormatTimestamp(data.GetProject().GetStartDate()),
			EndDate:   services.FormatTimestamp(data.GetProject().GetEndDate()),
			Project: &resources.Project{
				ID:   data.GetProject().GetID(),
				Name: data.GetProject().GetName(),
			},
		},
	})
}

func (h Handler) UpdateProjectPhase(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateProjectPhaseRequest{}

	/*
		projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
			return
		}
	*/

	phaseID, err := strconv.ParseInt(c.Param("phase_id"), 10, 64)
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

	data := &entities.ProjectPhase{
		ID: phaseID,
		Base: entities.Base{
			UpdatedBy: userID,
		},
		Name: req.Name,
	}
	data.StartDate, _ = services.StringToTimestamp(req.StartDate)
	data.EndDate, _ = services.StringToTimestamp(req.EndDate)

	err = h.Project.UpdatePhase(ctx, data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}

func (h Handler) ListProjectPhase(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ListPhase(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectPhaseResponse{}
	for _, val := range data {
		resp.Phases = append(resp.Phases, &resources.ProjectPhase{
			ID:        val.ID,
			Name:      val.Name,
			StartDate: services.FormatTimestamp(val.StartDate),
			EndDate:   services.FormatTimestamp(val.EndDate),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}

func (h Handler) RemoveProjectPhase(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	phaseID, err := strconv.ParseInt(c.Param("phase_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, _, _ := middlewares.ParseToken(c)

	err = h.Project.DeletePhase(ctx, &entities.ProjectPhase{
		ID:        phaseID,
		ProjectID: projectID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}
