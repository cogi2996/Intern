package services

import (
	"fmt"

	"github.com/ideal-forward/assistant-public-api/entities"
	"github.com/ideal-forward/assistant-public-api/models"
	"github.com/ideal-forward/assistant-public-api/pkg/excel"
	"github.com/xuri/excelize/v2"
)

type IExcel interface {
	// Create create file excel and response relative path
	Create(data []*entities.Task, pathFile string) error
}

type Excel struct {
	Model models.IExecutor
}

func NewExcel() IExcel {
	return &Excel{}
}

const ExcelSheet1 string = "Sheet1"

// STT | Mã nhiệm vụ | Nhiệm vụ | Mã NTP/NCC | Tên NTP/NCC | Giám sát | Số lương | Đơn giá | Thành tiền | Ngày bắt đầu | Ngày kết thúc | Ghi chú
func (e *Excel) Create(tasks []*entities.Task, pathFile string) error {
	f := excelize.NewFile()
	defer func() error {
		return f.Close()
	}()

	maker := excel.NewExcelMaker(f, "FF0000", "00FF00", "0000FF")

	index, err := f.NewSheet(ExcelSheet1)
	if err != nil {
		return err
	}

	f.SetDefaultFont("Arial")

	row := 1

	// Set header
	maker.SetHeader(ExcelSheet1, "A", row, "STT")
	maker.SetHeader(ExcelSheet1, "B", row, "Mã việc")
	maker.SetHeader(ExcelSheet1, "C", row, "Tên việc")
	maker.SetHeader(ExcelSheet1, "D", row, "Mã NTP/NCC")
	maker.SetHeader(ExcelSheet1, "E", row, "Tên NTP/NCC")
	maker.SetHeader(ExcelSheet1, "F", row, "Giám sát")
	maker.SetHeader(ExcelSheet1, "G", row, "Số lượng")
	maker.SetHeader(ExcelSheet1, "H", row, "Đơn giá")
	maker.SetHeader(ExcelSheet1, "I", row, "Thành tiền")
	maker.SetHeader(ExcelSheet1, "J", row, "Bắt đầu")
	maker.SetHeader(ExcelSheet1, "K", row, "Kết thúc")
	maker.SetHeader(ExcelSheet1, "L", row, "Trạng thái")
	maker.SetHeader(ExcelSheet1, "M", row, "Ghi chú")

	for i, val := range tasks {
		row = row + 1
		maker.SetInt(ExcelSheet1, "A", row, int64(i)+1)
		maker.SetString(ExcelSheet1, "B", row, val.Code)
		maker.SetString(ExcelSheet1, "C", row, val.Name)
		maker.SetString(ExcelSheet1, "D", row, val.GetExecutor().GetCode())
		maker.SetString(ExcelSheet1, "E", row, val.GetExecutor().GetName())
		maker.SetString(ExcelSheet1, "F", row, val.GetAcceptor().GetName())
		maker.SetInt(ExcelSheet1, "G", row, val.Quantity)
		maker.SetInt(ExcelSheet1, "H", row, val.Price)
		maker.SetFormula(ExcelSheet1, "I", row, fmt.Sprintf("G%d*H%d", row, row))
		maker.SetTime(ExcelSheet1, "J", row, TimestampToDate(val.StartTime))
		maker.SetTime(ExcelSheet1, "K", row, TimestampToDate(val.EndTime))
		maker.SetString(ExcelSheet1, "L", row, "")
		maker.SetString(ExcelSheet1, "M", row, "")
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs(pathFile); err != nil {
		return err
	}

	return nil
}
