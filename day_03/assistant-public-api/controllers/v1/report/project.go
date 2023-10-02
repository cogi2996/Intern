package report

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ideal-forward/assistant-public-api/controllers/resources"
)

func (h Handler) CompareProject(c *gin.Context) {
	/*
		ctx := c.Request.Context()
		req := &resources.CompareProjectRequest{}

		err := http_parser.BindAndValid(c, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		var (
			startTime time.Time
			endTime   time.Time
		)

		firstDate := time.Date(req.First.Year, time.Month(req.First.Month), 1, 0, 0, 0, 0, time.UTC)
		secondDate := time.Date(req.Second.Year, time.Month(req.Second.Month), 1, 0, 0, 0, 0, time.UTC)
		if firstDate.Before(secondDate) {
			startTime = firstDate
			endTime = time.Date(req.Second.Year, time.Month(req.Second.Month+1), 1, 0, 0, 0, 0, time.UTC).Add(-1 * time.Second)
		} else {
			startTime = secondDate
			endTime = time.Date(req.First.Year, time.Month(req.First.Month+1), 1, 0, 0, 0, 0, time.UTC).Add(-1 * time.Second)
		}

		tasks, err := h.Task.ListByCons(ctx, &services.TaskFilters{
			StartTime:  startTime.UnixMilli(),
			EndTime:    endTime.UnixMilli(),
			ProjectIDs: []int64{req.First.ProjectID, req.Second.ProjectID},
		}, true)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	*/

	c.JSON(http.StatusOK, &resources.Response{
		Data: makeRandonProjectChart(),
	})
}

func makeRandonProjectChart() *resources.CompareProjectResponse {
	resp := &resources.CompareProjectResponse{
		BoardName: "Tổng số tiền",
		Unit:      "Tỉ đồng",
	}

	// months
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 1/2019",
		FirstValue:  100,
		SecondValue: 0,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 2/2019",
		FirstValue:  110,
		SecondValue: 180,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 3/2019",
		FirstValue:  130,
		SecondValue: 160,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 4/2019",
		FirstValue:  240,
		SecondValue: 180,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 5/2019",
		FirstValue:  100,
		SecondValue: 200,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 6/2019",
		FirstValue:  90,
		SecondValue: 110,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 7/2019",
		FirstValue:  160,
		SecondValue: 90,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 8/2019",
		FirstValue:  10,
		SecondValue: 200,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 9/2019",
		FirstValue:  120,
		SecondValue: 190,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 10/2019",
		FirstValue:  250,
		SecondValue: 220,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 11/2019",
		FirstValue:  100,
		SecondValue: 200,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 12/2019",
		FirstValue:  110,
		SecondValue: 290,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 1/2020",
		FirstValue:  100,
		SecondValue: 220,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 2/2020",
		FirstValue:  90,
		SecondValue: 150,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 3/2020",
		FirstValue:  190,
		SecondValue: 120,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 4/2020",
		FirstValue:  95,
		SecondValue: 190,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 5/2020",
		FirstValue:  100,
		SecondValue: 290,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 6/2020",
		FirstValue:  300,
		SecondValue: 200,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 7/2020",
		FirstValue:  130,
		SecondValue: 210,
	})
	resp.Months = append(resp.Months, &resources.DoubleCompareResultItem{
		Name:        "Tháng 8/2020",
		FirstValue:  100,
		SecondValue: 200,
	})

	// quarters
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 1/2019",
		FirstValue:  140,
		SecondValue: 540,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 2/2019",
		FirstValue:  430,
		SecondValue: 490,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 3/2019",
		FirstValue:  290,
		SecondValue: 480,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 4/2019",
		FirstValue:  460,
		SecondValue: 710,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 1/2020",
		FirstValue:  380,
		SecondValue: 490,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 2/2020",
		FirstValue:  495,
		SecondValue: 680,
	})
	resp.Quarters = append(resp.Quarters, &resources.DoubleCompareResultItem{
		Name:        "Quý 3/2020",
		FirstValue:  230,
		SecondValue: 410,
	})

	// years
	resp.Years = append(resp.Years, &resources.DoubleCompareResultItem{
		Name:        "2019",
		FirstValue:  1320,
		SecondValue: 2220,
	})
	resp.Years = append(resp.Years, &resources.DoubleCompareResultItem{
		Name:        "2020",
		FirstValue:  1105,
		SecondValue: 1580,
	})

	return resp
}
