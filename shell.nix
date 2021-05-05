{ unstable ? import <unstable> {} }:

unstable.mkShell {
	buildInputs = with unstable; [
		# General GTK dependencies.
		glib
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
}
