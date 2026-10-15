import html
import os
import pathlib
import subprocess
import sys
import json
from datetime import date

from docutils import nodes
from sphinx.highlighting import lexers

sys.path.insert(0, os.path.abspath("."))

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
release = "v5.10.2"

# Default to localhost:8000, so the version switcher looks OK on livehtml
version_list_url = os.environ.get(
    "VERSION_LIST_URL", "http://localhost:8000/_static/versions.json"
)
print("Using version_list_url:", version_list_url)

# -- General configuration ---------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#general-configuration

extensions = [
    "myst_parser",
    "sphinx_design",
    "notfound.extension",
    "sphinx_copybutton",
    "sphinx_favicon",
    "sphinxcontrib.datatemplates",
]

myst_enable_extensions = [
    "colon_fence",
    "fieldlist",
    "substitution",
    "html_image",
]


templates_path = ["_templates"]
exclude_patterns = [
    ".github",
    "README.md",
    "CONTRIBUTING.md",
    "_build",
    "Thumbs.db",
    ".DS_Store",
    "env",
    "_vendor",
    "_tools",
]

language = "en"

# -- Options for HTML output -------------------------------------------------
# https://www.sphinx-doc.org/en/master/usage/configuration.html#options-for-html-output

# Title suffix for every page's <title> tag. Overrides the default
# "Gravwell vX.Y.Z documentation" so titles don't change with each release.
html_title = "Gravwell Documentation"

# Production URL of the latest docs. Sphinx uses this to emit a
# <link rel="canonical"> on every page; it's also used for the sitemap and
# Open Graph URLs below.
html_baseurl = "https://docs.gravwell.io/"

html_show_sourcelink = False
html_copy_source = False

html_theme = "pydata_sphinx_theme"
html_static_path = ["_static"]
html_extra_path = ["open_source"]
html_css_files = ["css/custom.css"]
html_search_scorer = "search_scorer.js"
html_theme_options = {
    "logo": {
        "image_light": "_static/images/Gravwell-Color.svg",
        "image_dark": "_static/images/Gravwell-Color-Reverse.svg",
    },
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
    "footer_start": [
        "sphinx-version",
        "theme-version",
    ],
    "footer_end": [
        "copyright",
        "git-commit-footer",
    ],
    "navigation_with_keys": False,
    #
    # Version switcher
    #
    "switcher": {
        "json_url": version_list_url,
        # The `version` field of each entry in verions.json must match a vN.N.N release name
        "version_match": release,
    },
    # Show a warning banner if the user's looking at any page other than latest
    "show_version_warning_banner": True,
    # Don't fail to compile just because the version switcher file (json_url) isn't reachable
    "check_switcher": False,
    # include the version switcher next to the logo
    "navbar_start": ["navbar-logo", "version-switcher"],
    # use custom center for navbar, so that we can better manage responsiveness
    "navbar_center": ["dynamic-navbar"],
}

# sphinx-favicon
favicons = [
    {
        "rel": "icon",
        "sizes": "48x48",
        "href": "favicon.ico",
    },
]

# -- Gravwell Query Language Config ----------------------------

lexers["gw"] = lexers["gravwell"] = GravwellLexer(startinline=True)


# -- "Not Found" Extension Config ----------------------------

notfound_urls_prefix = ""


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
    "git_commit": commit_id,
    "git_dirty_tree": is_dirty_tree,
}

# Variables to substitute in Markdown files
myst_substitutions = {}

# Copy button
copybutton_selector = "div.highlight pre,div.docutils pre.literal-block"


