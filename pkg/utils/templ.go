package utils

import (
	"bytes"
	"context"

	"github.com/a-h/templ"
)

func ComponentToString(ctx context.Context, c templ.Component) (string, error) {
	buf := new(bytes.Buffer)

	err := c.Render(ctx, buf)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
