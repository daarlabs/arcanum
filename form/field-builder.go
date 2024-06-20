package form

import (
	"fmt"
	"strings"
	"time"
	
	"golang.org/x/exp/constraints"
)

type FieldBuilder struct {
	dataType   string
	fieldType  string
	id         string
	autofocus  bool
	disabled   bool
	multiple   bool
	valid      bool
	name       string
	label      string
	text       string
	size       int
	value      any
	validators []validator
	messages   Messages
}

type FieldConfig struct {
	fieldType string
	dataType  string
	value     any
}

const (
	fieldTypeButton        = "button"
	fieldTypeCheckbox      = "checkbox"
	fieldTypeColor         = "color"
	fieldTypeDate          = "date"
	fieldTypeDateTimeLocal = "datetime-local"
	fieldTypeEmail         = "email"
	fieldTypeFile          = "file"
	fieldTypeHidden        = "hidden"
	fieldTypeImage         = "image"
	fieldTypeMonth         = "month"
	fieldTypeNumber        = "number"
	fieldTypePassword      = "password"
	fieldTypeRadio         = "radio"
	fieldTypeRange         = "range"
	fieldTypeReset         = "reset"
	fieldTypeSearch        = "search"
	fieldTypeSubmit        = "submit"
	fieldTypeTel           = "tel"
	fieldTypeText          = "text"
	fieldTypeTime          = "time"
	fieldTypeUrl           = "url"
	fieldTypeWeek          = "week"
	
	fieldDataTypeBool   = "bool"
	fieldDataTypeFile   = "file"
	fieldDataTypeFloat  = "float"
	fieldDataTypeInt    = "int"
	fieldDataTypeInt64  = "int64"
	fieldDataTypeString = "string"
	fieldDataTypeTime   = "time"
	
	fieldTimeFormat = "2006-01-02 15:04:05.999999999 +0000 UTC"
)

func Add(name string) *FieldBuilder {
	return &FieldBuilder{
		name:       strings.ReplaceAll(name, "_", "-"),
		validators: make([]validator, 0),
	}
}

func (b *FieldBuilder) Label(label string) *FieldBuilder {
	b.label = label
	return b
}

func (b *FieldBuilder) Autofocus(autofocus ...bool) *FieldBuilder {
	af := true
	if len(autofocus) > 0 {
		af = autofocus[0]
	}
	b.autofocus = af
	return b
}

func (b *FieldBuilder) Disabled(disabled ...bool) *FieldBuilder {
	d := true
	if len(disabled) > 0 {
		d = disabled[0]
	}
	b.disabled = d
	return b
}

func (b *FieldBuilder) Text(text any) *FieldBuilder {
	b.text = fmt.Sprintf("%v", text)
	return b
}

func (b *FieldBuilder) With(config FieldConfig, validators ...Validator) *FieldBuilder {
	switch config.value.(type) {
	case []any:
		createFieldType[any](b, config.fieldType, config.dataType, config.value.([]any)...)
	case []string:
		createFieldType[string](b, config.fieldType, config.dataType, config.value.([]string)...)
	case []int:
		createFieldType[int](b, config.fieldType, config.dataType, config.value.([]int)...)
	case []int64:
		createFieldType[int64](b, config.fieldType, config.dataType, config.value.([]int64)...)
	case []float32:
		createFieldType[float32](b, config.fieldType, config.dataType, config.value.([]float32)...)
	case []float64:
		createFieldType[float64](b, config.fieldType, config.dataType, config.value.([]float64)...)
	case []bool:
		createFieldType[bool](b, config.fieldType, config.dataType, config.value.([]bool)...)
	case []Multipart:
		createFieldType[Multipart](b, config.fieldType, config.dataType, config.value.([]Multipart)...)
	case []time.Time:
		createFieldType[time.Time](b, config.fieldType, config.dataType, config.value.([]time.Time)...)
	}
	for _, v := range validators {
		b.validators = append(b.validators, v.(validator))
	}
	return b
}

func (b *FieldBuilder) isRequired() bool {
	for _, v := range b.validators {
		if v.validatorType == validatorTypeRequired {
			return true
		}
	}
	return false
}

func Button(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeButton,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Checkbox(value ...bool) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeCheckbox,
		dataType:  fieldDataTypeBool,
		value:     value,
	}
}

func Color(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeColor,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Date(value ...time.Time) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeDate,
		dataType:  fieldDataTypeTime,
		value:     value,
	}
}

func DateTimeLocal(value ...time.Time) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeDateTimeLocal,
		dataType:  fieldDataTypeTime,
		value:     value,
	}
}

func Email(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeEmail,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func File(value ...Multipart) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeFile,
		dataType:  fieldDataTypeFile,
		value:     value,
	}
}

func Hidden[T comparable](value ...T) FieldConfig {
	var dataType string
	switch any(*new(T)).(type) {
	case float32, float64:
		dataType = fieldDataTypeFloat
	case int:
		dataType = fieldDataTypeInt
	case int64:
		dataType = fieldDataTypeInt64
	case string:
		dataType = fieldDataTypeString
	case bool:
		dataType = fieldDataTypeBool
	case time.Time:
		dataType = fieldDataTypeTime
	}
	return FieldConfig{
		fieldType: fieldTypeHidden,
		dataType:  dataType,
		value:     value,
	}
}

func Image(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeImage,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func (b *FieldBuilder) Id(id string) *FieldBuilder {
	b.id = strings.ReplaceAll(id, "_", "-")
	return b
}

func Month(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeMonth,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func (b *FieldBuilder) Multiple(size ...int) *FieldBuilder {
	b.multiple = true
	if len(size) > 0 {
		b.size = size[0]
	}
	return b
}

func Number[T constraints.Float | constraints.Integer](value ...T) FieldConfig {
	v := *new(T)
	switch any(v).(type) {
	case int64:
		return FieldConfig{
			fieldType: fieldTypeNumber,
			dataType:  fieldDataTypeInt64,
			value:     value,
		}
	case int:
		return FieldConfig{
			fieldType: fieldTypeNumber,
			dataType:  fieldDataTypeInt,
			value:     value,
		}
	}
	return FieldConfig{
		fieldType: fieldTypeNumber,
		dataType:  fieldDataTypeFloat,
		value:     value,
	}
}

func Password(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypePassword,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Radio(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeRadio,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Range(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeRange,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Reset(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeReset,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Search(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeSearch,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Submit(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeSubmit,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Tel(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeTel,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Text(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeText,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Time(value ...time.Time) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeTime,
		dataType:  fieldDataTypeTime,
		value:     value,
	}
}

func Url(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeUrl,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func Week(value ...string) FieldConfig {
	return FieldConfig{
		fieldType: fieldTypeWeek,
		dataType:  fieldDataTypeString,
		value:     value,
	}
}

func createFieldType[T any](b *FieldBuilder, fieldType, dataType string, values ...T) {
	b.fieldType = fieldType
	b.dataType = dataType
	if !b.multiple {
		b.multiple = len(values) > 1
	}
	if b.multiple {
		b.value = values
	}
	if !b.multiple && len(values) > 0 {
		b.value = values[0]
	}
	if b.value == nil {
		if b.multiple {
			b.value = make([]T, 0)
		}
		if !b.multiple {
			b.value = *new(T)
		}
	}
}
