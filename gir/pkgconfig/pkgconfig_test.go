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
			{"/nix/store/k6wcjrnbnjn4spyiikvm9ysj3b1g2acd-gtk4-4.0.3-dev/include"},
		},
		{
			{"gtk4", "pango", "cairo", "glib-2.0", "gdk-3.0"},
			{
				"/nix/store/k6wcjrnbnjn4spyiikvm9ysj3b1g2acd-gtk4-4.0.3-dev/include",
				"/nix/store/ddkmvr82cx5risya88zhcz49ngfcpbmc-pango-1.48.3-dev/include",
				"/nix/store/0dj85avpdlzrh97fmp7sh02a0lz1z5nv-cairo-1.16.0-dev/include",
				"/nix/store/wkw8b59nljjv649vwf2v05i8qkx5p1ns-glib-2.66.8-dev/include",
				"/nix/store/lzyp6d969dc407n2jbwg2grv28ss5pxn-gtk+3-3.24.27-dev/include",
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
