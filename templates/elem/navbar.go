package elem

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func Navbar() g.Node {
	return h.Nav(h.Class("flex items-center justify-between flex-wrap bg-teal-500 p-6"),
		h.Ol(
			navbarLink("/home", "Home"),
			navbarLink("/about", "About"),
		),
	)
}

func navbarLink(href, text string) g.Node {
	return h.Ul(
		h.Class("flex"),
		h.Li(
			h.Class("mr-6"),
			h.A(
				h.Class("text-blue-500 hover:text-blue-800"),
				h.Href(href),
				g.Text(text),
			),
		),
	)
}
