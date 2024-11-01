package bookmarks

import (
	"github.com/adysyukri/bookemarker-go/templates/elem"
	"github.com/adysyukri/bookemarker-go/templates/layout"
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Page(bml BookmarkList) g.Node {
	return layout.Layout(
		h.Div(
			h.Class("container mx-auto flex justify-between py-8"),
			// left column
			h.Nav(
				h.Class("w-1/4 bg-gray-200 p-4"),
				elem.GModal("Add", "Add New Book", h.Form(
					g.Attr("x-data", inputDefault),
					g.Attr("hx-post", "/add"),
					g.Attr("hx-target", "main"),
					g.Attr("hx-swap", "beforeend"),
					g.Attr("hx-on:after-request", "this.reset()"),
					elem.GInput("text", "Book Title", "title", ""),
					elem.GInput("text", "Book Author", "author", ""),
					elem.GInput("number", "Pages Total", "total", ""),
					elem.GInput("number", "Pages read", "read", ""),
					// elem.GButton(elem.BtnPrimary, "Save", "submit", map[string]string{"@click": "modalOpen=false"}),
					// elem.GButton(elem.BtnNeutral, "Close", "button", map[string]string{"@click": resetDefault}),
					elem.GButton(elem.BtnPrimary, "Save", "submit", g.Group{g.Attr("@click", "modalOpen=false")}),
					elem.GButton(elem.BtnNeutral, "Close", "button", g.Group{g.Attr("@click", resetDefault)}),
				)),
			),
			// right column
			h.Main(
				h.Class("w-3/4 bg-white p-4 flex-col"),
				g.Text("This is developed with Gomponents"),
				g.Map(bml, func(bm *Bookmark) g.Node {
					return GBookmarkCard(bm)
				}),
			),
		),
	)
}

// func leftColumn() g.Node {
// 	return h.Nav(
// 		h.Class("w-1/4 bg-gray-200 p-4"),
// 		elem.GModal("Add", "Add New Book", h.Form(
// 			g.Attr("x-data", fmt.Sprintf("{ %v }", inputDefault)),
// 			g.Attr("hx-post", "/add"),
// 			g.Attr("hx-target", "main"),
// 			g.Attr("hx-swap", "beforehand"),
// 			g.Attr("hx-on:after-request", "this.reset()"),
// 			elem.GInput("text", "Book Title", "title", ""),
// 			elem.GInput("text", "Book Author", "author", ""),
// 			elem.GInput("number", "Pages Total", "total", ""),
// 			elem.GInput("number", "Pages read", "read", ""),
// 			elem.GButton(elem.BtnPrimary, "Save", "submit", map[string]string{"@click": "modalOpen=false"}),
// 			elem.GButton(elem.BtnNeutral, "Close", "button", map[string]string{"@click": resetDefault}),
// 		)),
// 	)
// }

// func rightColumn(bml BookmarkList) g.Node {
// 	return h.Main(
// 		h.Class("w-3/4 bg-white p-4 flex-col"),
// 		g.Map(bml, func(bm *Bookmark) g.Node {
// 			return GBookmarkCard(bm)
// 		}),
// 	)
// }

// func form() g.Node {
// 	return h.Form(
// 		g.Attr("x-data", fmt.Sprintf("{ %v }", inputDefault)),
// 		g.Attr("hx-post", "/add"),
// 		g.Attr("hx-target", "main"),
// 		g.Attr("hx-swap", "beforehand"),
// 		g.Attr("hx-on:after-request", "this.reset()"),
// 		elem.GInput("text", "Book Title", "title", ""),
// 		elem.GInput("text", "Book Author", "author", ""),
// 		elem.GInput("number", "Pages Total", "total", ""),
// 		elem.GInput("number", "Pages read", "read", ""),
// 		elem.GButton(elem.BtnPrimary, "Save", "submit", map[string]string{"@click": "modalOpen=false"}),
// 		elem.GButton(elem.BtnNeutral, "Close", "button", map[string]string{"@click": resetDefault}),
// 	)
// }
