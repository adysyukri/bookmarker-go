package elem

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func GInput(inputType, label, name, value string) g.Node {
	return html.Div(
		html.Input(
			// html.Type(fmt.Sprintf("{ %v }", inputType)),
			// html.Placeholder(fmt.Sprintf("{ %v }", label)),
			// html.Name(fmt.Sprintf("{ %v }", name)),
			// g.If(value != "", html.Value(fmt.Sprintf("{ %v }", value))),
			// g.If(value == "", g.Attr("x-model", fmt.Sprintf("{ %v }", name))),
			html.Type(inputType),
			html.Placeholder(label),
			html.Name(name),
			g.If(value != "", html.Value(value)),
			g.If(value == "", g.Attr("x-model", name)),
			html.Class("flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md border-neutral-300 ring-offset-background placeholder:text-neutral-500 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50"),
		),
	)
}
