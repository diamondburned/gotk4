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
		rev    = "ad91dabf706946c4380d0a105f0937e4e8ffd75f";
		sha256 = "0rkw9k98qy7ifwypkh2fqhdn7y2qphy2f8xjisj0cyp5pjja62im";
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
					in
						upstreamPkgs.go;
			})
		];
	};
}
