# Creamsensation - Form

## Example code
```go
package example

import . "github.com/creamsensation/form"

type ExampleForm struct {
  Form
  Name  Field[string] 
  Age   Field[int]
}

func createExampleForm() (ExampleForm, error) {
  form := New(
    Add("name").With(Text(), Validate.Required()),
    Add("age").With(Number[int](), Validate.Required()),
  )
  return Build[ExampleForm](form)
}
```

## New()
Creates new form builder, which accept field builders
```go
New(fields...)
```

### Builder - Action()
Set form action
```go
formBuilder.Action(action)
```

### Builder - Add()
Same as Add() function, only alternative
```go
formBuilder.Add(name)
```

### Builder - Get()
Get form field
```go
formBuilder.Get(name)
```

### Builder - Limit()
Data limit (MBs)
```go
formBuilder.Limit(limit)
```

### Builder - Method()
Set form method
```go
formBuilder.Method(method)
```

### Builder - Name()
Set form name
```go
formBuilder.Name(name)
```

### Builder - Request()
Provide request to form, it uses native *http.Request
```go
formBuilder.Request(request)
```

## Add()
Creates new field builder
```go
Add(config, validators...)
```

## Validate
### Validate - Required()
Use when form field value is required, it works with string, int, floats, bool and Multipart
```go
Validate.Required()
--
Add("example").With(Text(), Validate.Required())
```
### Validate - Min()
Use when form field value must have minimal value or minimal length, it works with string, int, floats
```go
Validate.Min(1)
--
Add("text").With(Text(), Validate.Min(1))
Add("amount").With(Number[float64](), Validate.Min(1))
```
### Validate - Max()
Use when form field value must have maximum value or maximum length, it works with string, int, floats
```go
Validate.Max(10)
--
Add("text").With(Text(), Validate.Max(10))
Add("amount").With(Number[float64](), Validate.Max(10))
```
### Validate - Email()
Use when form field value must have email pattern, it works with string
```go
Validate.Email()
--
Add("email").With(Email("test@test.cz"), Validate.Email())
```

## Build()
Creates form from form builder, you have to provide result type
```go
Build[ExampleForm](formBuilder)
```

## CreateStruct()
Convert form struct to any data model struct, you have to provide source and result type
```go
CreateStruct[ExampleForm, Model](&form)
```