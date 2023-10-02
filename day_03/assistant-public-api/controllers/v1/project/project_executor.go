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

func (h Handler) AddProjectExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.AddProjectExecutorRequest{}

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

	err = h.Project.AddExecutor(ctx, &entities.ProjectExecutor{
		Base: entities.Base{
			CreatedBy: userID,
		},
		ProjectID:  projectID,
		ExecutorID: req.ExecutorID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}

func (h Handler) ListProjectExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ListExecutor(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectExecutorResponse{}
	for _, val := range data {
		resp.Executors = append(resp.Executors, &resources.Executor{
			ID:      val.GetExecutor().GetID(),
			Name:    val.GetExecutor().GetName(),
			Phone:   val.GetExecutor().GetPhone(),
			Email:   val.GetExecutor().GetEmail(),
			Address: val.GetExecutor().GetAddress(),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}

func (h Handler) RemoveProjectExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	executorID, err := strconv.ParseInt(c.Param("executor_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}
	// userID, _, _ := middlewares.ParseToken(c)

	err = h.Project.DeleteExecutor(ctx, &entities.ProjectExecutor{
		ProjectID:  projectID,
		ExecutorID: executorID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}
