package elem

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	"maragu.dev/gomponents/html"
)

const (
	BtnNeutral = "btn-neutral"
	BtnPrimary = "btn-primary"
	BtnError   = "btn-error"
	BtnSuccess = "btn-success"
	BtnWarning = "btn-warning"
)

func GButton(btnClass, btnText, btnType string, attrs g.Group) g.Node {
	return html.Button(
		html.Type(btnType),
		c.Classes{
			"inline-flex items-center justify-center px-4 py-2 text-sm font-medium tracking-wide":                                                   true,
			"transition-colors duration-100 bg-white border-2 rounded-md text-neutral-600 hover:text-white border-neutral-600 hover:bg-neutral-600": btnClass == BtnNeutral,
			"text-blue-600 transition-colors duration-100 bg-white border-2 border-blue-600 rounded-md hover:text-white hover:bg-blue-600":          btnClass == BtnPrimary,
			"text-red-600 transition-colors duration-100 bg-white border-2 border-red-600 rounded-md hover:text-white hover:bg-red-600":             btnClass == BtnError,
			"text-green-600 transition-colors duration-100 bg-white border-2 border-green-600 rounded-md hover:text-white hover:bg-green-600":       btnClass == BtnSuccess,
			"text-yellow-600 transition-colors duration-100 bg-white border-2 border-yellow-500 rounded-md hover:text-white hover:bg-yellow-500":    btnClass == BtnWarning,
		},
		g.Map(attrs, func(attr g.Node) g.Node { return attr }),
		g.Text(btnText),
	)
}
