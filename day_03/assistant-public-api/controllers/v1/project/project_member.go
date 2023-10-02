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

func (h Handler) AddProjectMember(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.AddProjectMemberRequest{}

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

	err = h.Project.AddMember(ctx, &entities.ProjectMember{
		Base: entities.Base{
			CreatedBy: userID,
		},
		ProjectID: projectID,
		MemberID:  req.UserID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}

func (h Handler) ListProjectMember(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	data, err := h.Project.ListMember(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectMemberResponse{}
	for _, val := range data {
		resp.Users = append(resp.Users, &resources.User{
			ID:      val.GetMember().GetID(),
			Name:    val.GetMember().GetName(),
			Phone:   val.GetMember().GetPhone(),
			Email:   val.GetMember().GetEmail(),
			Address: val.GetMember().GetAddress(),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}

func (h Handler) RemoveProjectMember(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	memberID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	// userID, _, _ := middlewares.ParseToken(c)

	err = h.Project.DeleteMember(ctx, &entities.ProjectMember{
		ProjectID: projectID,
		MemberID:  memberID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{})
}
