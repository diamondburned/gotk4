{
	unstable ? (import ./pkgs.nix {}),
	buildInputs ? [],
	nativeBuildInputs ? [],
}:

# minitime is a mini-output time wrapper.
let minitime = unstable.writeShellScriptBin
	"minitime"
	"command time --format $'%C -> %es\\n' \"$@\"";

in unstable.mkShell rec {
	# The build inputs, which contains dependencies needed during generation
	# time, build time and runtime.
	buildInputs = with unstable; [
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers
	] ++ buildInputs;

	# The native build inputs, which contains dependencies only needed during
	# generation time and build time.
	nativeBuildInputs = with unstable; [
		pkgconfig
		go
		minitime
	] ++ nativeBuildInputs;

	CGO_ENABLED = "1";

	# Use /tmp, since /run/user/1000 (XDG_RUNTIME_DIRECTORY) might be too small.
	# See https://github.com/NixOS/nix/issues/395.
	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
