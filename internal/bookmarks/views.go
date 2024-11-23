package bookmarks

import (
	"fmt"
	"strconv"

	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/elem/icons"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func GBookmarkCard(bm *Bookmark) g.Node {
	return html.Div(
		html.Class("bookmarkcard"),
		elem.GCard(
			bm.Title,
			bm.ID,
			elem.GButton(
				elem.BtnError,
				"Delete",
				"submit",
				g.Group{
					g.Attr("hx-delete", fmt.Sprintf("/delete/%s", bm.ID)),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
				},
			),
			g.Group{
				icons.GTrashBin(g.Group{
					g.Attr("hx-delete", fmt.Sprintf("/delete/%s", bm.ID)),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
					g.Attr("hx-trigger", "click"),
					html.Class("hover:cursor-pointer"),
				}),
				icons.GEditPencil(g.Group{
					html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').attribute({'style': null}) })", bm.ID))),
					html.Class("hover:cursor-pointer"),
				}),
			},
			html.P(html.Strong(g.Text("Author: ")), g.Text(bm.Author)),
			html.P(html.Strong(g.Text("Total Page: ")), g.Textf("%v", bm.Total)),
			html.P(html.Strong(g.Text("Total Read: ")), g.Textf("%v", bm.Read)),
		),
		GModalInput(bm),
	)
}

func GModalInput(bm *Bookmark) g.Node {
	return html.Div(
		html.ID(fmt.Sprintf("modal-%s", bm.ID)),
		html.Class("modalinput fixed top-0 left-0 z-[99] flex items-center justify-center w-screen h-screen"),
		html.Style("display: none"),
		html.Script(g.Raw("me().on('closeForm', ev => { halt(ev); me(ev).style.display = 'none' })")),
		elem.SurrealModalBody("Update", html.Form(
			html.ID(fmt.Sprintf("modal-form-%s", bm.ID)),
			g.Attr("hx-on", "click"),
			g.Attr("hx-put", fmt.Sprintf("/edit/%s", bm.ID)),
			g.Attr("hx-target", fmt.Sprintf("#card-%s", bm.ID)),
			g.Attr("hx-swap", "outerHTML"),
			g.Attr("hx-on::after-request", "this.reset()"),

			elem.GInput("text", "Book Title", "title", bm.Title),
			elem.GInput("text", "Book Author", "author", bm.Author),
			elem.GInput("number", "Pages Total", "total", strconv.Itoa(bm.Total)),
			elem.GInput("number", "Pages Read", "read", strconv.Itoa(bm.Read)),

			elem.GButton(elem.BtnPrimary, "Save", "submit", g.Group{
				html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').style.display = 'none' })", bm.ID))),
			}),
			elem.GButton(elem.BtnNeutral, "Close", "button", g.Group{
				html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').style.display = 'none' })", bm.ID))),
			}),
			elem.GButton(elem.BtnNeutral, "Reset", "button", g.Group{
				html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { any('input', me('#modal-form-%s')).run(ev => ev.value = '') })", bm.ID))),
			}),
		)),
	)
}
