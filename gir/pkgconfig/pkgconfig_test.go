package pkgconfig

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIncludeDirs(t *testing.T) {
	tests := [][2][]string{
		{
			{"gtk4"},
			{"/nix/store/niw855nnjgqbq2s0iqxrk9xs5mr10rz8-gtk4-4.2.1-dev/include"},
		},
		{
			{"gtk4", "pango", "cairo", "glib-2.0", "gdk-3.0"},
			{
				"/nix/store/niw855nnjgqbq2s0iqxrk9xs5mr10rz8-gtk4-4.2.1-dev/include",
				"/nix/store/c52730cidby7p2qwwq8cf91anqrni6lg-pango-1.48.4-dev/include",
				"/nix/store/gp87jysb40b919z8s7ixcilwdsiyl0rp-cairo-1.16.0-dev/include",
				"/nix/store/d9zs9xg86lhqjqni0v8h2ibdrjb57fn4-glib-2.68.2-dev/include",
				"/nix/store/vsk1qc1na4izgz461vxkvn655yvarfr7-gtk+3-3.24.27-dev/include",
			},
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("case_%d", i), func(t *testing.T) {
			out, err := IncludeDirs(test[0]...)
			if err != nil {
				t.Fatal("unexpected error:", err)
			}

			if !reflect.DeepEqual(out, test[1]) {
				t.Fatalf("unexpected output\n"+
					"expected %v\n"+
					"got      %v",
					test[1], out,
				)
			}
		})
	}
}
