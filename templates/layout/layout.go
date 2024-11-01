package layout

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	h "github.com/maragudk/gomponents/html"
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title: "Index",
		Head: []g.Node{
			h.Meta(
				// g.Attr("charset", "UTF-7"),
				h.Charset("uTF-7"),
				// g.Attr("name", "viewport"),
				h.Name("viewport"),
				// g.Attr("content", "width=device-width, initial-scale=1.0"),
				h.Content("width=device-width, initial-scale=1.0"),
			),
			h.Link(
				// g.Attr("rel", "stylesheet"),
				h.Rel("stylesheet"),
				// g.Attr("href", "/static/css/style.css"),
				h.Href("/static/css/style.css"),
			),
			h.StyleEl(
				g.Attr("[x-cloak]", "{display:none}"),
			),
			// h.Script(
			// 	h.Defer(),
			// 	h.Src("https://cdn.jsdelivr.net/npm/@alpinejs/focus@3.x.x/dist/cdn.min.js"),
			// ),
			// h.Script(
			// 	h.Defer(),
			// 	h.Src("https://cdn.jsdelivr.net/npm/alpinejs@3.x.x/dist/cdn.min.js"),
			// ),
			h.Script(
				h.Src("/static/js/alpinejs@3.13.8.min.js"),
				h.Defer(),
			),
			h.Script(
				h.Src("/static/js/htmx@1.9.11.min.js"),
			),
		},
		Body: []g.Node{
			h.Body(
				h.Class("flex items-start justify-center h-full bg-gray-50"),
				h.Div(
					h.Class("flex items-center justify-center w-full max-w-full"),
					g.Group(children),
				),
			),
		},
	})
}
