with (import ./packages.nix);
pkgs.mkShell { buildInputs = chosenPackages ++ [ pkgs.coreutils ]; }
