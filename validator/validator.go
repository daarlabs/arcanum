package validator

import (
	"reflect"
	"regexp"
)

type Validator interface {
	Json(s Schema, data any) (bool, Errors)
}

type validator struct {
	config Config
}

const (
	emailValidator = "[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+(?:\\.[-A-Za-z0-9!#$%&'*+/=?^_`{|}~]+)*@(?:[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?\\.)+[A-Za-z0-9](?:[-A-Za-z0-9]*[A-Za-z0-9])?"
)

var (
	emailMatcher = regexp.MustCompile(emailValidator)
)

func New(config ...Config) Validator {
	v := &validator{}
	if len(config) > 0 {
		v.config = config[0]
	}
	v.config = v.createDefaultConfig(v.config)
	return v
}

func (v *validator) Json(s Schema, data any) (bool, Errors) {
	errors := make(Errors)
	dt := reflect.TypeOf(data)
	dv := reflect.ValueOf(data)
	switch dt.Kind() {
	case reflect.Struct:
		errors = v.validateStruct(s.(*schema), dt, dv)
	case reflect.Map:
		errors = v.validateMap(s.(*schema), dv)
	}
	return len(errors) == 0, errors
}

func (v *validator) validateStruct(s *schema, mt reflect.Type, mv reflect.Value) Errors {
	errors := make(Errors)
	for i := 0; i < mv.NumField(); i++ {
		jsonKey := mt.Field(i).Tag.Get("json")
		if len(jsonKey) == 0 {
			continue
		}
		f, ok := s.shape[jsonKey]
		if !ok {
			continue
		}
		errs := v.validateField(f, mv.Field(i).Interface())
		if len(errs) > 0 {
			errors[jsonKey] = errs
		}
	}
	return errors
}

func (v *validator) validateMap(s *schema, mv reflect.Value) Errors {
	errors := make(Errors)
	keys := mv.MapKeys()
	for name := range s.shape {
		exist := false
		for _, k := range keys {
			if k.String() == name {
				exist = true
			}
		}
		if exist {
			continue
		}
		errors[name] = []string{v.config.Messages.Required}
	}
	for _, k := range keys {
		f, ok := s.shape[k.String()]
		if !ok {
			continue
		}
		errs := v.validateField(f, mv.MapIndex(k).Interface())
		if len(errs) > 0 {
			errors[k.String()] = errs
		}
	}
	return errors
}

func (v *validator) validateField(f *field, value any) []string {
	errs := make([]string, 0)
	switch val := value.(type) {
	case string:
		if f.email && !emailMatcher.MatchString(val) {
			errs = append(errs, v.config.Messages.Email)
		}
		if f.required && len(val) == 0 {
			errs = append(errs, v.config.Messages.Required)
		}
		if f.min > 0 && len(val) < f.min {
			errs = append(errs, v.config.Messages.MinText)
		}
		if f.max > 0 && len(val) > f.max {
			errs = append(errs, v.config.Messages.MaxText)
		}
	case int:
		if f.required && val == 0 {
			errs = append(errs, v.config.Messages.Required)
		}
		if f.min > 0 && val < f.min {
			errs = append(errs, v.config.Messages.MinNumber)
		}
		if f.max > 0 && val > f.max {
			errs = append(errs, v.config.Messages.MaxNumber)
		}
	case float32:
		if f.required && val == 0 {
			errs = append(errs, v.config.Messages.Required)
		}
		if f.min > 0 && val < float32(f.min) {
			errs = append(errs, v.config.Messages.MinNumber)
		}
		if f.max > 0 && val > float32(f.max) {
			errs = append(errs, v.config.Messages.MaxNumber)
		}
	case float64:
		if f.required && val == 0 {
			errs = append(errs, v.config.Messages.Required)
		}
		if f.min > 0 && val < float64(f.min) {
			errs = append(errs, v.config.Messages.MinNumber)
		}
		if f.max > 0 && val > float64(f.max) {
			errs = append(errs, v.config.Messages.MaxNumber)
		}
	case bool:
		if f.required && !val {
			errs = append(errs, v.config.Messages.Required)
		}
	}
	return errs
}

func (v *validator) createDefaultConfig(config Config) Config {
	if len(config.Messages.Email) == 0 {
		config.Messages.Email = defaultEmailMessage
	}
	if len(config.Messages.Required) == 0 {
		config.Messages.Required = defaultRequiredMessage
	}
	if len(config.Messages.MinText) == 0 {
		config.Messages.MinText = defaultMinTextMessage
	}
	if len(config.Messages.MaxText) == 0 {
		config.Messages.MaxText = defaultMaxTextMessage
	}
	if len(config.Messages.MinNumber) == 0 {
		config.Messages.MinNumber = defaultMinNumberMessage
	}
	if len(config.Messages.MaxNumber) == 0 {
		config.Messages.MaxNumber = defaultMaxNumberMessage
	}
	return config
}
