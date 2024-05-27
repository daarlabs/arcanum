package exporter

import (
	"bytes"
	"encoding/csv"
)

type Csv interface {
	Divider(divider rune) Csv
	Row() Row
	Export() ([]byte, error)
	
	MustExport() []byte
}

type csvExporter struct {
	divider rune
	rows    []*row
}

func createCsvExporter() *csvExporter {
	return &csvExporter{
		divider: ';',
		rows:    make([]*row, 0),
	}
}

func (e *csvExporter) Divider(divider rune) Csv {
	e.divider = divider
	return e
}

func (e *csvExporter) Row() Row {
	r := createRow()
	e.rows = append(e.rows, r)
	return r
}

func (e *csvExporter) Export() ([]byte, error) {
	result := new(bytes.Buffer)
	writer := csv.NewWriter(result)
	writer.Comma = e.divider
	for _, r := range e.rows {
		if err := writer.Write(createStringSliceFromRow(r)); err != nil {
			return []byte{}, err
		}
	}
	if err := writer.Error(); err != nil {
		return []byte{}, err
	}
	writer.Flush()
	return result.Bytes(), nil
}

func (e *csvExporter) MustExport() []byte {
	r, err := e.Export()
	if err != nil {
		panic(err)
	}
	return r
}
