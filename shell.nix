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
					url    = "https://golang.org/dl/go1.17rc1.linux-arm64.tar.gz";
					sha256 = "sha256:0kps5kw9yymxawf57ps9xivqrkx2p60bpmkisahr8jl1rqkf963l";
				};
				# postInstall = (old.postInstall or "") + ''
				# 	echo "Building std as shared library"
				# 	$out/bin/go install -buildmode=shared -linkshared std
				# '';
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
			ccache = super.ccache.overrideAttrs (old: {
				version = "tip-f2f9993";
				doCheck = false;

				buildInputs = (old.nativeBuildInputs) ++ (with super; [
					pkgconfig
					zstd
					hiredis
				]);

				src = systemPkgs.fetchFromGitHub {
					owner = "ccache";
					repo  = "ccache";
					rev   = "f2f9993db6042de6e5f5b55dd8bb4dc5987cf210";
					hash  = "sha256:0f7kkbyk6hi3cbhlapaxk1km6x8jakf61493c1bidr807z25j1vz";
				};
			});
		})
	];
};

in unstable.mkShell.override { stdenv = unstable.ccacheStdenv; } {
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
			ccache
			gopls
			goimports

			# minitime is a mini-output time wrapper.
			(sh "minitime" "command time --format $'%C -> %es\\n' \"$@\"")
			# Alias for gocopy.
			# (sh "gocopy" "exec $TMP/gocopy/${unstable.go.version}/go/bin/go \"$@\"")
		];

	# shellHook = ''
	# 	[[ ! -d $TMP/gocopy/${unstable.go.version} ]] && {
	# 		rm -rf $TMP/gocopy/${unstable.go.version}
	# 		mkdir -p $TMP/gocopy/${unstable.go.version}
	# 		cp -rf ${unstable.go}/share/go $TMP/gocopy/${unstable.go.version}
	# 		chmod u+wx -R $TMP/gocopy/${unstable.go.version}
	# 	}
	# '';

	CGO_ENABLED = "1";

	# Use /tmp, since /run/user/1000 (XDG_RUNTIME_DIRECTORY) might be too small.
	# See https://github.com/NixOS/nix/issues/395.
	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
