package tempest

import "fmt"

type Shadow struct {
	Value   string
	Hex     string
	Color   string
	Opacity float64
}

var (
	standardizedBoxShadow = map[string]string{
		SizeSm:   fmt.Sprintf("0 1px 2px 0 var(%s)", shadowColorVar),
		SizeMain: fmt.Sprintf("0 1px 3px 0 var(%[1]s), 0 1px 2px -1px var(%[1]s)", shadowColorVar),
		SizeMd:   fmt.Sprintf("0 4px 6px -1px var(%[1]s), 0 2px 4px -2px var(%[1]s)", shadowColorVar),
		SizeLg:   fmt.Sprintf("0 10px 15px -3px var(%[1]s), 0 4px 6px -4px var(%[1]s)", shadowColorVar),
		SizeXl:   fmt.Sprintf("0 20px 25px -5px var(%[1]s), 0 8px 10px -6px var(%[1]s)", shadowColorVar),
		SizeXxl:  fmt.Sprintf("0 25px 50px -12px var(%s)", shadowColorVar),
		Inner:    fmt.Sprintf("inset 0 2px 4px 0 var(%s)", shadowColorVar),
		None:     "0 0 #0000",
	}
)

var (
	BoxShadow = map[string][]Shadow{
		SizeSm: {
			{Value: "0 1px 2px 0", Opacity: 5},
		},
		SizeMain: {
			{Value: "0 1px 3px 0", Opacity: 10},
			{Value: "0 1px 2px -1px", Opacity: 10},
		},
		SizeMd: {
			{Value: "0 4px 6px -1px", Opacity: 10},
			{Value: "0 2px 4px -2px", Opacity: 10},
		},
		SizeLg: {
			{Value: "0 10px 15px -3px", Opacity: 10},
			{Value: "0 4px 6px -4px", Opacity: 10},
		},
		SizeXl: {
			{Value: "0 20px 25px -5px", Opacity: 10},
			{Value: "0 8px 10px -6px", Opacity: 10},
		},
		SizeXxl: {
			{Value: "0 25px 50px -12px", Opacity: 25},
		},
		Inner: {
			{Value: "inset 0 2px 4px 0", Opacity: 5},
		},
		None: {
			{Value: "0 0"},
		},
	}
)

var (
	DefaultShadowColor = "#000000"
)
