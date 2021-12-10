{ pkgs ? (import ./pkgs.nix {}) }:

# minitime is a mini-output time wrapper.
let minitime = pkgs.writeShellScriptBin "minitime"
		"command time --format $'%C -> %es\\n' \"$@\"";

	generate = pkgs.writeShellScriptBin "generate"
		"go generate";

	build = pkgs.writeShellScriptBin "build" ''
		cd pkg
		go build -v ./...
	'';

	CFLAGS = "-O1";

in pkgs.mkShell {
	name = "gotk4-shell";

	# Runtime dependencies.
	buildInputs = with pkgs; [
		# Runtime dependencies.
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers
	];

	nativeBuildInputs = with pkgs; [
		# Build dependencies.
		pkgconfig
		go

		# Tools.
		minitime
		generate
		build
	];

	CGO_CFLAGS   = CFLAGS;
	CGO_CXXFLAGS = CFLAGS;
	CGO_FFLAGS   = CFLAGS;
	CGO_LDFLAGS  = CFLAGS;

	CGO_ENABLED = "1";

	# Use /tmp, since /run/user/1000 (XDG_RUNTIME_DIRECTORY) might be too small.
	# See https://github.com/NixOS/nix/issues/395.
	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
