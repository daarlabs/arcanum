package sender

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strings"
	
	"github.com/daarlabs/arcanum/util/constant/contentType"
	"github.com/daarlabs/arcanum/util/constant/dataType"
	"github.com/daarlabs/arcanum/util/constant/fileSuffix"
)

type Send interface {
	Status(statusCode int) Send
}

type ExtendableSend interface {
	Header() http.Header
	Error(e any) error
	Json(value any) error
	Html(value string) error
	Xml(value any) error
	Text(value string) error
	Bool(value bool) error
	Redirect(url string) error
	File(name string, bytes []byte) error
	Asset(name string, bytes []byte) error
}

type Sender struct {
	Bytes       []byte
	DataType    string
	ContentType string
	Value       string
	StatusCode  int
	r           *http.Request
	w           http.ResponseWriter
	write       *bool
}

func New(request *http.Request, response http.ResponseWriter, write ...*bool) *Sender {
	var w *bool
	if len(write) == 0 {
		w = new(bool)
		*w = true
	}
	if len(write) > 0 {
		w = write[0]
	}
	return &Sender{
		StatusCode: http.StatusOK,
		write:      w,
		w:          response,
		r:          request,
	}
}

func (s *Sender) Header() http.Header {
	return s.w.Header()
}

func (s *Sender) Status(statusCode int) Send {
	if !*s.write {
		return s
	}
	s.StatusCode = statusCode
	return s
}

func (s *Sender) Error(e any) error {
	if !*s.write {
		return nil
	}
	var err error
	switch v := e.(type) {
	case nil:
		return errors.New(http.StatusText(s.StatusCode))
	case string:
		err = errors.New(v)
	case error:
		err = v
	default:
		err = errors.New(fmt.Sprintf("%v", e))
	}
	s.Bytes = []byte(err.Error())
	s.DataType = dataType.Error
	s.ContentType = contentType.Text
	if s.StatusCode == http.StatusOK {
		s.StatusCode = http.StatusBadRequest
	}
	return err
}

func (s *Sender) Json(value any) error {
	if !*s.write {
		return nil
	}
	bytes, err := json.Marshal(Json{Result: value})
	s.Bytes = bytes
	s.DataType = dataType.Json
	s.ContentType = contentType.Json
	return err
}

func (s *Sender) Html(value string) error {
	if !*s.write {
		return nil
	}
	s.Bytes = []byte(value)
	s.DataType = dataType.Html
	s.ContentType = contentType.Html
	return nil
}

func (s *Sender) Xml(value any) error {
	if !*s.write {
		return nil
	}
	bytes, err := xml.Marshal(value)
	s.Bytes = bytes
	s.DataType = dataType.Xml
	s.ContentType = contentType.Xml
	return err
}

func (s *Sender) Text(value string) error {
	if !*s.write {
		return nil
	}
	s.Bytes = []byte(value)
	s.DataType = dataType.Text
	s.ContentType = contentType.Text
	return nil
}

func (s *Sender) Bool(value bool) error {
	if !*s.write {
		return nil
	}
	bytes, err := json.Marshal(Json{Result: value})
	s.Bytes = bytes
	s.DataType = dataType.Bool
	s.ContentType = contentType.Json
	return err
}

func (s *Sender) Redirect(url string) error {
	if !*s.write {
		return nil
	}
	s.Value = url
	if s.StatusCode == http.StatusOK {
		s.StatusCode = http.StatusMovedPermanently
	}
	s.DataType = dataType.Redirect
	return nil
}

func (s *Sender) File(name string, bytes []byte) error {
	if !*s.write {
		return nil
	}
	s.Value = name
	s.Bytes = bytes
	s.DataType = dataType.Stream
	s.ContentType = contentType.OctetStream
	return nil
}

func (s *Sender) Asset(name string, bytes []byte) error {
	if !*s.write {
		return nil
	}
	s.Value = name
	s.Bytes = bytes
	s.DataType = dataType.Asset
	s.ContentType = contentType.OctetStream
	if strings.HasSuffix(name, fileSuffix.Css) {
		s.ContentType = contentType.Css
	}
	if strings.HasSuffix(name, fileSuffix.Js) {
		s.ContentType = contentType.Javascript
	}
	return nil
}
