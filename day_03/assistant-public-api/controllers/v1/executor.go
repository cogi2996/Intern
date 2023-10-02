package v1

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

func (h Handler) CreateExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateExecutorRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.Executor{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Code:          req.Code,
		Name:          req.Name,
		Address:       req.Address,
		RepresentedBy: req.RepresenterID,
		Phone:         req.Phone,
		Email:         req.Email,
	}
	id, err := h.ExecutorService.Create(ctx, data)
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

func (h Handler) ReadExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.ExecutorService.Read(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Executor{
			ID:      data.ID,
			Code:    data.Code,
			Name:    data.Name,
			Address: data.Address,
			Phone:   data.Phone,
			Email:   data.Email,
			Representer: &resources.User{
				ID:   data.GetRepresenter().GetID(),
				Name: data.GetRepresenter().GetName(),
			},
		},
	})
}

func (h Handler) UpdateExecutor(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateExecutorRequest{}

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
	err = h.ExecutorService.Update(ctx, &entities.Executor{
		ID:            id,
		Code:          req.Code,
		Name:          req.Name,
		Address:       req.Address,
		RepresentedBy: req.RepresenterID,
		Phone:         req.Phone,
		Email:         req.Email,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}

func (h Handler) DeleteExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = h.ExecutorService.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}

func (h Handler) ListExecutor(c *gin.Context) {
	ctx := c.Request.Context()

	executors, err := h.ExecutorService.GetList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListExecutorResponse{}
	for _, val := range executors {
		members := make([]*resources.User, 0)
		members = append(members, &resources.User{
			ID:   val.GetRepresenter().GetID(),
			Name: val.GetRepresenter().GetName(),
		})
		for _, member := range val.GetRepresenter().GetMembers() {
			members = append(members, &resources.User{
				ID:   member.GetID(),
				Name: member.GetName(),
			})
		}

		resp.Executors = append(resp.Executors, &resources.Executor{
			ID:   val.ID,
			Code: val.Code,
			Name: val.Name,
			Representer: &resources.User{
				ID:   val.GetRepresenter().GetID(),
				Name: val.GetRepresenter().GetName(),
			},
			Address: val.Address,
			Phone:   val.Phone,
			Members: members,
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
