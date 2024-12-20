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
			html.Link(
				html.Rel("stylesheet"),
				html.Href("/static/css/bulma.min.css"),
			),
			html.Script(
				html.Src("/static/js/htmx@1.9.11.min.js"),
			),
			html.Script(
				html.Src("/static/js/surreal@1.3.2.js"),
			),
		},
		Body: []g.Node{
			html.Body(
				html.Class("section"),
				html.Script(g.Raw("me().on('keydown', ev => { if(ev.key === 'Escape') any('.modal').removeClass('is-active') })")),
				html.Div(
					html.Class("container"),
					g.Group(children),
				),
			),
		},
	})
}
