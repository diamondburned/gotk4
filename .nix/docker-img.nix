{ pkgs ? (import ./pkgs.nix {}) }:

let shell = (import ./shell.nix { inherit pkgs; });

	deps = with pkgs; [
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers

		pkgconfig
		go

		bash
		coreutils
	];

	envVars = pkgs.runCommand "gotk4-env" {
		buildInputs = deps;
	} ''
		mkdir -p $out/bin/
		declare -xp > $out/bin/gotk4-env
	'';

in pkgs.dockerTools.buildImage {
	name = "gotk4";
	contents =
		# (with shell.inputDerivation; (buildInputs ++ nativeBuildInputs)) ++
		(with pkgs; [
			gobjectIntrospection
			glib
			graphene
			gdk-pixbuf
			gnome3.gtk
			gtk4
			vulkan-headers

			envVars
			gnugrep
	
			pkgconfig
			go

			ncdu
			bash
			coreutils
		]);
	extraCommands = ''
		mkdir -p /tmp
		chmod -R 777 /tmp
	'';
	config = {
		Env = [
			"TMP=/tmp"
		];
	};
}
