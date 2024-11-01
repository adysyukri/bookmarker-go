package elem

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func GModal(btnText, modalTitle string, children ...g.Node) g.Node {
	return h.Div(
		g.Attr("x-data", "{ modalOpen: false }"),
		g.Attr("@keydown.escape.window", "modalOpen = false"),
		h.Class("relative z-50 w-auto h-auto"),
		h.Button(
			g.Attr("@click", "modalOpen=true"),
			h.Class("inline-flex items-center justify-center h-10 px-4 py-2 text-sm font-medium transition-colors bg-white border rounded-md hover:bg-neutral-100 active:bg-white focus:bg-white focus:outline-none focus:ring-2 focus:ring-neutral-200/60 focus:ring-offset-2 disabled:opacity-50 disabled:pointer-events-none"),
			g.Text(btnText),
		),
		g.El(
			"template",
			g.Attr("x-teleport", "body"),
			h.Div(
				g.Attr("x-show", "modalOpen"),
				h.Class("fixed top-0 left-0 z-[99] flex items-center justify-center w-screen h-screen"),
				g.Attr("x-cloak"),
				h.Div(
					g.Attr("x-show", "modalOpen"),
					g.Attr("x-transition:enter", "ease-out duration-300"),
					g.Attr("x-transition:enter-start", "opacity-0"),
					g.Attr("x-transition:enter-end", "opacity-100"),
					g.Attr("x-transition:leave", "ease-in duration-300"),
					g.Attr("x-transition:leave-start", "opacity-100"),
					g.Attr("x-transition:leave-end", "opacity-0"),
					g.Attr("@click", "modalOpen=false"),
					h.Class("absolute inset-0 w-full h-full bg-black bg-opacity-40"),
				),
				h.Div(
					g.Attr("x-show", "modalOpen"),
					g.Attr("x-trap.inert.noscroll", "modalOpen"),
					g.Attr("x-transition:enter", "ease-out duration-300"),
					g.Attr("x-transition:enter-start", "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"),
					g.Attr("x-transition:enter-end", "opacity-100 translate-y-0 sm:scale-100"),
					g.Attr("x-transition:leave", "ease-in duration-300"),
					g.Attr("x-transition:leave-start", "opacity-100 translate-y-0 sm:scale-100"),
					g.Attr("x-transition:leave-end", "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"),
					h.Class("relative w-full py-6 bg-white px-7 sm:max-w-lg sm:rounded-lg"),
					h.Div(
						h.Class("flex items-center justify-between pb-2"),
						h.H3(
							h.Class("text-lg font-semibold"),
							g.Text(modalTitle),
						),
						h.Button(
							g.Attr("@click", "modalOpen=false"),
							h.Class("absolute top-0 right-0 flex items-center justify-center w-8 h-8 mt-5 mr-5 text-gray-600 rounded-full hover:text-gray-800 hover:bg-gray-50"),
							h.SVG(
								h.Class("w-5 h-5"),
								g.Attr("xmlns", "http://www.w3.org/2000/svg"),
								g.Attr("fill", "none"),
								g.Attr("viewBox", "0 0 24 24"),
								g.Attr("stroke-width", "1.5"),
								g.Attr("stroke", "currentColor"),
								g.El(
									"path",
									g.Attr("stroke-linecap", "round"),
									g.Attr("stroke-linejoin", "round"),
									g.Attr("d", "M6 18L18 6M6 6l12 12"),
								),
							),
						),
					),
					h.Div(
						h.Class("relative w-auto"),
						g.Group(children),
					),
				),
			),
		),
	)
}

func GmodalBody(modalTitle string, children ...g.Node) g.Node {
	return g.El(
		"template",
		g.Attr("x-teleport", "body"),
		h.Div(
			g.Attr("x-show", "modalOpen"),
			h.Class("fixed top-0 left-0 z-[99] flex items-center justify-center w-screen h-screen"),
			g.Attr("x-cloak"),
			h.Div(
				g.Attr("x-show", "modalOpen"),
				g.Attr("x-transition:enter", "ease-out duration-300"),
				g.Attr("x-transition:enter-start", "opacity-0"),
				g.Attr("x-transition:enter-end", "opacity-100"),
				g.Attr("x-transition:leave", "ease-in duration-300"),
				g.Attr("x-transition:leave-start", "opacity-100"),
				g.Attr("x-transition:leave-end", "opacity-0"),
				g.Attr("@click", "modalOpen=false"),
				h.Class("absolute inset-0 w-full h-full bg-black bg-opacity-40"),
			),
			h.Div(
				g.Attr("x-show", "modalOpen"),
				g.Attr("x-trap.inert.noscroll", "modalOpen"),
				g.Attr("x-transition:enter", "ease-out duration-300"),
				g.Attr("x-transition:enter-start", "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"),
				g.Attr("x-transition:enter-end", "opacity-100 translate-y-0 sm:scale-100"),
				g.Attr("x-transition:leave", "ease-in duration-300"),
				g.Attr("x-transition:leave-start", "opacity-100 translate-y-0 sm:scale-100"),
				g.Attr("x-transition:leave-end", "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"),
				h.Class("relative w-full py-6 bg-white px-7 sm:max-w-lg sm:rounded-lg"),
				h.Div(
					h.Class("flex items-center justify-between pb-2"),
					h.H3(
						h.Class("text-lg font-semibold"),
						g.Text(modalTitle),
					),
					h.Button(
						g.Attr("@click", "modalOpen=false"),
						h.Class("absolute top-0 right-0 flex items-center justify-center w-8 h-8 mt-5 mr-5 text-gray-600 rounded-full hover:text-gray-800 hover:bg-gray-50"),
						h.SVG(
							h.Class("w-5 h-5"),
							g.Attr("xmlns", "http://www.w3.org/2000/svg"),
							g.Attr("fill", "none"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("stroke-width", "1.5"),
							g.Attr("stroke", "currentColor"),
							g.El(
								"path",
								g.Attr("stroke-linecap", "round"),
								g.Attr("stroke-linejoin", "round"),
								g.Attr("d", "M6 18L18 6M6 6l12 12"),
							),
						),
					),
				),
				h.Div(
					h.Class("relative w-auto"),
					g.Group(children),
				),
			),
		),
	)
}
