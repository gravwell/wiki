let
  packages = (import ./packages.nix);
  pkgs = packages.nixpkgs;
  chosenPackages = packages.chosenPackages;
in pkgs.dockerTools.buildLayeredImage {
  # docker image name
  name = "gw-sphinx-docs";

  # image tag
  tag = "latest";

  # packages/files in docker image
  contents = chosenPackages ++ [ pkgs.bash pkgs.coreutils ];
}
