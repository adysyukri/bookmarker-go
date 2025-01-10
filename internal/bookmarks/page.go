package bookmarks

import (
	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/layout"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Page(bml BookmarkList, counter int) g.Node {
	return layout.Layout(
		html.Div(
			html.Class("columns"),
			// left column
			html.Aside(
				html.Class("column is-2"),
				elem.Modal("Add", "Add New Book", "add", html.Form(
					html.Class("box has-background-light"),
					g.Attr("hx-post", "/add"),
					g.Attr("hx-target", "main"),
					g.Attr("hx-swap", "innerHTML show:bottom"),

					html.Script(g.Raw(`
						me().on('htmx:afterRequest', ev => { 
							me(ev).reset()
							me(ev).send('get-count')
						})
					`)),

					elem.Input("text", "Book Title", "title", ""),
					elem.Input("text", "Book Author", "author", ""),
					elem.Input("number", "Pages Total", "total", ""),
					elem.Input("number", "Pages read", "read", ""),

					html.Div(
						html.Class("field is-grouped"),
						html.Div(
							html.Class("control"),
							elem.Button(elem.BtnPrimary, "Save", "submit", g.Group{
								html.Script(g.Raw("me().on('click', ev => { me('.modal').removeClass('is-active') })")),
							}),
						),
						html.Div(
							html.Class("control"),
							elem.Button(elem.BtnNeutral, "Close", "button", g.Group{
								html.Script(g.Raw(`me().on('click', ev => { 
									any('input', me('#modal-add')).run(ev => ev.value = '')
									me('.modal').removeClass('is-active') 
								})`)),
							}),
						),
						html.Div(
							html.Class("control"),
							elem.Button(elem.BtnNeutral, "Reset", "button", g.Group{
								html.Script(g.Raw("me().on('click', ev => { any('input', me('#modal-add')).run(ev => ev.value = '') })")),
							}),
						),
					),
				)),
				BookmarkCounter(counter),
			),
			// right column
			html.Main(
				html.Class("column"),
				ListBookmark(bml),
			),
		),
	)
}
