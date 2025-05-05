# Contributing to the Gravwell Wiki

We welcome issues and pull requests from external contributors. This file documents a few guidelines to keep in mind.

* You should only ever need to modify `.md` files, which are compiled to HTML automatically in deployment. To deploy a local version of the docs server to see your changes as you work, run `nix-shell` in the repo's top level directory, then run `make livehtml` and go to [http://localhost:8000/](http://localhost:8000)

* Please make issues as descriptive as possible. If you find a particular portion of the documentation confusing, for instance, try to explain what makes it hard to understand and suggest what might make it clearer.

* Pull requests should include a description of what you've changed and why. Spelling and grammar corrections are welcome, but if you're going to fix one spelling error, please take the time to check the rest of the file while you're at it.

* Please don't modify any `toctree` sections in the Markdown files; these help define navigation in the Gravwell docs and it's too easy to break them.

Here are some basic style rules:

* American English spellings (color rather than colour, etc.).
* Markdown `single backtick` code for filenames and variables.
* Gravwell queries should be in triple-backtick code blocks.
* Please use ASCII quote marks, not Unicode quotes or the grave character (see https://www.cl.cam.ac.uk/~mgk25/ucs/quotes.html for more information about quotation marks).
