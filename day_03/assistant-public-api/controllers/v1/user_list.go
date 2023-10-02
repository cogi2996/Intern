package v1

import (
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/middlewares"
)

func (h Handler) ListUser(c *gin.Context) {
	ctx := c.Request.Context()
	userID, _, _ := middlewares.ParseToken(c)

	users, err := h.UserService.List(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	resp := &resources.ListUserResponse{}
	for _, val := range users {
		if len(val.Executors) > 0 {
			continue
		}

		name := val.Name
		if val.ID == userID {
			name = fmt.Sprintf("%s (bản thân)", val.Name)
		}
		resp.Users = append(resp.Users, &resources.User{
			ID:      val.ID,
			Code:    val.Code,
			Name:    name,
			Phone:   val.Phone,
			Email:   val.Email,
			Address: val.Address,
			Manager: &resources.User{
				ID:   val.GetManager().GetID(),
				Code: val.GetManager().GetCode(),
				Name: val.GetManager().GetName(),
			},
		})
	}

	sort.Slice(resp.Users, func(i, j int) bool {
		if resp.Users[i].ID == userID {
			return true
		}

		if resp.Users[j].ID == userID {
			return false
		}

		s1 := strings.Split(resp.Users[i].Name, " ")
		s2 := strings.Split(resp.Users[i].Name, " ")
		if (len(s1) == 0) || (len(s2) == 0) {
			return strings.Compare(resp.Users[i].Name, resp.Users[j].Name) < 0
		}

		return strings.Compare(s1[len(s1)-1], s2[len(s2)-1]) < 0
	})

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
