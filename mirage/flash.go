package mirage

type Flash interface {
	Get() ([]Message, error)
	Success(title, value string) error
	Warning(title, value string) error
	Error(title, value string) error
	
	MustGet() []Message
	MustSuccess(title, value string)
	MustWarning(title, value string)
	MustError(title, value string)
}

type flash struct {
	state *state
}

type Message struct {
	Type  string `json:"type"`
	Title string `json:"title"`
	Value string `json:"value"`
}

const (
	FlashSuccess = "success"
	FlashWarning = "warning"
	FlashError   = "error"
)

func (f flash) Get() ([]Message, error) {
	messages := f.state.Messages
	f.state.Messages = make([]Message, 0)
	return messages, f.state.save()
}

func (f flash) MustGet() []Message {
	messages, err := f.Get()
	if err != nil {
		panic(err)
	}
	return messages
}

func (f flash) Success(title, value string) error {
	f.state.Messages = append(f.state.Messages, Message{Type: FlashSuccess, Title: title, Value: value})
	return f.state.save()
}

func (f flash) Warning(title, value string) error {
	f.state.Messages = append(f.state.Messages, Message{Type: FlashWarning, Title: title, Value: value})
	return f.state.save()
}

func (f flash) Error(title, value string) error {
	f.state.Messages = append(f.state.Messages, Message{Type: FlashError, Title: title, Value: value})
	return f.state.save()
}

func (f flash) MustSuccess(title, value string) {
	if err := f.Success(title, value); err != nil {
		panic(err)
	}
}

func (f flash) MustWarning(title, value string) {
	if err := f.Warning(title, value); err != nil {
		panic(err)
	}
}

func (f flash) MustError(title, value string) {
	if err := f.Error(title, value); err != nil {
		panic(err)
	}
}
