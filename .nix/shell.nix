{ pkgs ? (import ./pkgs.nix {}) }:

# minitime is a mini-output time wrapper.
let minitime = pkgs.writeShellScriptBin
	"minitime"
	"command time --format $'%C -> %es\\n' \"$@\"";

in pkgs.mkShell {
	# The build inputs, which contains dependencies needed during generation
	# time, build time and runtime.
	buildInputs = with pkgs; [
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers
	];

	# The native build inputs, which contains dependencies only needed during
	# generation time and build time.
	nativeBuildInputs = with pkgs; [
		pkgconfig
		go
		minitime
	];

	CGO_ENABLED = "1";

	# Use /tmp, since /run/user/1000 (XDG_RUNTIME_DIRECTORY) might be too small.
	# See https://github.com/NixOS/nix/issues/395.
	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
