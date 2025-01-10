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
	return g.Group{
		html.Div(
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
						g.Attr("hx-swap", "innerHTML swap:1s show:window:top"),
						g.Attr("hx-target", "main"),
						html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
							if(ev.detail.successful) {
								me('#card-%s').fadeOut(undefined, 1000, false)
								me(ev).send('get-count')
								await sleep(1000)
							}
						})`, bm.ID)),
					},
				),
				g.Group{
					icons.TrashBin(g.Group{
						g.Attr("hx-delete", fmt.Sprintf("/delete/%s", bm.ID)),
						g.Attr("hx-confirm", fmt.Sprintf("Confirm delete %s?", bm.Title)),
						g.Attr("hx-swap", "innerHTML swap:1s show:window:top"),
						g.Attr("hx-target", "main"),
						g.Attr("hx-trigger", "click"),
						html.Class("hover:cursor-pointer"),
						html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
							if(ev.detail.successful) {
								console.log(ev)
								me('#card-%s').fadeOut(undefined, 1000, false)
								me(ev).send('get-count')
								await sleep(1000)
							}
						})`, bm.ID)),
					}),
					icons.EditPencil(g.Group{
						html.Script(g.Rawf(`
							me().on('click', ev => { me('#modal-%s').addClass('is-active') })
						`, bm.ID)),
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
		),
	}
}

func ListBookmark(bml BookmarkList) g.Node {
	fmt.Printf("bml: %#v\n", bml)
	return g.Group{
		g.Iff(len(bml) == 0, func() g.Node {
			return html.Div(
				html.Class("notification"),
				html.P(g.Text("No book to show")),
			)
		}),
		g.Iff(len(bml) != 0, func() g.Node {
			return g.Map(bml, func(bm *Bookmark) g.Node {
				return BookmarkCard(bm)
			})
		}),
	}
}

func BookmarkCounter(counter int) g.Node {
	return html.Div(
		html.ID("counter"),
		g.Attr("hx-get", "/count"),
		g.Attr("hx-trigger", "get-count from:body"),
		g.Attr("hx-swap", "textContent"),
		html.H3(
			html.Class("subtitle is-6"),
			g.Textf("book count: %v", counter),
		),
	)
}

func UpdateModal(bm *Bookmark) g.Node {
	return elem.Modal("", "Update", bm.ID, html.Form(
		html.ID(fmt.Sprintf("modal-form-%s", bm.ID)),
		g.Attr("hx-put", fmt.Sprintf("/edit/%s", bm.ID)),
		g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
		g.Attr("hx-swap", "outerHTML show:bottom"),

		html.Script(g.Rawf(`me().on('htmx:afterRequest', async ev => {
			if(ev.detail.successful) {
				me(ev).reset()
			}
		})`)),

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
