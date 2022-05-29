{ action }:

let src = import ./src.nix;

in import "${src.gotk4-nix}/${action}.nix" {
	base = {
		pname   = "gotk4";
		version = "dev";
	};
	pkgs = import "${src.gotk4-nix}/pkgs.nix" {
		sourceNixpkgs = builtins.fetchTarball {
			name   = "gotk4-nixpkgs";
			url    = "https://github.com/NixOS/nixpkgs/archive/fbfb79400a08bf754e32b4d4fc3f7d8f8055cf94.tar.gz";
			sha256 = "0pgyx1l1gj33g5i9kwjar7dc3sal2g14mhfljcajj8bqzzrbc3za";
		};
		useFetched = true;
	};
}
