package exporter

type Exporter interface {
	Csv() Csv
	Excel() Excel
	Pdf(host string) Pdf
}

type exporter struct {
}

func New() Exporter {
	return &exporter{}
}

func (e *exporter) Csv() Csv {
	return createCsvExporter()
}

func (e *exporter) Excel() Excel {
	return createExcelExporter()
}

func (e *exporter) Pdf(host string) Pdf {
	return createPdfExporter(host)
}
