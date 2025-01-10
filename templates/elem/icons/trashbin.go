package icons

import (
	g "maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func TrashBin(attrs g.Group) g.Node {
	return html.SVG(
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		html.Width("25"),
		html.Height("25"),
		g.Attr("fill", "#F9476C"),
		g.Attr("viewBox", "0 0 256 256"),
		g.Map(attrs, func(attr g.Node) g.Node { return attr }),
		g.El("path", g.Attr("d", "M216,48H176V40a24,24,0,0,0-24-24H104A24,24,0,0,0,80,40v8H40a8,8,0,0,0,0,16h8V208a16,16,0,0,0,16,16H192a16,16,0,0,0,16-16V64h8a8,8,0,0,0,0-16ZM112,168a8,8,0,0,1-16,0V104a8,8,0,0,1,16,0Zm48,0a8,8,0,0,1-16,0V104a8,8,0,0,1,16,0Zm0-120H96V40a8,8,0,0,1,8-8h48a8,8,0,0,1,8,8Z")),
	)
}
