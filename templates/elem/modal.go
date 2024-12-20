package elem

import (
	"fmt"

	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Modal(btnText, modalTitle, id string, children ...g.Node) g.Node {
	return html.Div(
		g.If(btnText != "",
			// html.Button(
			// 	html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').classAdd('is-active') })", id))),
			// 	html.Class("button is-light"),
			// 	g.Text(btnText),
			// ),
			Button(BtnNeutral, btnText, "button", g.Group{
				html.Script(g.Rawf("me().on('click', ev => { me('#modal-%s').classAdd('is-active') })", id)),
				// html.Style("position: fixed"),
			}),
		),
		html.Div(
			html.ID(fmt.Sprintf("modal-%s", id)),
			html.Class("modal"),
			html.Div(
				html.Class("modal-background"),
				html.Script(g.Raw("me().on('click', ev => { any('.modal').removeClass('is-active') })")),
			),
			html.Div(
				html.Class("modal-card"),
				html.Header(
					html.Class("modal-card-head"),
					html.H3(
						html.Class("modal-card-title"),
						g.Text(modalTitle),
					),
					html.Button(
						html.Script(g.Raw("me().on('click', ev => { any('.modal').removeClass('is-active') })")),
						html.Class("delete"),
						html.SVG(
							g.Attr("xmlns", "http://www.w3.org/2000/svg"),
							g.Attr("fill", "none"),
							g.Attr("viewBox", "0 0 24 24"),
							g.Attr("stroke-width", "1.5"),
							g.Attr("stroke", "currentColor"),
							g.El(
								"path",
								g.Attr("stroke-linecap", "round"),
								g.Attr("stroke-linejoin", "round"),
								g.Attr("d", "M6 18L18 6M6 6l12 12"),
							),
						),
					),
				),
				html.Section(
					html.Class("modal-card-body"),
					g.Group(children),
				),
			),
		),
	)
}

// func ModalBody(modalTitle, id string, children ...g.Node) g.Node {
// 	return html.Div(
// 		html.ID(fmt.Sprintf("modal-%s", id)),
// 		html.Class("modal"),
// 		html.Div(
// 			html.Class("modal-background"),
// 			html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').removeClass('is-active') })", id))),
// 		),
// 		html.Div(
// 			html.Class("modal-card"),
// 			html.Header(
// 				html.Class("modal-card-head"),
// 				html.H3(
// 					html.Class("modal-card-title"),
// 					g.Text(modalTitle),
// 				),
// 				html.Button(
// 					html.Class("delete"),
// 					html.Script(g.Raw(fmt.Sprintf("me().on('click', ev => { me('#modal-%s').removeClass('is-active') })", id))),
// 					html.SVG(
// 						g.Attr("xmlns", "http://www.w3.org/2000/svg"),
// 						g.Attr("fill", "none"),
// 						g.Attr("viewBox", "0 0 24 24"),
// 						g.Attr("stroke-width", "1.5"),
// 						g.Attr("stroke", "currentColor"),
// 						g.El(
// 							"path",
// 							g.Attr("stroke-linecap", "round"),
// 							g.Attr("stroke-linejoin", "round"),
// 							g.Attr("d", "M6 18L18 6M6 6l12 12"),
// 						),
// 					),
// 				),
// 			),
// 			html.Section(
// 				html.Class("modal-card-body"),
// 				g.Group(children),
// 			),
// 		),
// 	)
// }
