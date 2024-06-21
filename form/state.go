package form

type state interface {
	GetForm(key string) (map[string][]string, error)
	SaveForm(key string, form map[string][]string) error
	
	MustGetForm(key string) map[string][]string
	MustSaveForm(key string, form map[string][]string)
}
