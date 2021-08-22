FROM nixpkgs/nix

WORKDIR /gotk4

# Prepare the Nix files.
COPY .nix .nix
COPY shell.nix shell.nix

# Prepare docker-env.
RUN nix-env -i -f .nix/docker-env.nix

# Initialize shell environment variables.
RUN /root/.nix-profile/bin/docker-env init

# Execute everything with the shell environment variables.
SHELL ["/root/.nix-profile/bin/docker-env", "exec"]

# Set the proper environment variables so -u works.
ENV HOME="/user/"
RUN mkdir -p /user && chmod -R 777 /user

# Populate Go things as a layer.
COPY go.mod go.sum ./
RUN  go mod download

# Copy the rest of the source code.
COPY . .

# Drop into a shell by default.
CMD bash
