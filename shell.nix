{ pkgs ? import <unstable> {} }:

pkgs.mkShell {
	buildInputs = with pkgs; [
		# General GTK dependencies.
		glib
		gnome3.gtk

		# GTK4 dependencies.
		gtk4
		vulkan-headers
	];

	nativeBuildInputs = with pkgs; [
		# Build dependencies.
		gobjectIntrospection
		pkgconfig
		go
	];
}
