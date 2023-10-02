package resources

type SingleCompareResultItem struct {
	Name  string  `json:"name"`
	Value float32 `json:"value"`
}

type DoubleCompareResultItem struct {
	Name        string  `json:"name"`
	FirstValue  float32 `json:"first_value"`
	SecondValue float32 `json:"second_value"`
}

type ComparedProject struct {
	ProjectID int64 `json:"project_id"`
	Year      int   `json:"year"`
	Month     int   `json:"month"`
}

type CompareProjectRequest struct {
	First      *ComparedProject `json:"first"`
	Second     *ComparedProject `json:"second"`
	IsSameTime bool             `json:"is_same_time"`
}

type CompareProjectResponse struct {
	BoardName string                     `json:"board_name"`
	Unit      string                     `json:"unit"`
	Months    []*DoubleCompareResultItem `json:"months"`
	Quarters  []*DoubleCompareResultItem `json:"quarters"`
	Years     []*DoubleCompareResultItem `json:"years"`
}

type ConstructorSelfCompareRequest struct {
	ExecutorID           int    `form:"executor_id"`
	PackageWorkStartDate string `form:"package_work_start_date"`
	PackageWorkEndDate   string `form:"package_work_end_date"`
	DayWorkStartDate     string `form:"day_work_start_date"`
	DayWorkEndDate       string `form:"day_work_end_date"`
}

type CompareChart struct {
	Name  string                     `json:"board_name"`
	Unit  string                     `json:"unit"`
	Chart []*DoubleCompareResultItem `json:"chart"`
}

type PackageWorkChart struct {
	Name  string                     `json:"board_name"`
	Unit  string                     `json:"unit"`
	Chart []*SingleCompareResultItem `json:"chart"`
}

type DayWorkChart struct {
	Name  string                     `json:"board_name"`
	Unit  string                     `json:"unit"`
	Chart []*SingleCompareResultItem `json:"chart"`
}

type ConstructorSelfCompareResponse struct {
	CompareChart     *CompareChart     `json:"compare"`
	PackageWorkChart *PackageWorkChart `json:"package_work"`
	DayWorkChart     *DayWorkChart     `json:"day_work"`
}
