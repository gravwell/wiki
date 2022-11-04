import os
import sys

sys.path.insert(0, os.path.abspath("."))

from sphinx.highlighting import lexers

from gravy_lexer import GravwellLexer

# Configuration file for the Sphinx documentation builder.
#
# For the full list of built-in configuration values, see the documentation:
# https://www.sphinx-doc.org/en/master/usage/configuration.html

# -- Project information -----------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#project-information

project = "Gravwell"
copyright = "2022, Gravwell, Inc."
author = "Gravwell, Inc."
release = "v5.1.2"

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = ["myst_parser", "sphinx_design", "sphinxcontrib.spelling"]
myst_enable_extensions = ["colon_fence"]


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
    "header_links_before_dropdown": 6,
}


lexers["gw"] = lexers["gravwell"] = GravwellLexer(startinline=True)
