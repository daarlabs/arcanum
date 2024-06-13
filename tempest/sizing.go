package tempest

const (
	DefaultFontSize = 16
)

var (
	DefaultBreakpoints = map[string]string{
		SizeXs:  "320px",
		SizeSm:  "576px",
		SizeMd:  "768px",
		SizeLg:  "992px",
		SizeXl:  "1200px",
		SizeXxl: "1400px",
	}
	DefaultContainer = map[string]string{
		None:    "100%",
		SizeSm:  "640px",
		SizeMd:  "768px",
		SizeLg:  "1024px",
		SizeXl:  "1280px",
		SizeXxl: "1536px",
	}
)

const (
	Rem = "rem"
	Em  = "em"
	Px  = "px"
	Pct = "%"
	Vw  = "vw"
	Vh  = "vh"
	Deg = "deg"
)

const (
	Height = "height"
	Width  = "width"
	Full   = "full"
	Screen = "screen"
	Auto   = "auto"
)

const (
	SizeXs   = "xs"
	SizeSm   = "sm"
	SizeMd   = "md"
	SizeMain = "main"
	SizeLg   = "lg"
	SizeXl   = "xl"
	SizeXxl  = "xxl"
)

var (
	standardizedSize = map[string]string{
		SizeXs:   "0.75rem",
		SizeSm:   "0.875rem",
		SizeMain: "1rem",
		SizeLg:   "1.125rem",
		SizeXl:   "1.25rem",
		SizeXxl:  "1.5rem",
	}
)

var (
	standardizedLineHeight = map[string]string{
		SizeXs:   "1rem",
		SizeSm:   "1.25rem",
		SizeMain: "1.5rem",
		SizeLg:   "1.75rem",
		SizeXl:   "2rem",
		SizeXxl:  "2.25rem",
	}
)

var (
	standardizedRadius = map[string]string{
		Full:     "100%",
		None:     "0",
		SizeSm:   "0.125rem",
		SizeMain: "0.25rem",
		SizeLg:   "0.375rem",
		SizeXl:   "0.5rem",
		SizeXxl:  "0.75rem",
	}
)
