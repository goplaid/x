package pagui_test

import (
	"strings"
	"testing"

	"github.com/theplant/testingutils"

	ui "github.com/sunfmin/pagui"
)

var cases = []struct {
	operation func(b *ui.DefaultPageInjector)
	expected  string
}{
	{
		operation: func(b *ui.DefaultPageInjector) {
			b.Title("Hello")
		},
		expected: `<title>Hello</title>
<meta charset="utf8"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
		`,
	},
	{
		operation: func(b *ui.DefaultPageInjector) {
			b.Title("Hello")
			b.Meta("charset", "shiftjis")
		},
		expected: `<title>Hello</title>
<meta charset="shiftjis"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
`,
	},
	{
		operation: func(b *ui.DefaultPageInjector) {
			b.Title("Hello")
			b.Meta("charset", "shiftjis")
		},
		expected: `<title>Hello</title>
<meta charset="shiftjis"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
`,
	},
	{
		operation: func(b *ui.DefaultPageInjector) {
			b.Title("Hello")
			b.Meta("charset", "shiftjis")
			b.Meta("charset", "utf8")
			b.MetaNameContent("keywords", "Hello")
		},
		expected: `<title>Hello</title>
<meta charset="shiftjis"/>
<meta charset="utf8"/>
<meta name="keywords" content="Hello"/>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no"/>
`,
	},
}

func TestDefaultPageInjector(t *testing.T) {
	for _, c := range cases {
		var b ui.DefaultPageInjector
		c.operation(&b)
		diff := testingutils.PrettyJsonDiff(strings.TrimSpace(c.expected), strings.TrimSpace(b.HeadString()))
		if len(diff) > 0 {
			t.Error(diff)
		}
	}
}
