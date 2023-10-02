package excel

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

type ExcelMaker struct {
	file *excelize.File

	headerStyleID     int
	numberStyleID     int
	numberGrayStyleID int
	dateStyleID       int
	dateGrayStyleID   int
	stringStyleID     int
	stringGrayStyleID int
}

func NewExcelMaker(f *excelize.File, headerColor, lightColor, grayColor string) *ExcelMaker {
	result := &ExcelMaker{
		file: f,
	}

	top := excelize.Border{Type: "top", Style: 1, Color: "FF0000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "FF0000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "FF0000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "FF0000"}

	result.headerStyleID, _ = f.NewStyle(&excelize.Style{
		Border:        []excelize.Border{top, left, right, bottom},
		Font:          &excelize.Font{Color: "1f7f3b", Bold: true, Family: "Arial"},
		Fill:          excelize.Fill{Type: "pattern", Color: []string{headerColor}, Pattern: 1},
		Alignment:     &excelize.Alignment{Vertical: "center", Horizontal: "right"},
		NumFmt:        4, // https://xuri.me/excelize/en/style.html#number_format
		DecimalPlaces: 3,
	})

	result.numberStyleID, _ = result.makeNumberStyle(lightColor)
	result.dateStyleID, _ = result.makeDateStyle(lightColor)
	result.stringStyleID, _ = result.makeStringStyle(lightColor)

	result.numberGrayStyleID, _ = result.makeNumberStyle(grayColor)
	result.dateGrayStyleID, _ = result.makeDateStyle(grayColor)
	result.stringGrayStyleID, _ = result.makeStringStyle(grayColor)

	return result
}

func (m *ExcelMaker) makeNumberStyle(bgColor string) (int, error) {
	top := excelize.Border{Type: "top", Style: 1, Color: "FF0000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "FF0000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "FF0000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "FF0000"}

	// format fomular
	return m.file.NewStyle(&excelize.Style{
		Border:        []excelize.Border{top, left, right, bottom},
		Font:          &excelize.Font{Color: "1f7f3b", Bold: false, Family: "Arial"},
		Fill:          excelize.Fill{Type: "pattern", Color: []string{bgColor}, Pattern: 1},
		Alignment:     &excelize.Alignment{Vertical: "center", Horizontal: "right"},
		NumFmt:        4, // https://xuri.me/excelize/en/style.html#number_format
		DecimalPlaces: 3,
	})
}

func (m *ExcelMaker) makeDateStyle(bgColor string) (int, error) {
	top := excelize.Border{Type: "top", Style: 1, Color: "FF0000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "FF0000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "FF0000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "FF0000"}

	exp := "dd-mm-yyyy HH:ss"
	return m.file.NewStyle(&excelize.Style{
		Border:       []excelize.Border{top, left, right, bottom},
		Font:         &excelize.Font{Color: "1f7f3b", Bold: false, Family: "Arial"},
		Fill:         excelize.Fill{Type: "pattern", Color: []string{bgColor}, Pattern: 1},
		Alignment:    &excelize.Alignment{Vertical: "center", Horizontal: "right"},
		CustomNumFmt: &exp,
	})
}

func (m *ExcelMaker) makeStringStyle(bgColor string) (int, error) {
	top := excelize.Border{Type: "top", Style: 1, Color: "FF0000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "FF0000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "FF0000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "FF0000"}

	return m.file.NewStyle(&excelize.Style{
		Border:    []excelize.Border{top, left, right, bottom},
		Font:      &excelize.Font{Color: "0000FF", Bold: false, Family: "Arial"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{bgColor}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "left"},
	})
}

func (m *ExcelMaker) SetInt(sheet, column string, row int, value int64) {
	cell := fmt.Sprintf("%s%d", column, row)
	m.file.SetCellValue(sheet, cell, int(value))
	if row%2 == 0 {
		m.file.SetCellStyle(sheet, cell, cell, m.numberStyleID)
	} else {
		m.file.SetCellStyle(sheet, cell, cell, m.numberGrayStyleID)
	}
}

func (m *ExcelMaker) SetHeader(sheet, column string, row int, title string) {
	cell := fmt.Sprintf("%s%d", column, row)
	m.file.SetCellStr(sheet, cell, title)
	m.file.SetCellStyle(sheet, cell, cell, m.headerStyleID)
}

func (m *ExcelMaker) SetTime(sheet, column string, row int, value time.Time) {
	cell := fmt.Sprintf("%s%d", column, row)
	m.file.SetCellValue(sheet, cell, value)
	if row%2 == 0 {
		m.file.SetCellStyle(sheet, cell, cell, m.dateStyleID)
	} else {
		m.file.SetCellStyle(sheet, cell, cell, m.dateGrayStyleID)
	}
}

func (m *ExcelMaker) SetString(sheet, column string, row int, value string) {
	cell := fmt.Sprintf("%s%d", column, row)
	m.file.SetCellStr(sheet, cell, value)
	if row%2 == 0 {
		m.file.SetCellStyle(sheet, cell, cell, m.stringStyleID)
	} else {
		m.file.SetCellStyle(sheet, cell, cell, m.stringGrayStyleID)
	}
}

func (m *ExcelMaker) SetFormula(sheet, column string, row int, formula string) {
	cell := fmt.Sprintf("%s%d", column, row)
	m.file.SetCellFormula(sheet, cell, formula)
	if row%2 == 0 {
		m.file.SetCellStyle(sheet, cell, cell, m.numberStyleID)
	} else {
		m.file.SetCellStyle(sheet, cell, cell, m.numberGrayStyleID)
	}
}
