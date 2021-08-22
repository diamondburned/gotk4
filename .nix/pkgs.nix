let githubTarURL = owner: repo: hash:
	"https://github.com/${owner}/${repo}/archive/${hash}.tar.gz";

in {
	overlays ? [],
	# The declarations, where a pinned Nixpkgs Unstable is fetched. When
	# updating, only rev and sha256 should be changed.
	src ? (builtins.fetchTarball {
		name   = "gotk4-nixpkgs";
		url    = "${githubTarURL "nixos" "nixpkgs" "fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94"}";
		sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
	}),
}:

import src {
	overlays = [ (import ./overlay.nix) ] ++ overlays;
}
