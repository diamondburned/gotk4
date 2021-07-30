{
	overlays ? [],
	# The declarations, where a pinned Nixpkgs Unstable is fetched. When
	# updating, only rev and sha256 should be changed.
	src ? ((import <nixpkgs> {}).fetchFromGitHub {
		owner  = "NixOS";
		repo   = "nixpkgs";
		rev    = "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94";
		sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
	}),
}:

import src {
	overlays = [ (import ./overlay.nix) ] ++ overlays;
}
