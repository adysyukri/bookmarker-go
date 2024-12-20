package bookmarks

import (
	"fmt"
	"strconv"

	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/elem/icons"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func BookmarkCard(bm *Bookmark) g.Node {
	return html.Div(
		html.ID(fmt.Sprintf("card-%s", bm.ID)),
		elem.Card(
			bm.Title,
			elem.Button(
				elem.BtnError,
				"Delete",
				"submit",
				g.Group{
					g.Attr("hx-delete", fmt.Sprintf("/delete/%s", bm.ID)),
					g.Attr("hx-confirm", fmt.Sprintf("Confirm delete %s?", bm.Title)),
					g.Attr("hx-swap", "outerHTML swap:1s show:window:top"),
					g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
					html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
						if(ev.detail.successful) {
							me('#card-%s').fadeOut(undefined, 1000, false)
							await sleep(1000)
							me('.notification').textContent = '%s Deleted'
							me('.notification').className = 'notification is-danger'
							await sleep(3000)
							me('.notification').textContent = '%s'
							me('.notification').removeClass('is-danger')
						}
					})`, bm.ID, bm.Title, notificationContent)),
				},
			),
			g.Group{
				icons.TrashBin(g.Group{
					g.Attr("hx-delete", fmt.Sprintf("/delete/%s", bm.ID)),
					g.Attr("hx-confirm", fmt.Sprintf("Confirm delete %s?", bm.Title)),
					g.Attr("hx-swap", "outerHTML swap:1s show:window:top"),
					g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
					g.Attr("hx-trigger", "click"),
					html.Class("hover:cursor-pointer"),
					html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
						if(ev.detail.successful) {
							me('#card-%s').fadeOut(undefined, 1000, false)
							await sleep(1000)
							me('.notification').textContent = '%s Deleted';
							me('.notification').className = 'notification is-danger';
							await sleep(3000);
							me('.notification').textContent = '%s';
							me('.notification').removeClass('is-danger');
						}
					})`, bm.ID, bm.Title, notificationContent)),
				}),
				icons.EditPencil(g.Group{
					html.Script(g.Rawf("me().on('click', ev => { me('#modal-%s').addClass('is-active') })", bm.ID)),
					html.Class("hover:cursor-pointer"),
				}),
			},
			html.Div(
				html.Class("content"),
				html.P(html.Strong(g.Text("Author: ")), g.Text(bm.Author)),
				html.P(html.Strong(g.Text("Total Page: ")), g.Textf("%v", bm.Total)),
				html.P(html.Strong(g.Text("Total Read: ")), g.Textf("%v", bm.Read)),
			),
		),
		UpdateModal(bm),
	)

}

func BookmarkCounter(counter int) g.Node {
	return html.Div(
		html.ID("counter"),
		g.Attr("hx-swap-oob", "true"),
		html.H3(
			html.Class("subtitle is-6"),
			g.Textf("book count: %v", counter),
		),
	)
}

func UpdateModal(bm *Bookmark) g.Node {
	return elem.Modal("", "Update", bm.ID, html.Form(
		html.ID(fmt.Sprintf("modal-form-%s", bm.ID)),
		g.Attr("hx-on", "click"),
		g.Attr("hx-put", fmt.Sprintf("/edit/%s", bm.ID)),
		g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
		g.Attr("hx-swap", "outerHTML show:bottom"),
		g.Attr("hx-on::after-request", "this.reset()"),

		elem.Input("text", "Book Title", "title", bm.Title),
		elem.Input("text", "Book Author", "author", bm.Author),
		elem.Input("number", "Pages Total", "total", strconv.Itoa(bm.Total)),
		elem.Input("number", "Pages Read", "read", strconv.Itoa(bm.Read)),

		html.Div(
			html.Class("field is-grouped"),
			html.Div(
				html.Class("control"),
				elem.Button(elem.BtnPrimary, "Save", "submit", g.Group{
					html.Script(g.Raw("me().on('click', ev => { any('.modal').removeClass('is-active') })")),
				}),
			),
			html.Div(
				html.Class("control"),
				elem.Button(elem.BtnNeutral, "Close", "button", g.Group{
					html.Script(g.Raw("me().on('click', ev => { any('.modal').removeClass('is-active') })")),
				}),
			),
			html.Div(
				html.Class("control"),
				elem.Button(elem.BtnNeutral, "Reset", "button", g.Group{
					html.Script(g.Rawf("me().on('click', ev => { any('input', me('#modal-form-%s')).run(ev => ev.value = '') })", bm.ID)),
				}),
			),
		),
	),
	)
}
