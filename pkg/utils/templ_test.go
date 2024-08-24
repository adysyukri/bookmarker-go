package utils_test

import (
	"context"
	"testing"

	"github.com/adysyukri/bookemarker-go/pkg/utils"
	"github.com/adysyukri/bookemarker-go/templates/elem"
)

func TestComponentToString(t *testing.T) {
	str, err := utils.ComponentToString(
		context.Background(),
		elem.Input("text", "LabelInput", "input1"),
	)

	if err != nil {
		t.Error("error ComponentToString error not nil: ", err)
	}

	if str == "" {
		t.Error("error ComponentToString string empty")
	}
}

func TestMapToJSObject(t *testing.T) {
	var cardData = []map[string]string{
		{
			"label": "Title",
			"value": "title1",
		},
		{
			"label": "Author",
			"value": "author1",
		},
	}
	expectedValue := `[{Title:"title1"},{Author:"author1"}]`

	result := utils.MapToJSObject(cardData)

	if expectedValue != result {
		t.Errorf("expected\n '%s'\nbut got\n '%s'\n", expectedValue, result)
	}
}
