package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
	"github.com/ideal-forward/assistant-public-api/pkg/http_parser"
)

func (h Handler) CompareSeflConstructor(c *gin.Context) {
	// ctx := c.Request.Context()
	req := &resources.ConstructorSelfCompareRequest{}

	err := http_parser.BindAndValid(c, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, &resources.Response{
		Data: makeRandonSeflConstructorChart(),
	})
}

func makeRandonSeflConstructorChart() *resources.ConstructorSelfCompareResponse {
	resp := &resources.ConstructorSelfCompareResponse{
		CompareChart: &resources.CompareChart{
			Name: "Tổng tiền",
			Unit: "Triệu đồng",
		},
		PackageWorkChart: &resources.PackageWorkChart{
			Name: "Số khoán",
			Unit: "Việc",
		},
		DayWorkChart: &resources.DayWorkChart{
			Name: "Số công nhật",
			Unit: "Công nhật",
		},
	}

	// months
	resp.CompareChart.Chart = append(resp.CompareChart.Chart, &resources.DoubleCompareResultItem{
		Name:        "Tuần 29/05/23",
		FirstValue:  100,
		SecondValue: 0,
	})
	resp.CompareChart.Chart = append(resp.CompareChart.Chart, &resources.DoubleCompareResultItem{
		Name:        "Tuần 05/06/23",
		FirstValue:  110,
		SecondValue: 180,
	})
	resp.CompareChart.Chart = append(resp.CompareChart.Chart, &resources.DoubleCompareResultItem{
		Name:        "Tuần 12/06/23",
		FirstValue:  130,
		SecondValue: 160,
	})
	resp.CompareChart.Chart = append(resp.CompareChart.Chart, &resources.DoubleCompareResultItem{
		Name:        "Tuần 19/06/23",
		FirstValue:  240,
		SecondValue: 180,
	})

	///
	resp.PackageWorkChart.Chart = append(resp.PackageWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 29/05/23",
		Value: 50,
	})
	resp.PackageWorkChart.Chart = append(resp.PackageWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 05/06/23",
		Value: 100,
	})
	resp.PackageWorkChart.Chart = append(resp.PackageWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 12/06/23",
		Value: 75,
	})
	resp.PackageWorkChart.Chart = append(resp.PackageWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 19/06/23",
		Value: 125,
	})

	///
	resp.DayWorkChart.Chart = append(resp.DayWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 29/05/23",
		Value: 15,
	})
	resp.DayWorkChart.Chart = append(resp.DayWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 05/06/23",
		Value: 20,
	})
	resp.DayWorkChart.Chart = append(resp.DayWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 12/06/23",
		Value: 50,
	})
	resp.DayWorkChart.Chart = append(resp.DayWorkChart.Chart, &resources.SingleCompareResultItem{
		Name:  "Tuần 19/06/23",
		Value: 35,
	})

	return resp
}
