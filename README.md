# Gravwell Wiki

Current version of Gravwell documentation can always be found at http://docs.gravwell.io
This repo is served up at that url and can be cloned for offline use.

There is a shell.nix provided in the repo. To install Nix, visit https://nixos.org/download.html
Once nix is installed, just run `nix-shell` in the wiki root directory.

Inside of the nix-shell:

- build the documentation by running `make html`
- view changes as you edit by running `make livehtml` and visit http://127.0.0.1:8000

## Deploying

### Paths

In order for the version switcher to work properly, the paths where we serve the docs have to have the following structure:

- Docs for latest release are at `/`
  - Ends up being the target for old links, so old links continue to work as they always have
- Docs for previous releases are at `/vN.N.N/`

For example, if we're publishing v5.4.5, v5.4.4, and v5.4.3...

- `https://docs.gravwell.io` is v5.4.5
- `https://docs.gravwell.io/v5.4.4/` is v5.4.4
- `https://docs.gravwell.io/v5.4.3/` is v5.4.3

### Directory structure

Assume we're using a web server to serve a directory as `docs.gravwell.io`.

If we wanted to serve v5.4.5, v5.4.4, and v5.4.3 as shown above, we could structure the directory like this:

```
docs/     <- the full build of v5.4.5
├─ 404.html
├─ index.html
├─ ...
├─ v5.4.4   <- the full build of v5.4.4
│  ├─ 404.html
│  ├─ index.html
│  ├─ ...
├─ v5.4.3   <- the full build of v5.4.3
│  ├─ 404.html
│  ├─ index.html
│  ├─ ...
```

Nesting previous versions in the latest version works the way we want it to.

To deploy a new version:

1. Copy all `docs/vN.N.N/` directories somewhere safe
2. Rename `docs/` to `vN.N.N+1` (or whatever version it is)
3. Add the docs for the new version in `docs/`
   - Don't forget to update `_static/versions.json`!
4. Move all `v.N.N.N` and `vN.N.N+1` back into `docs/`

Links would also work:

```
docs/
├─ latest/      <- a link to v5.4.5
│  ├─ v5.4.4/   <- a link to v5.4.4
│  ├─ v5.4.3/   <- a link to v5.4.3
│  ├─ ...
├─ v5.4.5   <- the full build of v5.4.5
│  ├─ 404.html
│  ├─ index.html
│  ├─ ...
├─ v5.4.4   <- the full build of v5.4.4
│  ├─ 404.html
│  ├─ index.html
│  ├─ ...
├─ v5.4.3   <- the full build of v5.4.3
│  ├─ 404.html
│  ├─ index.html
│  ├─ ...
```

To deploy a new version:

1. Remove all `vN.N.N` links from `latest/`
2. Remove `latest/`
3. Add new docs build to `docs/`
   - Don't forget to update `_static/versions.json`!
4. Re-create `latest/`, this time pointing to the new release
5. Re-create `vN.N.N` links as desired in `latest/`

### versions.json

`https://docs.gravwell.io/_static/verions.json` dictates how the version selector will behave in all docs versions, past and current.

As implied by the path, only the `versions.json` that's included in the latest docs version needs to be up-to-date. It doesn't matter if any `.../vN.N.N/_static/versions.json` is out of date. Only the file at `/_static/verions.json` is used.

The [PyData Sphinx Theme docs](https://pydata-sphinx-theme.readthedocs.io/en/stable/user_guide/version-dropdown.html#add-a-json-file-to-define-your-switcher-s-versions) explain the optins in more detail, but here's a quick summary / example:

```js
[
  {
    // An optional special name to show in the selector
    name: "v5.4.5 (latest)",

    // The version. Should match `release` in `conf.py`
    version: "v5.4.5",

    // The path where one will be able to find this release (root path for latest)
    url: "https://docs.gravwell.io/",

    // Set to true if this is the latest version.
    // If set to false or omited, an "Old Version" warning banner will show on that version.
    // Only one entry in the versions array may have preferred set to true.
    preferred: true,
  },

  {
    // Since "name" is missing, the selector will just show "v5.4.4"
    version: "v5.4.4",
    url: "https://docs.gravwell.io/v5.4.4/",
    //  Since preferred is omitted, the "Old Version" warning banner will show for v5.4.4
  },

  {
    // Since "name" is missing, the selector will just show "v5.4.3"
    version: "v5.4.3",
    url: "https://docs.gravwell.io/v5.4.3/",
    //  Since preferred is omitted, the "Old Version" warning banner will show for v5.4.3
  },
];
```

When updating version.json...

- push a new entry in the front of the array
  - Ensure `name` and `preferred` are set as necessary
- unset `preferred` and `name` for the entry that's now at index 1
- remove old releases from the end as necessary (see below)

### Removing old releases

As verions accumulate, we may decide to stop hosting old versions. This requires:

1. Updating `_static/versions.json` hosted at `https://docs.gravwell.io/_static/versions.json`, so that it contains only those versions we want to continue hosting
2. Removing old versions from the web server

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
