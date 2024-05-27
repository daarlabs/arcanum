package mjml

import "github.com/daarlabs/arcanum/gox"

func Align(value ...string) gox.Node {
	return gox.CreateAttribute[string]("align")(value...)
}

func ContainerBackgroundColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("container-background-color")(value...)
}

func BackgroundColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-color")(value...)
}

func BackgroundHeight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-height")(value...)
}

func BackgroundWidth(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-width")(value...)
}

func BackgroundPosition(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-position")(value...)
}

func BackgroundPositionX(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-position-x")(value...)
}

func BackgroundPositionY(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-position-y")(value...)
}

func BackgroundRepeat(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-repeat")(value...)
}

func BackgroundSize(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-size")(value...)
}

func BackgroundUrl(value ...string) gox.Node {
	return gox.CreateAttribute[string]("background-url")(value...)
}

func Border(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border")(value...)
}

func BorderColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-color")(value...)
}

func BorderRadius(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-radius")(value...)
}

func BorderStyle(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-style")(value...)
}

func BorderWidth(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-width")(value...)
}

func BorderTop(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-top")(value...)
}

func BorderRight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-right")(value...)
}

func BorderBottom(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-bottom")(value...)
}

func BorderLeft(value ...string) gox.Node {
	return gox.CreateAttribute[string]("border-left")(value...)
}

func Cellpadding(value ...string) gox.Node {
	return gox.CreateAttribute[string]("cellpadding")(value...)
}

func Cellspacing(value ...string) gox.Node {
	return gox.CreateAttribute[string]("cellspacing")(value...)
}

func Color(value ...string) gox.Node {
	return gox.CreateAttribute[string]("color")(value...)
}

func CssClass(value ...string) gox.Node {
	return gox.CreateAttribute[string]("css-class")(value...)
}

func Direction(value ...string) gox.Node {
	return gox.CreateAttribute[string]("direction")(value...)
}

func FluidOnMobile(value ...any) gox.Node {
	return gox.CreateAttribute[any]("fluid-on-mobile")(value...)
}

func FontFamily(value ...any) gox.Node {
	return gox.CreateAttribute[any]("font-family")(value...)
}

func FontSize(value ...any) gox.Node {
	return gox.CreateAttribute[any]("font-size")(value...)
}

func FontWeight(value ...any) gox.Node {
	return gox.CreateAttribute[any]("font-weight")(value...)
}

func FullWidth(value ...any) gox.Node {
	return gox.CreateAttribute[any]("full-width")(value...)
}

func Hamburger(value ...string) gox.Node {
	return gox.CreateAttribute[string]("hamburger")(value...)
}

func Height(value ...string) gox.Node {
	return gox.CreateAttribute[string]("height")(value...)
}

func IcoAlign(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-align")(value...)
}

func IcoClose(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-close")(value...)
}

func IcoColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-color")(value...)
}

func IcoFontFamily(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-font-family")(value...)
}

func IcoFontSize(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-font-size")(value...)
}

func IcoLineHeight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-line-height")(value...)
}

func IcoOpen(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-open")(value...)
}

func IcoPadding(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-padding")(value...)
}

func IcoPaddingTop(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-padding-top")(value...)
}

func IcoPaddingRight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-padding-right")(value...)
}

func IcoPaddingBottom(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-padding-bottom")(value...)
}

func IcoPaddingLeft(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-padding-left")(value...)
}

func IcoTextDecoration(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-text-decoration")(value...)
}

func IcoTextTransform(value ...string) gox.Node {
	return gox.CreateAttribute[string]("ico-text-transform")(value...)
}

func IconAlign(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-align")(value...)
}

func IconHeight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-height")(value...)
}

func IconPadding(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-padding")(value...)
}

func IconPosition(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-position")(value...)
}

func IconUnwrappedAlt(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-unwrapped-alt")(value...)
}

func IconUnwrappedUrl(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-unwrapped-url")(value...)
}

func IconWidth(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-width")(value...)
}

func IconWrappedAlt(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-wrapped-alt")(value...)
}

func IconWrappedUrl(value ...string) gox.Node {
	return gox.CreateAttribute[string]("icon-wrapped-url")(value...)
}

func InnerBorder(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border")(value...)
}

func InnerBorderTop(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border-top")(value...)
}

func InnerBorderRight(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border-right")(value...)
}

func InnerBorderBottom(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border-bottom")(value...)
}

func InnerBorderLeft(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border-left")(value...)
}

func InnerBorderRadius(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-border-radius")(value...)
}

func InnerPadding(value ...string) gox.Node {
	return gox.CreateAttribute[string]("inner-padding")(value...)
}

func LetterSpacing(value ...any) gox.Node {
	return gox.CreateAttribute[any]("letter-spacing")(value...)
}

func LeftIcon(value ...any) gox.Node {
	return gox.CreateAttribute[any]("left-icon")(value...)
}

func Mode(value ...any) gox.Node {
	return gox.CreateAttribute[any]("mode")(value...)
}

func Padding(value ...any) gox.Node {
	return gox.CreateAttribute[any]("padding")(value...)
}

func PaddingTop(value ...any) gox.Node {
	return gox.CreateAttribute[any]("padding-top")(value...)
}

func PaddingRight(value ...any) gox.Node {
	return gox.CreateAttribute[any]("padding-right")(value...)
}

func PaddingBottom(value ...any) gox.Node {
	return gox.CreateAttribute[any]("padding-bottom")(value...)
}

func PaddingLeft(value ...any) gox.Node {
	return gox.CreateAttribute[any]("padding-left")(value...)
}

func Position(value ...any) gox.Node {
	return gox.CreateAttribute[any]("position")(value...)
}

func RightIcon(value ...any) gox.Node {
	return gox.CreateAttribute[any]("right-icon")(value...)
}

func Sizes(value ...any) gox.Node {
	return gox.CreateAttribute[any]("sizes")(value...)
}

func TableLayout(value ...string) gox.Node {
	return gox.CreateAttribute[string]("table-layout")(value...)
}

func TdBorder(value ...string) gox.Node {
	return gox.CreateAttribute[string]("td-border")(value...)
}

func TdBorderRadius(value ...string) gox.Node {
	return gox.CreateAttribute[string]("td-border-radius")(value...)
}

func TdHoverBorderColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("td-hover-border-color")(value...)
}

func TdSelectedBorderColor(value ...string) gox.Node {
	return gox.CreateAttribute[string]("td-selected-border-color")(value...)
}

func TdWidth(value ...string) gox.Node {
	return gox.CreateAttribute[string]("td-width")(value...)
}

func TextAlign(value ...string) gox.Node {
	return gox.CreateAttribute[string]("text-align")(value...)
}

func TextDecoration(value ...string) gox.Node {
	return gox.CreateAttribute[string]("text-decoration")(value...)
}

func TextTransform(value ...string) gox.Node {
	return gox.CreateAttribute[string]("text-transform")(value...)
}

func VerticalAlign(value ...string) gox.Node {
	return gox.CreateAttribute[string]("vertical-align")(value...)
}

func Width(value ...string) gox.Node {
	return gox.CreateAttribute[string]("width")(value...)
}
