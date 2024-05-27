package exporter

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
)

type Pdf interface {
	Orientation() PdfOrientation
	Write() PdfWriter
	Margin(top, right, bottom, left float32) Pdf
	Assets(paths ...string) Pdf
	
	Export() ([]byte, error)
	MustExport() []byte
}

type PdfWriter interface {
	Header(nodes ...gox.Node) PdfWriter
	Body(nodes ...gox.Node) PdfWriter
	Footer(nodes ...gox.Node) PdfWriter
}

type PdfOrientation interface {
	Portrait() Pdf
	Landscape() Pdf
}

type pdfExporter struct {
	endpoint    string
	buffer      *bytes.Buffer
	writer      *multipart.Writer
	orientation string
	margin      [4]float32
	assets      []string
	header      []gox.Node
	body        []gox.Node
	footer      []gox.Node
}

const (
	Portrait  = "portrait"
	Landscape = "landscape"
)

const (
	documentWidth  = "8.27"
	documentHeight = "11.7"
	indexFileName  = "index.html"
	headerFileName = "header.html"
	footerFileName = "footer.html"
)

const (
	filesField        = "files"
	landscapeField    = "landscape"
	paperWidthField   = "paperWidth"
	paperHeightField  = "paperHeight"
	marginTopField    = "marginTop"
	marginRightField  = "marginRight"
	marginBottomField = "marginBottom"
	marginLeftField   = "marginLeft"
)

func createPdfExporter(host string) *pdfExporter {
	buffer := new(bytes.Buffer)
	return &pdfExporter{
		endpoint:    createGotenbergEndpoint(host),
		buffer:      buffer,
		writer:      multipart.NewWriter(buffer),
		orientation: Portrait,
		margin:      [4]float32{0.5, 0.5, 0.5, 0.5},
		assets:      make([]string, 0),
		
	}
}

func (e *pdfExporter) Orientation() PdfOrientation {
	return e
}

func (e *pdfExporter) Write() PdfWriter {
	return e
}

func (e *pdfExporter) Margin(top, right, bottom, left float32) Pdf {
	e.margin = [4]float32{top, right, bottom, left}
	return e
}

func (e *pdfExporter) Assets(paths ...string) Pdf {
	e.assets = append(e.assets, paths...)
	return e
}

func (e *pdfExporter) Header(nodes ...gox.Node) PdfWriter {
	e.header = nodes
	return e
}

func (e *pdfExporter) Body(nodes ...gox.Node) PdfWriter {
	e.body = nodes
	return e
}

func (e *pdfExporter) Footer(nodes ...gox.Node) PdfWriter {
	e.footer = nodes
	return e
}

func (e *pdfExporter) Portrait() Pdf {
	e.orientation = Portrait
	return e
}

func (e *pdfExporter) Landscape() Pdf {
	e.orientation = Landscape
	return e
}

func (e *pdfExporter) Export() ([]byte, error) {
	if err := e.createFields(); err != nil {
		return []byte{}, err
	}
	if err := e.writer.Close(); err != nil {
		return []byte{}, err
	}
	res, err := http.Post(e.endpoint, e.writer.FormDataContentType(), e.buffer)
	if err != nil {
		return []byte{}, err
	}
	processed, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	if err := res.Body.Close(); err != nil {
		return []byte{}, err
	}
	return processed, nil
}

func (e *pdfExporter) MustExport() []byte {
	r, err := e.Export()
	if err != nil {
		panic(err)
	}
	return r
}

func (e *pdfExporter) createFields() error {
	if err := e.writer.WriteField(landscapeField, fmt.Sprintf("%v", e.orientation == Landscape)); err != nil {
		return err
	}
	if err := e.writer.WriteField(paperWidthField, documentWidth); err != nil {
		return err
	}
	if err := e.writer.WriteField(paperHeightField, documentHeight); err != nil {
		return err
	}
	if err := e.writer.WriteField(marginTopField, fmt.Sprintf("%.2f", e.margin[0])); err != nil {
		return err
	}
	if err := e.writer.WriteField(marginRightField, fmt.Sprintf("%.2f", e.margin[1])); err != nil {
		return err
	}
	if err := e.writer.WriteField(marginBottomField, fmt.Sprintf("%.2f", e.margin[2])); err != nil {
		return err
	}
	if err := e.writer.WriteField(marginLeftField, fmt.Sprintf("%.2f", e.margin[3])); err != nil {
		return err
	}
	if len(e.body) > 0 {
		if err := e.createIndexField(); err != nil {
			return err
		}
	}
	if len(e.header) > 0 {
		if err := e.createHeaderField(); err != nil {
			return err
		}
	}
	if len(e.footer) > 0 {
		if err := e.createFooterField(); err != nil {
			return err
		}
	}
	if len(e.assets) > 0 {
		if err := e.createAssetsFields(); err != nil {
			return err
		}
	}
	return nil
}

func (e *pdfExporter) createIndexField() error {
	bodyWriter, err := e.writer.CreateFormFile(filesField, indexFileName)
	if err != nil {
		return err
	}
	if _, err := bodyWriter.Write([]byte(gox.Render(e.body...))); err != nil {
		return err
	}
	return nil
}

func (e *pdfExporter) createHeaderField() error {
	headerWriter, err := e.writer.CreateFormFile(filesField, headerFileName)
	if err != nil {
		return err
	}
	if _, err := headerWriter.Write([]byte(gox.Render(e.header...))); err != nil {
		return err
	}
	return nil
}

func (e *pdfExporter) createFooterField() error {
	footerWriter, err := e.writer.CreateFormFile(filesField, footerFileName)
	if err != nil {
		return err
	}
	if _, err := footerWriter.Write([]byte(gox.Render(e.footer...))); err != nil {
		return err
	}
	return nil
}

func (e *pdfExporter) createAssetsFields() error {
	for _, path := range e.assets {
		pathParts := strings.Split(path, "/")
		name := pathParts[len(pathParts)-1]
		assetWriter, err := e.writer.CreateFormFile(filesField, name)
		if err != nil {
			return err
		}
		assetBytes, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if _, err := assetWriter.Write(assetBytes); err != nil {
			return err
		}
	}
	return nil
}
