package elem

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Input(inputType, label, name, value string) g.Node {
	return html.Div(
		html.Class("field"),
		html.Div(
			html.Class("control"),
			html.Input(
				html.Class("input"),
				html.Type(inputType),
				html.Placeholder(label),
				html.Name(name),
				html.Value(value),
				// g.If(value == "", g.Attr("x-model", name)),
			),
		),
	)
}
