{ }:

let
	cPackages = [
		"gtk4"
		"gtk3"
	];

	shell = import ./.nix { action = "shell"; };
	pkgs  = shell.pkgs;
in

shell.overrideAttrs (old: {
	shellHook = ''
		${old.shellHook or ""}

		for pkg in ${pkgs.lib.concatStringsSep " " cPackages}; do
			__cpath=$(pkg-config $pkg --cflags-only-I | sed 's/ -I/:/g' | sed 's/^-I//')
			if [[ $__cpath == "" ]]; then
				continue
			fi

			if [[ -z $CPATH ]]; then
				export CPATH=$__cpath
			else
				export CPATH=$CPATH:$__cpath
			fi
		done
	'';
})
