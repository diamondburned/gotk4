{ systemPkgs ? import <nixpkgs> {} }:

# Pin Nixpkgs for a constant output.
let unstable = import (systemPkgs.fetchFromGitHub {
	owner  = "NixOS";
	repo   = "nixpkgs";
	rev    = "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94";
	sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
}) {};

in systemPkgs.mkShell {
	buildInputs = with unstable; [
		# General GTK dependencies.
		glib
		graphene
		# GTK3.
		gnome3.gtk
		# GTK4 dependencies.
		gtk4
		vulkan-headers
	];

	nativeBuildInputs = with unstable; [
		# Build dependencies.
		gobjectIntrospection
		pkgconfig
		go
	];

	shellHook = ''
		trash() {
			command rm /tmp/gotk4-pkg && \
				mv pkg /tmp/gotk4-pkg && \
				mkdir pkg && \
				echo "Moved pkg/ to /tmp/gotk4-pkg/"
		}
	'';
}
