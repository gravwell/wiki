import os
import subprocess
import sys
import pathlib

sys.path.insert(0, os.path.abspath("."))

from datetime import date
from sphinx.highlighting import lexers
from gravy_lexer import GravwellLexer

# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

project = "Gravwell"
copyright = f"Gravwell, Inc. {date.today().year}"
author = "Gravwell, Inc."
release = "v5.1.2"

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = ["myst_parser", "sphinx_design", "sphinxcontrib.spelling"]
myst_enable_extensions = [
    "colon_fence",
    "fieldlist",
    "substitution",
]


templates_path = ["_templates"]
exclude_patterns = ["README.md", "_build", "Thumbs.db", ".DS_Store", "env"]

language = "en"

# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

html_show_sourcelink = False
html_copy_source = False

html_theme = "pydata_sphinx_theme"
html_static_path = ["_static"]
html_css_files = ["css/custom.css"]
html_theme_options = {
    "logo": {
        "image_light": "images/Gravwell-Color.svg",
        "image_dark": "images/Gravwell-Color-Reverse.svg",
    },
    "favicons": [
        {
            "rel": "icon",
            "sizes": "48x48",
            "href": "favicon.ico",
        },
    ],
    "icon_links": [
        {
            # Label for this link
            "name": "GitHub",
            # URL where the link will redirect
            "url": "https://github.com/gravwell",  # required
            # Icon class (if "type": "fontawesome"), or path to local image (if "type": "local")
            "icon": "fa-brands fa-github",
            # The type of image to be used (see below for details)
            "type": "fontawesome",
        },
        {
            # Label for this link
            "name": "Discord",
            # URL where the link will redirect
            "url": "https://discord.com/invite/gravwell",  # required
            # Icon class (if "type": "fontawesome"), or path to local image (if "type": "local")
            "icon": "fa-brands fa-discord",
            # The type of image to be used (see below for details)
            "type": "fontawesome",
        },
    ],
    "header_links_before_dropdown": 6,
    "footer_items": [
        "git-commit-footer",
        "copyright",
        "sphinx-version",
    ],
}


lexers["gw"] = lexers["gravwell"] = GravwellLexer(startinline=True)


# -- Substitutions -------------------------------------------------

# Determine git commit ID
#
# Make sure we're checking the commit id for the wiki project, by passing the directory
# to the git subprocess call
git_dir = pathlib.Path(__file__).parent.resolve()
commit_id = (
    subprocess.check_output(["git", "rev-parse", "--short", "HEAD"], cwd=git_dir)
    .strip()
    .decode("ascii")
)
try:
    subprocess.check_call(["git", "diff", "--quiet"], cwd=git_dir)
    is_dirty_tree = False
except subprocess.CalledProcessError:
    is_dirty_tree = True

# Variables to substitute in HTML template files
html_context = {
    "git_commit": f"{commit_id}{'*' if is_dirty_tree else ''}",
}

# Variables to substitute in Markdown files
myst_substitutions = {}
