package utils

import (
	"bytes"
	"context"
	"fmt"
	"strings"

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

func MapToJSObject(s []map[string]string) string {
	jsv := []string{}
	for _, sv := range s {
		jsv = append(jsv, fmt.Sprintf("{%s:\"%s\"}", sv["label"], sv["value"]))

	}
	return fmt.Sprintf("[%s]", strings.Join(jsv, ","))
}

func XData(s []map[string]string) string {
	return fmt.Sprintf("{data:%s}", MapToJSObject(s))
}
