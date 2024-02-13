{ action }:

let

	systemPkgs = import <nixpkgs> {};

	# gotk4-nix pins the version of
	# <https://github.com/diamondburned/gotk4-nix> to use.
	#
	# gotk4-nix = ../../gotk4-nix;
	gotk4-nix = systemPkgs.fetchFromGitHub {
		owner  = "diamondburned";
		repo   = "gotk4-nix";
		rev    = "2d8319755215bf88730e912297cdfd1a6044645d";
		sha256 = "sha256-QyOZLjn/rY7qpTWSM+5IhdA/Sljrft8I5HyuPMS1Y9I=";
	};

	minGoVersion = "1.21";

in import "${gotk4-nix}/${action}.nix" rec {
	base = {
		pname   = "gotk4";
		version = "dev";
	};
	pkgs = import "${gotk4-nix}/pkgs.nix" {
		# sourceNixpkgs overrides the nixpkgs version pinned in
		# <https://github.com/diamondburned/gotk4-nix/blob/main/src.nix>.
		sourceNixpkgs = systemPkgs.fetchFromGitHub {
			owner = "NixOS";
			repo  = "nixpkgs";
			# This is pinning an older (2021-06-06) revision of nixpkgs.
			rev    = "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94";
			sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
		};
		# useFetched specifies whether to use the system copy of nixpkgs
		# (false), or whether to force use of the pinned version (true).
		useFetched = true;
		# usePatchedGo = true;
		overlays = [
			(self: super:
				let
					pkgs = import "${gotk4-nix}/pkgs.nix" {};
				in
				{
					go =
						with systemPkgs.lib;
						assert assertMsg
							(versionAtLeast pkgs.go.version minGoVersion)
							"go version ${pkgs.go.version} is too old, need at least ${minGoVersion}";
						pkgs.go;
					inherit (pkgs) gopls gotools;
				}
			)
		];
	};
	passthru = {
		inherit pkgs;
	};
}
