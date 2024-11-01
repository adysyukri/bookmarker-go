package elem

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func GInput(inputType, label, name, value string) g.Node {
	return h.Div(
		h.Input(
			// h.Type(fmt.Sprintf("{ %v }", inputType)),
			// h.Placeholder(fmt.Sprintf("{ %v }", label)),
			// h.Name(fmt.Sprintf("{ %v }", name)),
			// g.If(value != "", h.Value(fmt.Sprintf("{ %v }", value))),
			// g.If(value == "", g.Attr("x-model", fmt.Sprintf("{ %v }", name))),
			h.Type(inputType),
			h.Placeholder(label),
			h.Name(name),
			g.If(value != "", h.Value(value)),
			g.If(value == "", g.Attr("x-model", name)),
			h.Class("flex w-full h-10 px-3 py-2 text-sm bg-white border rounded-md border-neutral-300 ring-offset-background placeholder:text-neutral-500 focus:border-neutral-300 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-neutral-400 disabled:cursor-not-allowed disabled:opacity-50"),
		),
	)
}
