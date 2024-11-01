package elem

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func GButton(btnClass, btnText, btnType string, attrs g.Group) g.Node {
	var class string
	switch btnClass {
	case BtnNeutral:
		class = "inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide transition-colors duration-100 bg-white border-2 rounded-md text-neutral-600 hover:text-white border-neutral-600 hover:bg-neutral-600"
	case BtnPrimary:
		class = "inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-blue-600 transition-colors duration-100 bg-white border-2 border-blue-600 rounded-md hover:text-white hover:bg-blue-600"
	case BtnError:
		class = "inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-red-600 transition-colors duration-100 bg-white border-2 border-red-600 rounded-md hover:text-white hover:bg-red-600"
	case BtnSuccess:
		class = "inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-green-600 transition-colors duration-100 bg-white border-2 border-green-600 rounded-md hover:text-white hover:bg-green-600"
	case BtnWarning:
		class = "inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide text-yellow-600 transition-colors duration-100 bg-white border-2 border-yellow-500 rounded-md hover:text-white hover:bg-yellow-500"
	}

	// var attr g.Node
	// for key, value := range attrs {
	// 	attr = g.Attr(key, value)
	// }

	return h.Button(
		h.Type(btnType),
		h.Class(class),
		g.Map(attrs, func(attr g.Node) g.Node { return attr }),
		g.Text(btnText),
	)
}
