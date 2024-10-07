# dedup-links

`dedup-links` is a command line utility written in Go that walks a given static website's source to...

- locate image files that are duplicated within the source
- identify an arbitrary "original" copy of a given image
- rewrite `<img>` tag `src` attributes to use "original" images
- delete duplicate, "non-original" image files

... in order to reduce the size of the website source.

## Usage

```
$ dedup-links --help
Usage: dedup-links [--baseurl BASEURL] [--workers WORKERS] ROOT

Positional arguments:
  ROOT                   root directory to walk when replacing links

Options:
  --baseurl BASEURL      base URL to use when constructing absolute paths [default: /]
  --workers WORKERS [default: 4]
  --help, -h             display this help and exit
```

**ROOT**: The root path is a required positional argument. The tool will descend into that directory to identify "original" and "non-original" image files. It will also search the root path for HTML files that refer to those images.

**--baseurl**: The baseurl is an optional flag that tells the tool how to construct absolute url paths to images. Because the tool replaces relative url paths with absolute url paths, it's important to know if the site will be served at some subpath or not. If baseurl is not provided, the tool assumes there is no subpath, and absolute url paths will simply start with `/`

**--workers**: The number of workers to use when calculating file hashes and rewriting `img[src]` URL paths

**--help**: Show the command usage

## Limitations

- The tool only inspects `.html` files. If a CSS (or JS or WASM or whatever) source file references a "non-original" image, that source file will not be updated and the image file will be deleted anyway.

## Development

This is a pretty typical Go program.

```command
# enter the nix shell, so that you've got access to all the right versions of tools
nix-shell

# use "go run" to test changes. We're not using a vendor/ directory, so -mod=mod might be necessary.
go run -mod=mod ./main.go

# build the program with nix (this works inside the shell and outside the shell)
nix-build

# run the executable created by nix-build
./result/bin/dedup-links
```

## Relationship the wiki repo

The derivation present in this directory's `default.nix` is used by the top-level `default.nix`. That is, when you enter `nix-shell` at the top of `gravwell/wiki`, nix will automatically build `dedup-links` and put it on your `PATH`.

This makes it easy to build and invoke `dedup-links` without knowing how to build Go programs. Plus, it makes the tool easy to use within GitHub Actions.
