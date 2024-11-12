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
		g.Attr("x-data", "{modalOpen: false}"),
		elem.GCard(bm.Title, bm.ID, elem.GButton(
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
					g.Attr("@click", "modalOpen=true"),
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
		html.Class("modalinput"),
		elem.GmodalBody("Update", html.Form(
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
				g.Attr("@click", "modalOpen=false"),
			}),
			elem.GButton(elem.BtnNeutral, "Close", "button", g.Group{
				g.Attr("@click", "modalOpen=false"),
			}),
		)),
	)
}
