FROM nixpkgs/nix

WORKDIR /gotk4

# Prepare the Nix files.
COPY .nix .nix
COPY shell.nix shell.nix

ENV NIX_PATH="nixpkgs=channel:nixos-unstable"

# Prepare docker-env.
RUN nix-env --install --file .nix --argstr action ../../../gotk4/.nix/docker-env

# Initialize shell environment variables.
RUN /root/.nix-profile/bin/docker-env init

# Execute everything with the shell environment variables.
SHELL ["/root/.nix-profile/bin/docker-env", "exec"]
ENTRYPOINT ["/root/.nix-profile/bin/docker-env", "exec"]

# Set the proper environment variables so -u works.
ENV HOME="/user/"
RUN mkdir -p /user && chmod -R 777 /user

# Populate Go things as a layer.
COPY go.mod go.sum ./
RUN  go mod download
COPY pkg/go.mod pkg/go.sum ./pkg/
RUN  cd pkg && go mod download

# Copy the rest of the source code.
COPY . .

# Drop into a shell by default.
CMD bash
