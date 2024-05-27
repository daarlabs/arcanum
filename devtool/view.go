package devtool

import (
	"fmt"
	"sort"
	
	. "github.com/daarlabs/arcanum/gox"
)

const (
	Selector = "__devtool__"
)

func createView(
	style, script, renderDuration string, values []any, queries []Query, session map[string]any, route map[string]any,
) Node {
	routeKeys := make([]string, 0)
	sessionKeys := make([]string, 0)
	for key := range route {
		routeKeys = append(routeKeys, key)
	}
	for key := range session {
		sessionKeys = append(sessionKeys, key)
	}
	sort.Strings(routeKeys)
	sort.Strings(sessionKeys)
	
	return Div(
		Id(Selector),
		CreateAttribute[string]("hx-swap-oob")("true"),
		Style(Element(), Text(style)),
		Script(Text(script)),
		Div(
			Class("devtool"),
			Style(Text("display:none;")),
			Div(Class("devtool-infobox is-ok"), Text(renderDuration), CustomData("devtool-app-build")),
			Div(
				Class("devtool-control"),
				If(
					len(route) > 0,
					Div(
						Button(
							Class("devtool-button"),
							Type("button"),
							CustomData("devtool-popup-handler", "route"),
							iconRoute(),
						),
					),
				),
				If(
					len(session) > 0,
					Div(
						Button(
							Class("devtool-button"),
							Type("button"),
							CustomData("devtool-popup-handler", "session"),
							iconSession(),
						),
					),
				),
				If(
					len(queries) > 0,
					Div(
						Button(
							Class("devtool-button"),
							Type("button"),
							CustomData("devtool-popup-handler", "queries"),
							iconDb(),
						),
					),
				),
				If(
					len(values) > 0,
					Div(
						Button(
							Class("devtool-button"),
							Type("button"),
							CustomData("devtool-popup-handler", "logs"),
							iconEye(),
						),
					),
				),
			),
			// Route
			If(
				len(route) > 0,
				Div(
					Class("devtool-popup"),
					CustomData("devtool-popup", "route"),
					Range(
						routeKeys, func(key string, _ int) Node {
							var el Node
							switch route[key].(type) {
							case string:
								el = Span(Class("is-text"), Text(fmt.Sprintf("\"%s\"", route[key])))
							case float64, float32:
								el = Span(Class("is-number"), Text(fmt.Sprintf("%f", route[key])))
							case int:
								el = Span(Class("is-number"), Text(fmt.Sprintf("%d", route[key])))
							case bool:
								el = Span(Class("is-bool"), Text(fmt.Sprintf("%t", route[key])))
							default:
								el = Span(Text(fmt.Sprintf("%+v", route[key])))
							}
							return Fragment(
								Div(
									Class("devtool-popup-value"),
									Span(Text(key+": ")),
									el,
								),
								Div(Class("devtool-popup-divider")),
							)
						},
					),
				),
			),
			// Session
			If(
				len(session) > 0,
				Div(
					Class("devtool-popup"),
					CustomData("devtool-popup", "session"),
					Range(
						sessionKeys, func(key string, _ int) Node {
							var el Node
							switch session[key].(type) {
							case string:
								el = Span(Class("is-text"), Text(fmt.Sprintf("\"%s\"", session[key])))
							case float64, float32:
								el = Span(Class("is-number"), Text(fmt.Sprintf("%f", session[key])))
							case int:
								el = Span(Class("is-number"), Text(fmt.Sprintf("%d", session[key])))
							case bool:
								el = Span(Class("is-bool"), Text(fmt.Sprintf("%t", session[key])))
							default:
								el = Span(Text(fmt.Sprintf("%+v", session[key])))
							}
							return Fragment(
								Div(
									Class("devtool-popup-value"),
									Span(Text(key+": ")),
									el,
								),
								Div(Class("devtool-popup-divider")),
							)
						},
					),
				),
			),
			// Queries popup
			If(
				len(queries) > 0,
				Div(
					Class("devtool-popup"),
					Range(
						queries, func(item Query, i int) Node {
							return Fragment(
								Div(
									Class("devtool-popup-value"),
									Span(Class("is-number"), Text(fmt.Sprintf("[%s]", item.Duration.String()))),
									Span(Class("is-text"), Text(item.Value)),
								),
								Div(Class("devtool-popup-divider")),
							)
						},
					),
					CustomData("devtool-popup", "queries"),
				),
			),
			// Logs popup
			If(
				len(values) > 0,
				Div(
					Class("devtool-popup"),
					Range(
						values, func(item any, i int) Node {
							var el Node
							switch item.(type) {
							case string:
								el = Div(Class("devtool-popup-value is-text"), Text(fmt.Sprintf("\"%s\"", item)))
							case float64, float32:
								el = Div(Class("devtool-popup-value is-number"), Text(fmt.Sprintf("%f", item)))
							case int:
								el = Div(Class("devtool-popup-value is-number"), Text(fmt.Sprintf("%d", item)))
							case bool:
								el = Div(Class("devtool-popup-value is-bool"), Text(fmt.Sprintf("%t", item)))
							default:
								el = Div(Class("devtool-popup-value"), Text(fmt.Sprintf("%+v", item)))
							}
							return Fragment(
								el,
								Div(Class("devtool-popup-divider")),
							)
						},
					),
					CustomData("devtool-popup", "logs"),
				),
			),
		),
	)
}
