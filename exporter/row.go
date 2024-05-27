package exporter

type Row interface {
	Col() Col
}

type row struct {
	cols []*col
}

func createRow() *row {
	return &row{
		cols: make([]*col, 0),
	}
}

func (r *row) Col() Col {
	c := &col{}
	r.cols = append(r.cols, c)
	return c
}
