let
  # use a specific (although arbitrarily chosen) version of the Nix package collection
  default_pkgs = fetchTarball {
    url =
      "https://github.com/NixOS/nixpkgs/archive/c3616ce8370f4a00e6b62af80fdaace308c13b68.tar.gz";
    # the sha256 makes sure that the downloaded archive really is what it was when this
    # file was written
    sha256 = "0cbf5y66zllj63ndlcng5jlc5fhpp7ph1ribgi989xmdplf0h1r1";
  };
  # function header: we take one argument "pkgs" with a default defined above
in { pkgs ? import default_pkgs { } }:
let
  pythonBundle = pkgs.python310.withPackages (ps:
    with ps; [
      sphinx
      sphinx-autobuild
      myst-parser
      sphinxcontrib-spelling
      pydata-sphinx-theme
      sphinx-design
      black
    ]);
  # this is what the function returns: the result of a mkShell call with a buildInputs
  # argument that specifies all software to be made available in the shell
in pkgs.mkShell { buildInputs = [ pythonBundle ]; }
