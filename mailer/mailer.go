package mailer

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime"
	"mime/multipart"
	"net/http"
	"net/smtp"
	"strings"
	
	"github.com/daarlabs/arcanum/gox"
)

type Mailer interface {
	Attachment(name string, data []byte) Mailer
	Body(nodes ...gox.Node) Mailer
	Bytes() []byte
	Copy(values ...string) Mailer
	From(from string) Mailer
	HiddenCopy(values ...string) Mailer
	Subject(subject string) Mailer
	Title(title string) Mailer
	To(to ...string) Mailer
	Send() error
	String() string
	
	MustSend()
}

type mailer struct {
	config       Config
	attachments  []attachment
	from         string
	to           []string
	toCopy       []string
	toHiddenCopy []string
	subject      string
	title        string
	nodes        []gox.Node
}

type attachment struct {
	name string
	data []byte
}

func New(config Config) Mailer {
	return &mailer{config: config}
}

func (m *mailer) Attachment(name string, data []byte) Mailer {
	m.attachments = append(m.attachments, attachment{name, data})
	return m
}

func (m *mailer) Body(nodes ...gox.Node) Mailer {
	m.nodes = nodes
	return m
}

func (m *mailer) Bytes() []byte {
	return m.createBody()
}

func (m *mailer) Copy(values ...string) Mailer {
	m.toCopy = values
	return m
}

func (m *mailer) From(value string) Mailer {
	m.from = value
	return m
}

func (m *mailer) HiddenCopy(values ...string) Mailer {
	m.toHiddenCopy = values
	return m
}

func (m *mailer) Send() error {
	return smtp.SendMail(
		fmt.Sprintf("%s:%d", m.config.Host, m.config.Port),
		smtp.PlainAuth("", m.config.User, m.config.Password, m.config.Host),
		m.from,
		m.to,
		m.createBody(),
	)
}

func (m *mailer) MustSend() {
	err := m.Send()
	if err != nil {
		panic(err)
	}
}

func (m *mailer) String() string {
	return string(m.createBody())
}

func (m *mailer) Subject(value string) Mailer {
	m.subject = value
	return m
}

func (m *mailer) Title(value string) Mailer {
	m.title = value
	return m
}

func (m *mailer) To(values ...string) Mailer {
	m.to = values
	return m
}

func (m *mailer) createBody() []byte {
	buf := new(bytes.Buffer)
	attachmentsExist := len(m.attachments) > 0
	buf.WriteString(fmt.Sprintf("From: %s\r\n", m.from))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(m.to, ",")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mime.BEncoding.Encode("utf-8", m.subject)))
	if len(m.toCopy) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.toCopy, ",")))
	}
	if len(m.toHiddenCopy) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.toHiddenCopy, ",")))
	}
	buf.WriteString("MIME-version: 1.0;\r\n")
	w := multipart.NewWriter(buf)
	boundary := w.Boundary()
	if !attachmentsExist {
		buf.WriteString("Content-Type: text/html; charset=utf-8\n")
	}
	if attachmentsExist {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	}
	buf.WriteString(gox.Render(m.nodes...))
	if attachmentsExist {
		for _, a := range m.attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(a.data)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", a.name))
			b := make([]byte, base64.StdEncoding.EncodedLen(len(a.data)))
			base64.StdEncoding.Encode(b, a.data)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}
		
		buf.WriteString("--")
	}
	return buf.Bytes()
}

func (m *mailer) encodeRFC2047(value string) string {
	return mime.BEncoding.Encode("utf-8", value)
}
