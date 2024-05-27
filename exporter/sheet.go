package exporter

import "github.com/xuri/excelize/v2"

type Sheet interface {
	Row() Row
}

type sheet struct {
	file *excelize.File
	name string
	rows []*row
}

func createSheet(file *excelize.File, name string) *sheet {
	return &sheet{
		file: file,
		name: name,
		rows: make([]*row, 0),
	}
}

func (s *sheet) Row() Row {
	r := createRow()
	s.rows = append(s.rows, r)
	return r
}
