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
				"gtk4": "/nix/store/niw855nnjgqbq2s0iqxrk9xs5mr10rz8-gtk4-4.2.1-dev/share/gir-1.0",
			},
		},
		{
			[]string{"gtk4", "pango", "cairo", "glib-2.0", "gdk-3.0"},
			map[string]string{
				"gtk4":     "/nix/store/niw855nnjgqbq2s0iqxrk9xs5mr10rz8-gtk4-4.2.1-dev/share/gir-1.0",
				"pango":    "/nix/store/c52730cidby7p2qwwq8cf91anqrni6lg-pango-1.48.4-dev/share/gir-1.0",
				"cairo":    "/nix/store/gp87jysb40b919z8s7ixcilwdsiyl0rp-cairo-1.16.0-dev/share/gir-1.0",
				"glib-2.0": "/nix/store/d9zs9xg86lhqjqni0v8h2ibdrjb57fn4-glib-2.68.2-dev/share/gir-1.0",
				"gdk-3.0":  "/nix/store/vsk1qc1na4izgz461vxkvn655yvarfr7-gtk+3-3.24.27-dev/share/gir-1.0",
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
