with import ./packages.nix;
pkgs.dockerTools.buildLayeredImage {
  # docker image name
  name = "gw-sphinx-docs";

  # image tag
  tag = "latest";

  # packages/files in docker image
  contents = chosenPackages ++ [ pkgs.bash pkgs.coreutils ];
}
