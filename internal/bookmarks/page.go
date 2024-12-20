package bookmarks

import (
	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/layout"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

const (
	notificationContent = "This bookmark is developed with gomponents, bulmacss, surrealjs & htmx"
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
					g.Attr("hx-target", "#bookmark"),
					g.Attr("hx-swap", "beforeend show:bottom"),

					html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
						if(ev.detail.successful) {
							let title = ev.detail.requestConfig.parameters.title
							me(ev).reset()
							me('.notification').textContent = 'Succesfully Added ' + title
							me('.notification').className = 'notification is-success'
							await sleep(3000)
							me('.notification').textContent = '%s'
							me('.notification').removeClass('is-success')
						}
					})`, notificationContent)),

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
				html.Div(
					html.ID("counter"),
					html.H3(
						html.Class("subtitle is-6"),
						g.Textf("book count: %v", counter),
					),
				),
			),
			// right column
			html.Main(
				html.Class("column"),
				html.Div(
					html.Class("notification"),
					html.P(g.Text(notificationContent)),
				),
				html.Div(
					html.ID("bookmark"),
					g.Map(bml, func(bm *Bookmark) g.Node {
						return BookmarkCard(bm)
					}),
				),
			),
		),
	)
}
