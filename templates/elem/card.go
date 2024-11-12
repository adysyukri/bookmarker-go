package elem

import (
	"fmt"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func GCard(title, id string, btn g.Node, icons g.Group, children ...g.Node) g.Node {
	return html.Div(
		html.ID(fmt.Sprintf("card-%s", id)),
		html.Class("max-w-sm bg-white border rounded-lg shadow-sm p-7 border-neutral-200/60"),
		html.Div(
			html.Class("flex justify-between"),
			html.A(
				html.Href("#_"),
				html.Class("block mb-3"),
				html.H5(html.Class("text-xl font-bold leading-none tracking-tight text-neutral-900"), g.Text(title)),
			),
			html.Div(
				html.Class("flex"),
				g.Map(icons, func(icon g.Node) g.Node {
					return icon
				}),
			),
		),
		html.Div(
			html.Class("mb-4 text-neutral-500"),
			g.Group(children),
		),
		btn,
	)
}
