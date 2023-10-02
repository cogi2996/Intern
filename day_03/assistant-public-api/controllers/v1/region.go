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

func (h Handler) CreateRegion(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateRegionRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	userID, _, _ := middlewares.ParseToken(c)

	data := &entities.Region{
		Base: entities.Base{
			CreatedBy: userID,
		},
		Name: req.Name,
	}
	id, err := h.RegionService.Create(ctx, data)
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

func (h Handler) ReadRegion(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.RegionService.Read(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: &resources.User{
			ID:   data.ID,
			Name: data.Name,
		},
	})
}

func (h Handler) UpdateRegion(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.UpdateRegionRequest{}

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
	err = h.RegionService.Update(ctx, &entities.Region{
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

func (h Handler) DeleteRegion(c *gin.Context) {
	ctx := c.Request.Context()

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	err = h.RegionService.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resources.Empty{},
	})
}

func (h Handler) ListRegion(c *gin.Context) {
	ctx := c.Request.Context()

	regions, err := h.RegionService.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListRegionResponse{}
	for _, val := range regions {
		resp.Regions = append(resp.Regions, &resources.Region{
			ID:   val.ID,
			Name: val.Name,
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
