package exporter

import (
	"bytes"
	"fmt"
	
	"github.com/xuri/excelize/v2"
)

type Excel interface {
	CreateSheet(name string) Sheet
	Export() ([]byte, error)
	
	MustExport() []byte
}

type excelExporter struct {
	file   *excelize.File
	sheets []*sheet
}

func createExcelExporter() *excelExporter {
	return &excelExporter{
		file:   excelize.NewFile(),
		sheets: make([]*sheet, 0),
	}
}

func (e *excelExporter) CreateSheet(name string) Sheet {
	s := createSheet(e.file, name)
	e.sheets = append(e.sheets, s)
	return s
}

func (e *excelExporter) Export() ([]byte, error) {
	result := new(bytes.Buffer)
	for _, s := range e.sheets {
		sheetIndex, err := e.file.NewSheet(s.name)
		if err != nil {
			return []byte{}, err
		}
		e.file.SetActiveSheet(sheetIndex)
		for i, r := range s.rows {
			for j, c := range r.cols {
				if err := e.file.SetCellValue(s.name, fmt.Sprintf("%s%d", getLetterWithIndex(i), j+1), c.value); err != nil {
					return []byte{}, err
				}
			}
		}
	}
	if err := e.file.Write(result); err != nil {
		return []byte{}, err
	}
	if err := e.file.Close(); err != nil {
		return []byte{}, err
	}
	return result.Bytes(), nil
}

func (e *excelExporter) MustExport() []byte {
	r, err := e.Export()
	if err != nil {
		panic(err)
	}
	return r
}
