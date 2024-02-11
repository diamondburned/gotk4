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
		rev    = "4eab6a0";
		sha256 = "sha256-WsJ2Cf1hvKT3BUYYVxQ5rNMYi6z7NWccbSsw39lgqO8=";
	};

	minGoVersion = 
		with builtins;
		elemAt (elemAt (split "go ([0-9][^\n]*)" (readFile ../go.mod)) 1) 0;

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
			(self: super: {
				go =
					let
						upstreamPkgs = import "${gotk4-nix}/pkgs.nix" {};
						go = upstreamPkgs.go;
					in
						with systemPkgs.lib;
						assert assertMsg
							(versionAtLeast go.version minGoVersion)
							"go version ${go.version} is too old, need at least ${minGoVersion}";

						go;
			})
		];
	};
	passthru = {
		inherit pkgs;
	};
}
