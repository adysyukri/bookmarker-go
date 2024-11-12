package bookmarks

import (
	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/layout"
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Page(bml BookmarkList) g.Node {
	inputDefault := `{
		title: "",
		author: "",
		total: null,
		read: null
	}`

	return layout.Layout(
		html.Div(
			html.Class("container mx-auto flex justify-between py-8"),
			// left column
			html.Nav(
				html.Class("w-1/4 bg-gray-200 p-4"),
				elem.GModal("Add", "Add New Book", html.Form(
					g.Attr("x-data", inputDefault),
					g.Attr("hx-post", "/add"),
					g.Attr("hx-target", "main"),
					g.Attr("hx-swap", "beforeend"),
					g.Attr("hx-on:after-request", "this.reset()"),
					elem.GInput("text", "Book Title", "title", ""),
					elem.GInput("text", "Book Author", "author", ""),
					elem.GInput("number", "Pages Total", "total", ""),
					elem.GInput("number", "Pages read", "read", ""),
					elem.GButton(elem.BtnPrimary, "Save", "submit", g.Group{g.Attr("@click", "modalOpen=false")}),
					elem.GButton(elem.BtnNeutral, "Close", "button", g.Group{g.Attr("@click", `() => {
						modalOpen = false,
						title = ""
						author = ""
						total = null
						read = null	
					}`)}),
				)),
			),
			// right column
			html.Main(
				html.Class("w-3/4 bg-white p-4 flex-col"),
				g.Text("This is developed with Gomponents"),
				g.Map(bml, func(bm *Bookmark) g.Node {
					return GBookmarkCard(bm)
				}),
			),
		),
	)
}
