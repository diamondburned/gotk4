package pkgconfig

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestParseValues(t *testing.T) {
	tests := map[string]struct {
		inPkgs   []string
		inStdout string
		expVals  map[string]string
	}{
		"missing-leading-fdo": {
			[]string{"gtk4", "guile-3.0", "ruby-3.0"},
			"/usr/share/guile/site/3.0 /usr/lib/ruby/site_ruby\n",
			map[string]string{
				"gtk4":      "",
				"guile-3.0": "/usr/share/guile/site/3.0",
				"ruby-3.0":  "/usr/lib/ruby/site_ruby",
			},
		},
		"missing-leading-pkgconf1.8": {
			[]string{"gtk4", "guile-3.0", "ruby-3.0"},
			" /usr/share/guile/site/3.0 /usr/lib/ruby/site_ruby\n",
			map[string]string{
				"gtk4":      "",
				"guile-3.0": "/usr/share/guile/site/3.0",
				"ruby-3.0":  "/usr/lib/ruby/site_ruby",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out, err := parseValues(test.inPkgs, nil, test.inStdout)
			assert.NoError(t, err)

			assert.Equal(t, test.expVals, out)
		})
	}
}

func TestGIRDirs(t *testing.T) {
	tests := []struct {
		inPkgs     []string
		expGIRDirs map[string]string
	}{
		{
			[]string{"gtk4"},
			map[string]string{
				"gtk4": "/nix/store/j2vgh4x3wmxrycvwkbnp74vh32yip9i5-gtk4-4.14.4-dev/share/gir-1.0",
			},
		},
		{
			[]string{"gtk4", "pango", "cairo", "glib-2.0", "gdk-3.0"},
			map[string]string{
				"gtk4":     "/nix/store/j2vgh4x3wmxrycvwkbnp74vh32yip9i5-gtk4-4.14.4-dev/share/gir-1.0",
				"pango":    "/nix/store/fmqv9zpvf9fkb8n384ramhhbp8k20rnw-pango-1.52.2-dev/share/gir-1.0",
				"cairo":    "/nix/store/0i4xfw9ylgmg8f7z6rwwmfbin3ckyli5-cairo-1.18.0-dev/share/gir-1.0",
				"glib-2.0": "/nix/store/kd2a4kyia0lai7n7hn657mpcr2gykw00-glib-2.80.2-dev/share/gir-1.0",
				"gdk-3.0":  "/nix/store/kd2a4kyia0lai7n7hn657mpcr2gykw00-glib-2.80.2-dev/share/gir-1.0",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			out, err := GIRDirs(test.inPkgs...)
			assert.NoError(t, err)

			assert.Equal(t, test.expGIRDirs, out)
		})
	}
}