# Re-generate open_source.json at the beginning of each build.
# This keeps open_source.json up-to-date as files change.
def build_inited_handler(_app):
    open_source_dir = "open_source/"
    open_source_files = sorted(
        [
            {
                "name": os.path.join(root[len(open_source_dir) :], f),
                "path": os.path.join(root, f),
            }
            for root, _dirs, files in os.walk(open_source_dir)
            for f in files
        ],
        key=lambda x: x["path"],
    )

    # If the content of open_source.json hasn't changed, return early.
    #
    # Writing open_source.json will re-trigger a build on "make livehtml".
    # Without the early return, "make livehtml" will infinitely loop.
    try:
        with open("open_source.json", "r") as jsonFile:
            current_content = json.load(jsonFile)
            if current_content == open_source_files:
                return
    except OSError:
        # If the file doesn't exist, we'll just create it
        pass

    with open("open_source.json", "w") as jsonFile:
        json.dump(
            open_source_files,
            jsonFile,
            indent=2,
        )


# -- SEO: meta description, Open Graph tags, sitemap, robots.txt ---------------

# Pages that shouldn't appear in the sitemap
sitemap_exclude = {"404", "search", "genindex"}

# Paragraphs inside these nodes are poor page summaries, so skip them when
# looking for a description
description_skip_ancestors = (
    nodes.Admonition,
    nodes.table,
    nodes.list_item,
    nodes.field_list,
    nodes.sidebar,
    nodes.topic,
)


def page_description(doctree, max_len=160):
    """Return the text of the page's first body paragraph, truncated for use in a meta description."""
    for paragraph in doctree.findall(nodes.paragraph):
        parent = paragraph.parent
        skip = False
        while parent is not None:
            if isinstance(parent, description_skip_ancestors):
                skip = True
                break
            parent = parent.parent
        if skip:
            continue

        text = " ".join(paragraph.astext().split())
        if len(text) < 40:
            continue
        if len(text) > max_len:
            text = text[: max_len - 1].rsplit(" ", 1)[0].rstrip(",.;:") + "…"
        return text
    return None


def html_page_context_handler(app, pagename, _templatename, context, doctree):
    if doctree is None:
        return

    metatags = context.get("metatags", "")
    tags = []

    title = app.env.titles[pagename].astext() if pagename in app.env.titles else ""
    url = app.config.html_baseurl + app.builder.get_target_uri(pagename)

    # Pages can set their own description with myst front matter (myst.html_meta.description).
    # Only generate one when the page doesn't.
    description = None
    if 'name="description"' not in metatags:
        description = page_description(doctree)
        if description:
            tags.append(f'<meta name="description" content="{html.escape(description)}" />')

    tags.append('<meta property="og:type" content="website" />')
    tags.append('<meta property="og:site_name" content="Gravwell Documentation" />')
    tags.append(f'<meta property="og:url" content="{html.escape(url)}" />')
    if title:
        tags.append(f'<meta property="og:title" content="{html.escape(title)}" />')
    if description:
        tags.append(f'<meta property="og:description" content="{html.escape(description)}" />')

    context["metatags"] = metatags + "\n".join(tags) + "\n"


def build_finished_handler(app, exception):
    if exception is not None or app.builder.format != "html":
        return

    outdir = pathlib.Path(app.outdir)

    urls = [
        app.config.html_baseurl + app.builder.get_target_uri(docname)
        for docname in sorted(app.env.found_docs)
        if docname not in sitemap_exclude
    ]
    sitemap = ['<?xml version="1.0" encoding="UTF-8"?>']
    sitemap.append('<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">')
    sitemap.extend(f"  <url><loc>{html.escape(u)}</loc></url>" for u in urls)
    sitemap.append("</urlset>")
    (outdir / "sitemap.xml").write_text("\n".join(sitemap) + "\n")

    # Old versions are excluded from search results with a noindex tag added at deploy
    # time (see .github/workflows/build-html.yml), so crawlers must be allowed to fetch them.
    (outdir / "robots.txt").write_text(
        f"User-agent: *\nAllow: /\n\nSitemap: {app.config.html_baseurl}sitemap.xml\n"
    )


def setup(app):
    app.connect("builder-inited", build_inited_handler)
    app.connect("html-page-context", html_page_context_handler)
    app.connect("build-finished", build_finished_handler)
