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

func (h Handler) CreateProjectCategory(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateProjectCategoryRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.ProjectCategory{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Name: req.Name,
	}
	id, err := h.Category.Create(ctx, data)
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

func (h Handler) ReadProjectCategory(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Category.Read(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.ProjectCategory{
			ID:   data.ID,
			Name: data.Name,
		},
	})
}

func (h Handler) UpdateProjectCategory(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateProjectCategoryRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}
	err = h.Category.Update(ctx, &entities.ProjectCategory{
		ID:   id,
		Name: req.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.Empty{},
	})
}

func (h Handler) DeleteProjectCategory(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = h.Category.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}

func (h Handler) ListProjectCategory(c *gin.Context) {
	ctx := c.Request.Context()

	categories, err := h.Category.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectCategoryResponse{}
	for _, val := range categories {
		resp.Categories = append(resp.Categories, &resources.ProjectCategory{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
