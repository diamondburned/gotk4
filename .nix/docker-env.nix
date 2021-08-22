{ pkgs ? (import ./pkgs.nix {}) }:

let env = pkgs.writeShellScriptBin "docker-env" ''
	set -e
	
	cmd=$1
	shift
	
	case "$cmd" in
	"init")
		${pkgs.nix}/bin/nix-shell --pure --run 'declare -xp > /nix-environment'
		;;
	"exec")
		__user=$USER
		__home=$HOME

		source /nix-environment

		USER=$__user
		HOME=$__home

		eval "$@"
		;;
	esac
'';

in pkgs.buildEnv {
	name  = "gotk4-docker-env";
	paths = [ env ];
}
