package form

type Messages struct {
	Email     string `json:"email" toml:"email" yaml:"email"`
	Required  string `json:"required" toml:"required" yaml:"required"`
	MinText   string `json:"minText" toml:"minText" yaml:"minText"`
	MaxText   string `json:"maxText" toml:"maxText" yaml:"maxText"`
	MinNumber string `json:"minNumber" toml:"minNumber" yaml:"minNumber"`
	MaxNumber string `json:"maxNumber" toml:"maxNumber" yaml:"maxNumber"`
	Multipart string `json:"multipart" toml:"multipart" yaml:"multipart"`
	Invalid   string `json:"invalid" toml:"invalid" yaml:"invalid"`
}

const (
	defaultRequiredMessage  = "field is required"
	defaultEmailMessage     = "email value is invalid"
	defaultMinTextMessage   = "field length is smaller than should be"
	defaultMaxTextMessage   = "field length is higher than should be"
	defaultMinNumberMessage = "field value is smaller than should be"
	defaultMaxNumberMessage = "field value is higher than should be"
	defaultMultipartMessage = "invalid file"
	defaultInvalidMessage   = "invalid value"
)

var (
	defaultMessages = Messages{
		Email:     defaultEmailMessage,
		Required:  defaultRequiredMessage,
		MinText:   defaultMinTextMessage,
		MaxText:   defaultMaxTextMessage,
		MinNumber: defaultMinNumberMessage,
		MaxNumber: defaultMaxNumberMessage,
		Multipart: defaultMultipartMessage,
		Invalid:   defaultInvalidMessage,
	}
)
