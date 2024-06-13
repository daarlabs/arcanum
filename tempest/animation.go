package tempest

type Animation struct {
	Duration  string
	Timing    string
	Repeat    string
	Delay     string
	Keyframes []Keyframe
	Value     string
}

type Keyframe struct {
	Offset string
	Styles map[string]string
}

var (
	Animations = map[string]Animation{
		"spin": {
			Duration: "1s",
			Timing:   "linear",
			Repeat:   "infinite",
			Keyframes: []Keyframe{
				{Offset: "from", Styles: map[string]string{"transform": "rotate(0deg)"}},
				{Offset: "to", Styles: map[string]string{"transform": "rotate(360deg)"}},
			},
		},
	}
)
