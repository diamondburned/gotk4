self: super: {
	go = super.go.overrideAttrs (old: {
		version = "1.18beta1";
		src = builtins.fetchurl {
			url    = "https://go.dev/dl/go1.18beta1.src.tar.gz";
			sha256 = "sha256:18akwrw4lzl6cj3yy5hrnf4zfycx84xas1s95mdwp6a6n66h5321";
		};
		doCheck = false;
		patches = [
			# cmd/cgo: concurrent file generation
			(builtins.fetchurl "https://github.com/diamondburned/go/commit/2f584e7d3759ba19c13e694198087545f80b7ad3.patch")
		];
	});
	buildGoModule = super.buildGoModule.override {
		inherit (self) go;
	};
}
