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
}) {
	overlays = [
		(self: super: {
			go = super.go.overrideAttrs (old: {
				version = "1.17beta1";
				src = builtins.fetchurl {
				    url    = "https://dl.google.com/go/go1.17beta1.src.tar.gz";
					sha256 = "1w3m3p7b09cwakjxk7zn28ingjpnr3l7cmc6amcm9g7r4lvrgf02";
				};
				doCheck = false;
			});
			gopls = self.buildGoModule rec {
				pname = "gopls";
				version = "0.7.0";

				src = systemPkgs.fetchgit {
					rev = "gopls/v0.7.0";
					url = "https://go.googlesource.com/tools";
					sha256 = "0vylrsmpszij23yngk7mfysp8rjbf29nyskbrwwysf63r9xbrwbi";
				};

				modRoot = "gopls";
				vendorSha256 = "1mnc84nvl7zhl4pzf90cd0gvid9g1jph6hcxk6lrlnfk2j2m75mj";

				doCheck = false;
				subPackages = [ "." ];
			};
		})
	];
};

in unstable.mkShell {
	# The build inputs, which contains dependencies needed during generation
	# time, build time and runtime.
	buildInputs = with unstable; [
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers
	];

	# The native build inputs, which contains dependencies only needed during
	# generation time and build time.
	nativeBuildInputs =
		with unstable;
		let sh = systemPkgs.writeShellScriptBin;
		in [
			# Build/generation dependencies.
			gobjectIntrospection
			pkgconfig
			go

			# Development tools.
			gopls
			goimports

			# Compile tools.
			zig
			(sh "zcc" ''ZIG_LOCAL_CACHE_DIR=$TMP/z ${unstable.zig}/bin/zig cc  -target x86_64-linux "$@"'')
			(sh "zxx" ''ZIG_LOCAL_CACHE_DIR=$TMP/z ${unstable.zig}/bin/zig c++ -target x86_64-linux "$@"'')
		];
}
