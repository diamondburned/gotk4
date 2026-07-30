package strcases

import (
	"testing"
)

func TestPascalToGo(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// Existing edge cases in replaced.txt
		{"dbus", "DBus"},
		{"gicon", "GIcon"},
		{"gtype", "GType"},
		{"gvalue", "GValue"},
		{"gvariant", "GVariant"},
		{"ipv4", "IPv4"},
		{"ipv6", "IPv6"},
		{"etag", "ETag"},
		{"proxy", "Proxy"},
		{"tai_le", "TaiLe"},
		{"foreach", "ForEach"},

		// Network & Acronyms
		{"ip_tos", "IPToS"},
		{"ipv6_tclass", "IPv6TClass"},
		{"cicp", "CICP"},
		{"ecn", "ECN"},
		{"dscp", "DSCP"},
		{"dnd", "DND"},
		{"uri", "URI"},
		{"uris", "URIs"},
		{"id", "ID"},
		{"ids", "IDs"},
		{"xml", "XML"},
		{"svg", "SVG"},
		{"http", "HTTP"},
		{"https", "HTTPS"},
		{"rgba", "RGBA"},
		{"tls", "TLS"},
		{"dtls", "DTLS"},
		{"tcp", "TCP"},
		{"udp", "UDP"},
		{"uuid", "UUID"},
		{"mime", "MIME"},
		{"hmac", "HMAC"},
	}

	for _, tt := range tests {
		got := Go(tt.input)
		if got != tt.expected {
			t.Errorf("Go(%q) = %q; want %q", tt.input, got, tt.expected)
		}
	}
}
