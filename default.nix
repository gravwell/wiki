{ VERSION_LIST_URL ? null }:
let
  # use a specific (although arbitrarily chosen) version of the Nix package collection
  pkgs = import (fetchTarball {
    url =
      "https://github.com/NixOS/nixpkgs/archive/24fe8bb4f552ad3926274d29e083b79d84707da6.tar.gz";
    # the sha256 makes sure that the downloaded archive really is what it was when this
    # file was written
    sha256 = "1ica2sangr5daiv19bj743mysp9cs46zl1mfpy909qyamh85612p";
  }) { };

  custom-aspell =
    pkgs.aspellWithDicts (d: [ d.en d.en-computers d.en-science ]);

  sphinx-favicon = with pkgs.python310.pkgs;
    buildPythonPackage rec {
      pname = "sphinx_favicon";
      version = "1.0.1";
      format = "wheel";
      src = pkgs.fetchurl {
        url =
          "https://files.pythonhosted.org/packages/92/c2/152bd6c211b847e525d2c7004fd98e3ac5baeace192716da8cd9c9ec2427/sphinx_favicon-1.0.1-py3-none-any.whl";
        hash = "sha256-fJPWtjTLTJaHzqtnqFJvBdOwJnnflOJz5RpDKC5rA0w=";
      };
      propagatedBuildInputs = [ sphinx ];
      pythonImportsCheck = [ "sphinx_favicon" ];
    };

  custom-pydata-sphinx-theme = with pkgs.python310.pkgs;
    buildPythonPackage rec {
      pname = "pydata-sphinx-theme";
      version = "0.15.2";
      format = "wheel";

      src = pkgs.fetchurl {
        url =
          "https://files.pythonhosted.org/packages/bc/ef/1700096a5c1d17c2d99332b0759c7ca70346aac917ceafa4c380f085d359/pydata_sphinx_theme-0.15.2-py3-none-any.whl";
        hash = "sha256-DF+h+pipsm2uWQZm/1dvJ+Jse6cI/udU7Lngc1ntRYg=";
      };

      propagatedBuildInputs = [
        sphinx
        accessible-pygments
        beautifulsoup4
        docutils
        packaging
        typing-extensions
      ];

      pythonImportsCheck = [ "pydata_sphinx_theme" ];
    };

  pythonBundle = pkgs.python310.withPackages (ps:
    with ps; [
      sphinx
      sphinx-favicon
      sphinx-autobuild
      sphinx-copybutton
      myst-parser
      custom-pydata-sphinx-theme
      sphinx-design
      black
      sphinx-notfound-page
    ]);

in pkgs.stdenv.mkDerivation {
  name = "gravwell-wiki";
  src = ./.;

  VERSION_LIST_URL = VERSION_LIST_URL;

  buildInputs = [ pythonBundle pkgs.gnumake pkgs.git custom-aspell ];
  buildPhase = ''
    make clean html
  '';

  installPhase = ''
    mkdir -p $out
    cp -r _build/html $out
  '';

  LOCALE_ARCHIVE = "${pkgs.glibcLocales}/lib/locale/locale-archive";
}
