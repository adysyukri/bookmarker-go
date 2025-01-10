package elem

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Card(title string, btn g.Node, icons g.Group, children ...g.Node) g.Node {
	return html.Div(
		html.Class("columns"),
		html.Div(
			html.Class("column is-one-third"),
			html.Div(
				html.Class("card"),
				html.Div(
					html.Class("card-header"),
					html.P(
						html.H5(html.Class("card-header-title"), g.Text(title)),
					),
					html.Div(
						html.Class("card-header-icon"),
						g.Map(icons, func(icon g.Node) g.Node {
							return html.Span(html.Class("icon"), icon)
						}),
					),
				),
				html.Div(
					html.Class("card-content"),
					g.Group(children),
				),
				html.Footer(
					html.Class("card-footer"),
					html.Div(
						html.Class("card-footer-item"),
						btn,
					),
				),
			),
		),
	)
}
