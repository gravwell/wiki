// There's a race in setting the scorer, depending on order in which files are loaded.
// This definition omits a "var", "let" or "const" because there's a chance Scorer is already defined.
// By omitting the keyword, we can define or set the value of the variable, which is handy.
Scorer = {
  score: (result) => {
    const [docname, title, anchor, descr, score, filename] = result;
    if (docname.startsWith("changelog/")) {
      // Push matches on changelogs to the bottom
      return score - 100;
    }
    return score;
  },

  // Default values. See https://github.com/sphinx-doc/sphinx/blob/5ff3740063c1ac57f17ecd697bcd06cc1de4e75c/sphinx/themes/basic/static/searchtools.js#L10
  objNameMatch: 11,
  objPartialMatch: 6,
  objPrio: {
    0: 15,
    1: 5,
    2: -5,
  },
  objPrioDefault: 0,
  title: 15,
  partialTitle: 7,
  term: 5,
  partialTerm: 2,
};
