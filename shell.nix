let
  packages = (import ./packages.nix);
  pkgs = packages.nixpkgs;
  chosenPackages = packages.chosenPackages;
in pkgs.mkShell { buildInputs = chosenPackages ++ [ pkgs.coreutils ]; }
