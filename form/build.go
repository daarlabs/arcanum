package form

import (
	"fmt"
	"net/http"
	"reflect"
	"time"
	
	"github.com/iancoleman/strcase"
)

const (
	baseFormFieldName = "Form"
)

func MustBuild[T any](b *Builder) T {
	r, err := Build[T](b)
	if err != nil {
		panic(err)
	}
	return r
}

func getContentType(b *Builder) string {
	for _, f := range b.fields {
		if f.dataType == fieldDataTypeFile {
			return contentTypeMultipartForm
		}
	}
	return contentTypeForm
}

func Build[T any](b *Builder) (T, error) {
	if b.request == nil {
		return buildForm[T](b), nil
	}
	b.submitted = isFormSubmitted(b.request)
	b.contentType = getContentType(b)
	reqFormData, reqFormFiles, err := processRequest(b.request, b.limit)
	if err != nil {
		return *new(T), fmt.Errorf("error processing request to form: %w", err)
	}
	if len(reqFormData) > 0 {
		processFormData(b, reqFormData)
	}
	if len(reqFormFiles) > 0 {
		err := processFormFiles(b, reqFormFiles)
		if err != nil {
			return *new(T), err
		}
	}
	return buildForm[T](b), nil
}

func buildForm[T any](b *Builder) T {
	form := new(T)
	formRef := reflect.ValueOf(form)
	for i, fb := range b.fields {
		fb.messages = b.messages
		errors := buildFormField(formRef, fb, b.request)
		b.fields[i].valid = len(errors) == 0
	}
	buildBaseForm(formRef, b)
	return *form
}

func buildFormField(formRef reflect.Value, fb *FieldBuilder, req *http.Request) []string {
	errors := make([]string, 0)
	formField := formRef.Elem().FieldByName(strcase.ToCamel(fb.name))
	if !formField.IsValid() {
		return errors
	}
	switch fb.dataType {
	case fieldDataTypeString:
		if fb.multiple {
			field := createFormField[[]string](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[string](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	case fieldDataTypeFloat:
		if fb.multiple {
			field := createFormField[[]float64](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[float64](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	case fieldDataTypeInt:
		if fb.multiple {
			field := createFormField[[]int](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[int](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	case fieldDataTypeBool:
		if fb.multiple {
			field := createFormField[[]bool](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[bool](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	case fieldDataTypeFile:
		if fb.multiple {
			field := createFormField[[]Multipart](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[Multipart](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	case fieldDataTypeTime:
		if fb.multiple {
			field := createFormField[[]time.Time](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
		if !fb.multiple {
			field := createFormField[time.Time](fb, req)
			formField.Set(reflect.ValueOf(field))
			return field.Messages
		}
	}
	return errors
}

func buildBaseForm(formRef reflect.Value, b *Builder) {
	if formRef.Kind() != reflect.Ptr {
		return
	}
	baseFormField := formRef.Elem().FieldByName(baseFormFieldName)
	if !baseFormField.IsValid() {
		return
	}
	baseFormField.Set(reflect.ValueOf(createBaseForm(b)))
}

func createBaseForm(b *Builder) Form {
	return Form{
		Method:      b.method,
		Action:      b.action,
		ContentType: b.contentType,
		Security:    b.security,
		Valid:       b.isValid(),
		Submitted:   b.submitted,
		Hx:          b.hx,
	}
}

func createFormField[T any](fb *FieldBuilder, req *http.Request) Field[T] {
	return Field[T]{
		Id:        fb.id,
		Name:      fb.name,
		Type:      fb.fieldType,
		DataType:  fb.dataType,
		Label:     fb.label,
		Text:      fb.text,
		Value:     fb.value.(T),
		Multiple:  fb.multiple,
		Messages:  validateField(fb, req),
		Required:  fb.isRequired(),
		Disabled:  fb.disabled,
		Autofocus: fb.autofocus,
	}
}
