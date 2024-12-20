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
			"button":                 true,
			"":                       btnClass == BtnNeutral,
			"is-primary is-outlined": btnClass == BtnPrimary,
			"is-error is-outlined":   btnClass == BtnError,
			"is-success is-outlined": btnClass == BtnSuccess,
			"is-warning is-outlined": btnClass == BtnWarning,
		},
		g.Map(attrs, func(attr g.Node) g.Node { return attr }),
		g.Text(btnText),
	)
}
