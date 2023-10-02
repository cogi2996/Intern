package project

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/pkg/util"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ReadProject(c *gin.Context) {
	ctx := c.Request.Context()

	projectID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("invalid parameter"))
		return
	}

	data, err := h.Project.Read(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.Project{
		ID:        data.ID,
		Code:      data.Code,
		Name:      data.Name,
		Address:   data.Address,
		StartDate: util.TimestampToDateFormat(data.StartDate),
		EndDate:   util.TimestampToDateFormat(data.EndDate),
		Category: &resources.ProjectCategory{
			ID:   data.GetCategory().GetID(),
			Name: data.GetCategory().GetName(),
		},
		Region: &resources.Region{
			ID:   data.GetRegion().GetID(),
			Name: data.GetRegion().GetName(),
		},
		Manager: &resources.User{
			ID:   data.GetManager().GetID(),
			Name: data.GetManager().GetName(),
		},
		Areas:  make([]*resources.ProjectArea, 0),
		Phases: make([]*resources.ProjectPhase, 0),
	}

	for _, area := range data.Areas {
		resp.Areas = append(resp.Areas, &resources.ProjectArea{
			ID:   area.ID,
			Name: area.Name,
		})
	}

	for _, phase := range data.Phases {
		resp.Phases = append(resp.Phases, &resources.ProjectPhase{
			ID:        phase.ID,
			Name:      phase.Name,
			StartDate: services.FormatTimestamp(phase.StartDate),
			EndDate:   services.FormatTimestamp(phase.EndDate),
		})
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
