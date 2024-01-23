with (import ./packages.nix);
pkgs.mkShell {
  LOCALE_ARCHIVE = "${pkgs.glibcLocales}/lib/locale/locale-archive";
  buildInputs = chosenPackages ++ [ pkgs.coreutils ];
}
