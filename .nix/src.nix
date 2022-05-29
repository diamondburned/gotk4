let systemPkgs = import <nixpkgs> {};

in {
	# gotk4-nix = ../../gotk4-nix;
	gotk4-nix = systemPkgs.fetchFromGitHub {
		owner = "diamondburned";
		repo  = "gotk4-nix";
		rev   = "0a50408da4eb59ad4500db49785676714f4bd4e6";
		hash  = "sha256:1ryxxkxly298yr3m0868g69jgk8gagvg6zmpknv2l7jd8x3w4pz7";
	};
}
