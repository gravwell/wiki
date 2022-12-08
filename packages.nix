let
  # use a specific (although arbitrarily chosen) version of the Nix package collection
  pkgs = import (fetchTarball {
    url =
      "https://github.com/NixOS/nixpkgs/archive/c3616ce8370f4a00e6b62af80fdaace308c13b68.tar.gz";
    # the sha256 makes sure that the downloaded archive really is what it was when this
    # file was written
    sha256 = "0cbf5y66zllj63ndlcng5jlc5fhpp7ph1ribgi989xmdplf0h1r1";
  }) { };

  custom-aspell = pkgs.aspellWithDicts (d: [ d.en d.en-computers d.en-science ]);

  pythonBundle = pkgs.python310.withPackages (ps:
    with ps; [
      sphinx
      sphinx-autobuild
      myst-parser
      pydata-sphinx-theme
      sphinx-design
      black
    ]);

in {
  nixpkgs = pkgs;
  chosenPackages = [ pythonBundle pkgs.gnumake pkgs.git custom-aspell ];
}
