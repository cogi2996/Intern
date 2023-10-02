package v1

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) Authenticate(c *gin.Context) {
	ctx := c.Request.Context()
	req := &resources.AuthenticateRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !strings.Contains(req.Username, "@") {
		req.Username = req.Username + "@coteccons.vn"
	}

	resp := &resources.AuthenticateResponse{}
	resp.Token, err = h.RegisterService.Authenticate(ctx, req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: resp,
	})
}
