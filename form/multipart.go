package form

type Multipart struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Suffix string `json:"suffix"`
	Data   []byte `json:"data"`
}
