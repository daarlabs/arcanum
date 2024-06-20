package devtool

import (
	. "github.com/daarlabs/arcanum/gox"
	"github.com/daarlabs/arcanum/tempest"
)

func layout(assets Node, nodes ...Node) Node {
	return Html(
		Head(
			Title(Text("Recovered error")),
			Meta(
				Name("viewport"),
				Content("width=device-width, initial-scale=1"),
			),
			Raw(
				`
				<link rel="apple-touch-icon" sizes="180x180" href="/public/favicon/apple-touch-icon.png">
				<link rel="icon" type="image/png" sizes="32x32" href="/public/favicon/favicon-32x32.png">
				<link rel="icon" type="image/png" sizes="16x16" href="/public/favicon/favicon-16x16.png">
				<link rel="manifest" href="/public/favicon/site.webmanifest">
				<link rel="mask-icon" href="/public/favicon/safari-pinned-tab.svg" color="#5bbad5">
				<link rel="shortcut icon" href="/public/favicon/favicon.ico">
				<meta name="msapplication-TileColor" content="#00aba9">
				<meta name="msapplication-config" content="/public/favicon/browserconfig.xml">
				<meta name="theme-color" content="#ffffff">
			`,
			),
			assets,
		),
		Body(
			tempest.Class().BgSlate(900).TextWhite().TextXs().Grid().PlaceItemsCenter().
				H("screen").W("screen").Overflow("auto"),
			Div(
				Fragment(nodes...),
			),
		),
	)
}
