# The header, where the system's Nixpkgs is imported into. This imported Nixpkgs
# is not guaranteed to be constant on other people's machines, so they should
# only be used to fetch another Nixpkgs.
{ systemPkgs ? import <nixpkgs> {} }:

# The declarations, where a pinned Nixpkgs Unstable is fetched. When updating,
# only rev and sha256 should be changed.
let unstable = import (systemPkgs.fetchFromGitHub {
	owner  = "NixOS";
	repo   = "nixpkgs";
	rev    = "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94";
	sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
}) {};

in unstable.mkShell {
	# The build inputs, which contains dependencies needed during generation
	# time, build time and runtime.
	buildInputs = with unstable; [
		glib
		graphene
		gnome3.gtk
		gtk4
		vulkan-headers
	];

	# The native build inputs, which contains dependencies only needed during
	# generation time and build time.
	nativeBuildInputs = with unstable; [
		gobjectIntrospection
		pkgconfig
		go
		goimports
	];
}
