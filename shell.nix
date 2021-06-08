{ pkgs ? import <unstable> {} }:

pkgs.mkShell {
	buildInputs = with pkgs; [
		# General GTK dependencies.
		glib
		graphene
		# GTK3.
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

	shellHook = ''
		trash() {
			command rm /tmp/gotk4-pkg && \
				mv pkg /tmp/gotk4-pkg && \
				mkdir pkg && \
				echo "Moved pkg/ to /tmp/gotk4-pkg/"
		}
	'';
}
