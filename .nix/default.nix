{ action }:

let

	systemPkgs = import <nixpkgs> {};

	# gotk4-nix = ../../gotk4-nix;
	gotk4-nix = systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4-nix";
		rev   = "0a50408da4eb59ad4500db49785676714f4bd4e6";
		hash  = "sha256:1ryxxkxly298yr3m0868g69jgk8gagvg6zmpknv2l7jd8x3w4pz7";
	};

in import "${gotk4-nix}/${action}.nix" {
	base = {
		pname   = "gotk4";
		version = "dev";
	};
	pkgs = import "${gotk4-nix}/pkgs.nix" {
		sourceNixpkgs = builtins.fetchTarball {
			name   = "gotk4-nixpkgs";
			url    = "https://github.com/NixOS/nixpkgs/archive/fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94.tar.gz";
			sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
		};
		useFetched = true;
	};
}
