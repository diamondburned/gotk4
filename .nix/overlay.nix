self: super: {
	go = super.go.overrideAttrs (old: {
		version = "1.17beta1";
		src = builtins.fetchurl {
			url    = "https://golang.org/dl/go1.17rc1.linux-arm64.tar.gz";
			sha256 = "sha256:0kps5kw9yymxawf57ps9xivqrkx2p60bpmkisahr8jl1rqkf963l";
		};
		doCheck = false;
	});
	gopls = self.buildGoModule rec {
		pname = "gopls";
		version = "0.7.0";

		src = super.fetchgit {
			rev = "gopls/v0.7.0";
			url = "https://go.googlesource.com/tools";
			sha256 = "0vylrsmpszij23yngk7mfysp8rjbf29nyskbrwwysf63r9xbrwbi";
		};

		modRoot = "gopls";
		vendorSha256 = "1mnc84nvl7zhl4pzf90cd0gvid9g1jph6hcxk6lrlnfk2j2m75mj";

		doCheck = false;
		subPackages = [ "." ];
	};
}
