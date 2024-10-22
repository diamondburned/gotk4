{
	inputs = {
		nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
		nixpkgs-gotk4.url = "github:NixOS/nixpkgs/fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94"; # 2021-06-06
		gotk4-nix.url = "github:diamondburned/gotk4-nix/main";
		gotk4-nix.flake = false;
		flake-utils.url = "github:numtide/flake-utils";
		flake-compat.url = "https://flakehub.com/f/edolstra/flake-compat/1.tar.gz";
	};

	outputs =
		{
			self,
			nixpkgs,
			nixpkgs-gotk4,
			gotk4-nix,
			flake-utils,
			flake-compat,
		}:

		flake-utils.lib.eachDefaultSystem (system:
			let
				pkgs = nixpkgs.legacyPackages.${system};
				cmds = with pkgs; [
					(writeShellScriptBin
						"gir-generate"
						"go run github.com/diamondburned/gotk4/gir/cmd/gir-generate \"$@\"")
				];
			in
			{
				devShells.default = import "${gotk4-nix}/shell.nix" {
					base.pname = "gotk4";
					pkgs = import nixpkgs-gotk4 {
						inherit system;
						overlays = [
							(import "${gotk4-nix}/overlay.nix")
							(import "${gotk4-nix}/overlay-patchelf.nix")
							(self: super: {
								inherit (pkgs)
									go
									gopls
									gotools;
							})
						];
					};
					buildInputs = _: ([] ++ cmds);
				};
				packages.dockerEnv = let
					env = pkgs.writeShellScriptBin "docker-env" ''
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
				in
				pkgs.buildEnv {
					name  = "gotk4-docker-env";
					paths = [ env ] ++ (with pkgs.stdenv; [
						cc
						shellPackage
					]);
				};
			}
		);
}
