{ action }:

let

	systemPkgs = import <nixpkgs> {};

	# gotk4-nix pins the version of
	# <https://github.com/diamondburned/gotk4-nix> to use.
	#
	# gotk4-nix = ../../gotk4-nix;
	gotk4-nix = systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4-nix";
		# This is a commit from 2023-06-02.
		rev    = "b186ac5419c22f4b75b1bdc35ef4fd9349c6b65e";
		sha256 = "1pfx0p4w56d7pa1l9ipqfq52znfl7slc2wbjfqddq1jv1fp9z43q";
	};

in import "${gotk4-nix}/${action}.nix" {
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
			# This is pinning an older (2021-06-06) revision of nixpkgs than
			# the pinned gotk4-nix version pins (2022-01-29).
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
					in
						upstreamPkgs.go;
			})
		];
	};
}
