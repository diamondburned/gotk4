{ }:

let
	shell = import ./.nix { action = "shell"; };
	pkgs  = shell.pkgs;
in

shell
