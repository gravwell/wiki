let
  # use a specific (although arbitrarily chosen) version of the Nix package collection
  pkgs = import (fetchTarball {
    url =
      "https://github.com/NixOS/nixpkgs/archive/e0464e47880a69896f0fb1810f00e0de469f770a.tar.gz";
    sha256 = "sha256:1maakx00q48r6q6njxrajyhrq27xsnnayarc8j33p7x9f6pxlbyg";
  }) { };

  buildGoModule = pkgs.buildGoModule.override { go = pkgs.go_1_23; };

in buildGoModule {
  name = "dedup-links";
  src = ./.;

  vendorHash = "sha256-zSh3L3N7ak80nrx9c2lA1bBVcqRbVP1HYX4X3lWQmCc=";
}
