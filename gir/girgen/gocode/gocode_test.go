package gocode

import (
	"fmt"
	"strings"
	"testing"
)

func TestCoalesceTail(t *testing.T) {
	tests := []struct{ in, out string }{
		{"(hello string, world string)", "(hello, world string)"},
		{"(hello, world string)", "(hello, world string)"},
		{"(hello string, world, world2 string)", "(hello, world, world2 string)"},
		{"(hello string, world string, number int, number2 int)", "(hello, world string, number, number2 int)"},
		{"(hello string, number int, world string)", "(hello string, number int, world string)"},
		{"(hello *T, world *T, a string)", "(hello, world *T, a string)"},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			out := CoalesceTail(test.in)
			if out != test.out {
				t.Errorf("expected: %q", test.out)
				t.Errorf("got:      %q", out)
			}
		})
	}
}

func TestExtractDefer(t *testing.T) {
	tests := []struct{ in, out string }{{
		`	ret := functionName()
			defer ret.Destroy()
		`,
		`	ret.Destroy()
		`,
	}, {
		`	file := mustOpen("filename")
			defer file.Close()
			defer otherFunction()
		`,
		`	otherFunction()
			file.Close()
		`,
	},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case%d", i+1), func(t *testing.T) {
			test.out = strings.TrimSpace(test.out)
			test.out = strings.ReplaceAll(test.out, "\t", "")

			out := strings.TrimSpace(ExtractDefer(test.in))
			if out != test.out {
				t.Errorf("expected: %q", test.out)
				t.Errorf("got:      %q", out)
			}
		})
	}
}
