package service

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"logistic_api/pkg/logger"
	"os"
	"path/filepath"
	"time"
)

const (
	reportDir       = "table_reports"
	sheetReportName = "Отчет"
)

type Report interface {
	GetDateStart() time.Time
	GetDateEnd() time.Time
}

type TableReport struct {
	UniqueID  string    `json:"report_name"`
	DateStart time.Time `json:"date_start" binding:"required" json:"date_start"`
	DateEnd   time.Time `json:"date_end" binding:"required" json:"date_end"`
	TableData TableData `json:"table_data" json:"table_data"`
}

func NewTableReport(tableData TableData) *TableReport {
	return &TableReport{TableData: tableData}
}

func (r *TableReport) GetDateStart() time.Time {
	return r.DateStart
}

func (r *TableReport) GetDateEnd() time.Time {
	return r.DateEnd
}

func dirValidate() error {
	if _, err := os.Stat(reportDir); os.IsNotExist(err) {
		err := os.Mkdir(reportDir, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TableReport) ToExcel() (string, error) {
	fileName := fmt.Sprintf("%s_%s.xlsx", t.UniqueID, time.Now().Format("2006-01-02"))
	filePath := filepath.Join(reportDir, fileName)

	file := excelize.NewFile()

	err := file.SetSheetName("Sheet1", sheetReportName)
	if err != nil {
		return "", err
	}

	for rowIndex, row := range t.TableData.Rows {
		for colIndex, cellData := range row {
			cell, err := excelize.CoordinatesToCellName(colIndex+1, rowIndex+1)

			if err != nil {
				return "", err
			}

			err = file.SetCellStr(sheetReportName, cell, fmt.Sprintf("%v", cellData))

			if err != nil {
				return "", err
			}
		}
	}

	err = file.SaveAs(filePath)
	if err != nil {
		return "", err
	}

	return filePath, nil
}

type TableData struct {
	Rows [][]string `json:"rows"`
}

type DiagramReport struct {
	DateStart   time.Time   `json:"date_start" binding:"required"`
	DateEnd     time.Time   `json:"date_end" binding:"required"`
	DiagramData DiagramData `json:"diagram_data"`
}

func (r *DiagramReport) GetDateStart() time.Time {
	return r.DateStart
}

func (r *DiagramReport) GetDateEnd() time.Time {
	return r.DateEnd
}

type DiagramData struct {
	Labels []string  `json:"labels"`
	Data   []float64 `json:"data"`
}

func init() {
	err := dirValidate()
	if err != nil {
		logger.Log.Error(err.Error())
	}
}
