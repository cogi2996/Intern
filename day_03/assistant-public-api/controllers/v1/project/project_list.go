package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
	"github.com/ideal-forward/assistant-public-api/services"
)

func (h Handler) ListProject(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.CreateProjectRequest{}

	err := http_parser.BindAndValid(c, req, "MaxSize", "DateFormat")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// userID, organizationID, _ := middlewares.ParseToken(c)

	projects, err := h.Project.GetList(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListProjectResponse{}
	for _, val := range projects {
		element := &resources.Project{
			ID:        val.ID,
			Code:      val.Code,
			Name:      val.Name,
			Address:   val.Address,
			StartDate: services.FormatTimestamp(val.StartDate),
			EndDate:   services.FormatTimestamp(val.EndDate),
			Category: &resources.ProjectCategory{
				ID:   val.GetCategory().GetID(),
				Name: val.GetCategory().GetName(),
			},
			Region: &resources.Region{
				ID:   val.GetRegion().GetID(),
				Name: val.GetRegion().GetName(),
			},
			Manager: &resources.User{
				ID:   val.GetManager().GetID(),
				Name: val.GetManager().GetName(),
			},
		}

		element.Areas = make([]*resources.ProjectArea, 0)
		for _, area := range val.Areas {
			element.Areas = append(element.Areas, &resources.ProjectArea{
				ID:   area.ID,
				Name: area.Name,
			})
		}

		element.Phases = make([]*resources.ProjectPhase, 0)
		for _, phase := range val.Phases {
			element.Phases = append(element.Phases, &resources.ProjectPhase{
				ID:        phase.ID,
				Name:      phase.Name,
				StartDate: services.FormatTimestamp(phase.StartDate),
				EndDate:   services.FormatTimestamp(phase.EndDate),
			})
		}

		resp.Projects = append(resp.Projects, element)
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
