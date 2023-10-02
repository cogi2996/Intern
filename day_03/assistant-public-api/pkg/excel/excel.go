package excel

import (
	"fmt"
	"time"

	"github.com/xuri/excelize/v2"
)

func CreateExcel() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		fmt.Println(err)
		return
	}

	f.SetDefaultFont("Arial")

	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetCellInt("Sheet1", "C1", 200)
	f.SetCellInt("Sheet1", "C2", 400)

	// Set style
	// https://xuri.me/excelize/en/example/calendar.html

	top := excelize.Border{Type: "top", Style: 1, Color: "FF0000"}
	left := excelize.Border{Type: "left", Style: 1, Color: "FF0000"}
	right := excelize.Border{Type: "right", Style: 1, Color: "FF0000"}
	bottom := excelize.Border{Type: "bottom", Style: 1, Color: "FF0000"}

	// format fomular
	styleID, err := f.NewStyle(&excelize.Style{
		Border:        []excelize.Border{top, left, right, bottom},
		Font:          &excelize.Font{Color: "1f7f3b", Bold: true, Family: "Arial"},
		Fill:          excelize.Fill{Type: "pattern", Color: []string{"E6F4EA"}, Pattern: 1},
		Alignment:     &excelize.Alignment{Vertical: "center", Horizontal: "right"},
		NumFmt:        4, // https://xuri.me/excelize/en/style.html#number_format
		DecimalPlaces: 3,
	})
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellFormula("Sheet1", "C3", "SUM(C1,C2)")
	f.SetCellStyle("Sheet1", "C3", "C3", styleID)

	// format datetime
	exp := "dd-mm-yyyy HH:ss"
	styleID, err = f.NewStyle(&excelize.Style{
		Border:    []excelize.Border{top, left, right, bottom},
		Font:      &excelize.Font{Color: "1f7f3b", Bold: true, Family: "Arial"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"E6F4EA"}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "right"},
		// NumFmt:        22, // https://xuri.me/excelize/en/style.html#number_format
		CustomNumFmt: &exp,
		// DecimalPlaces: 3,
	})
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellValue("Sheet1", "C4", time.Now().Add(96*time.Hour))
	f.SetCellStyle("Sheet1", "C4", "C4", styleID)

	// format string
	styleID, err = f.NewStyle(&excelize.Style{
		Border:    []excelize.Border{top, right, bottom},
		Font:      &excelize.Font{Color: "0000FF", Bold: true, Family: "Arial"},
		Fill:      excelize.Fill{Type: "pattern", Color: []string{"00FF00"}, Pattern: 1},
		Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "left"},
	})
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle("Sheet1", "B3", "B3", styleID)
	f.SetCellStr("Sheet1", "B3", "Tổng:")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("ZBook1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ReadExcel() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Get value from cell by given worksheet name and cell reference.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}
