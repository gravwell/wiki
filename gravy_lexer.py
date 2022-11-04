from pygments.lexer import RegexLexer, words
from pygments.token import *


class GravwellLexer(RegexLexer):
    name = "Example Lexer with states"

    tokens = {
        "root": [
            #
            # Whitespace
            #
            (r"[ \t\r\n]+", Whitespace),
            #
            # Module invocations
            #
            (r"tag\s*=", Operator, ("pipeModuleInvocation", "tagModuleInvocation")),
            (r"\|", Operator, "pipeModuleInvocation"),
            #
            # Keywords (suffix \b says "must be standalone word")
            #
            (words(("as", "by", "over"), suffix=r"\b"), Keyword),
            #
            # Macro
            #
            (r"\$[\w]+", Name.Tag),
            #
            # Compund Queries
            #
            (r"@[\w]+", Name.Function),
            #
            # Flag
            #
            (r"-[\w]+", Name.Variable),
            #
            # Identifier
            #
            (r"[\w]+", Name),
            #
            # Braces
            #
            (r"[{}()[\]]", Punctuation),
            #
            # strings
            #
            (r'"([^"])*$', Error),  # non-teminated string
            (r'"', String.Delimiter, "string"),
            #
            # open comment
            #
            (r"\/\*", Comment.Multiline, "comment"),
            #
            # Operators
            #
            (
                words(
                    (
                        "+",
                        "-",
                        "*",
                        "/",
                        "%",
                        "&",
                        "^",
                        "<",
                        ">",
                        "=",
                        "!",
                        "[",
                        "]",
                        "{",
                        "}",
                        ",",
                        ";",
                        ".",
                        ":",
                        "<<",
                        ">>",
                        "+=",
                        "-=",
                        "&&",
                        "||",
                        "++",
                        "--",
                        "==",
                        "!=",
                        "~",
                    ),
                ),
                Operator,
            ),
        ],
        "tagModuleInvocation": [
            #
            # Whitespace
            #
            (r"[ \t\r\n]+", Whitespace),
            #
            # Macro
            #
            (r"\$[\w]+", Name.Tag, "#pop"),
            #
            # Compund Query
            #
            (r"@[\w]+", Name.Function, "#pop"),
            #
            # Identifier
            #
            (r"[\w]+", Name, "#pop"),
        ],
        "pipeModuleInvocation": [
            #
            # Whitespace
            #
            (r"[ \t\r\n]+", Whitespace),
            #
            # Macro
            #
            (r"\$[\w]+", Name.Tag, "#pop"),
            #
            # SearchModule
            #
            (r"\w+", Name.Class, "#pop"),
        ],
        "comment": [
            (r"[^/*]+", Comment.Multiline),
            (r"/\*", Comment.Multiline, "#push"),
            (r"\*\/", Comment.Multiline, "#pop"),
            (r"[/*]", Comment.Multiline),
        ],
        "string": [
            (r"\$[\w]+", Name.Tag),
            (r'[^"]', String),
            (r'"', String.Delimiter, "#pop"),
        ],
    }
