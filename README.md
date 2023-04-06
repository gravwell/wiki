# Gravwell Wiki

Current version of Gravwell documentation can always be found at http://docs.gravwell.io
This repo is served up at that url and can be cloned for offline use.

There is a shell.nix provided in the repo. To install Nix, visit https://nixos.org/download.html
Once nix is installed, just run `nix-shell` in the wiki root directory.

Inside of the nix-shell:

- build the documentation by running `make html`
- view changes as you edit by running `make livehtml` and visit http://127.0.0.1:8000

## Adding Sphinx Extensions

Adding a Sphinx extension requires two steps:

- Add relevant Python package(s) to the nix shell
- Add the extension to `conf.py`

Once the pacakges are installed and listed in `conf.py`, the extension will ready for use.

### Example adding `sphinx-copybutton`

**Adding `sphinx-copybutton` to nix shell**

1. Let's check if the extension is already available in nixpkgs by visiting https://search.nixos.org/packages and searching for sphinx-copybutton
   - If it's available, note the attribute path: `nixpkgs.python310Packages.sphinx-copybutton`
   - If it's not available, consult the nix docs for packaging python libs. We package `not-found-extension` ourselves in `package.nix`. That's a good starting point.
2. Because `sphinx-copybutton` is in nixpkgs, we can just add it to `packages.nix` in the `pythonBundle`.
3. Exit (if necessary) then re-enter the nix shell with `nix-shell`

**Adding `sphinx-copybutton` to `conf.py`**

1. Consult the docs. Looks like copy button just requires adding `"sphinx_copybutton"` to the list of extensions in `conf.py`.
2. Run `black conf.py` to format the file
