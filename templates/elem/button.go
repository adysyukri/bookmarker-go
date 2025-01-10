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

func Button(btnClass, btnText, btnType string, attrs g.Group) g.Node {
	return html.Button(
		html.Type(btnType),
		c.Classes{
			"button is-outlined": true,
			"is-dark":            btnClass == BtnNeutral,
			"is-primary":         btnClass == BtnPrimary,
			"is-danger":          btnClass == BtnError,
			"is-success":         btnClass == BtnSuccess,
			"is-warning":         btnClass == BtnWarning,
		},
		g.Map(attrs, func(attr g.Node) g.Node { return attr }),
		g.Text(btnText),
	)
}
