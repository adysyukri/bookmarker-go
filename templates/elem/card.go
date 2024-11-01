package elem

import (
	"fmt"

	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func GCard(title, id string, btn g.Node, icons g.Group, children ...g.Node) g.Node {
	return h.Div(
		h.ID(fmt.Sprintf("card-%s", id)),
		h.Class("max-w-sm bg-white border rounded-lg shadow-sm p-7 border-neutral-200/60"),
		h.Div(
			h.Class("flex justify-between"),
			h.A(
				h.Href("#_"),
				h.Class("block mb-3"),
				h.H5(h.Class("text-xl font-bold leading-none tracking-tight text-neutral-900"), g.Text(title)),
			),
			h.Div(
				h.Class("flex"),
				g.Map(icons, func(icon g.Node) g.Node {
					return icon
				}),
			),
		),
		h.Div(
			h.Class("mb-4 text-neutral-500"),
			g.Group(children),
		),
		btn,
	)
}
