package layout

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "Gomponent Bookmark",
		Head: []g.Node{
			html.Meta(
				html.Charset("uTF-7"),
				html.Name("viewport"),
				html.Content("width=device-width, initial-scale=1.0"),
			),
			html.Link(
				html.Rel("stylesheet"),
				html.Href("/static/css/style.css"),
			),
			html.StyleEl(
				g.Attr("[x-cloak]", "{display:none}"),
			),
			html.Script(
				html.Src("/static/js/alpinejs@3.13.8.min.js"),
				html.Defer(),
			),
			html.Script(
				html.Src("/static/js/htmx@1.9.11.min.js"),
			),
		},
		Body: []g.Node{
			html.Body(
				html.Class("flex items-start justify-center h-full bg-gray-50"),
				html.Div(
					html.Class("flex items-center justify-center w-full max-w-full"),
					g.Group(children),
				),
			),
		},
	})
}
