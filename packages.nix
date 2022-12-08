let
  # use a specific (although arbitrarily chosen) version of the Nix package collection
  pkgs = import (fetchTarball {
    url =
      "https://github.com/NixOS/nixpkgs/archive/c3616ce8370f4a00e6b62af80fdaace308c13b68.tar.gz";
    # the sha256 makes sure that the downloaded archive really is what it was when this
    # file was written
    sha256 = "0cbf5y66zllj63ndlcng5jlc5fhpp7ph1ribgi989xmdplf0h1r1";
  }) { };

  custom-aspell =
    pkgs.aspellWithDicts (d: [ d.en d.en-computers d.en-science ]);

  not-found-extension = with pkgs.python310.pkgs;
    buildPythonPackage rec {
      pname = "sphinx_notfound_page";
      version = "0.8.3";
      format = "wheel";
      src = fetchPypi {
        inherit pname version format;
        sha256 =
          "c4867b345afccef72de71fb410c412540dfbb5c2de0dc06bde70b331b8f30469";
      };
      buildInputs = [ pkgs.python310Packages.sphinx ];
    };

  pythonBundle = pkgs.python310.withPackages (ps:
    with ps; [
      sphinx
      sphinx-autobuild
      myst-parser
      pydata-sphinx-theme
      sphinx-design
      black
      not-found-extension
    ]);

in {
  nixpkgs = pkgs;
  chosenPackages = [ pythonBundle pkgs.gnumake pkgs.git custom-aspell ];
}
