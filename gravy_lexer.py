from pygments.lexer import RegexLexer, words
from pygments.token import *

macroToken = Name.Variable
compoundQueryToken = Name.Variable
identifierToken = Name


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
            (r"tag\s*=", Keyword, ("pipeModuleInvocation", "tagModuleInvocation")),
            (r"\|", Punctuation, "pipeModuleInvocation"),
            #
            # Keywords (suffix \b says "must be standalone word")
            #
            (words(("as", "by", "over"), suffix=r"\b"), Keyword),
            #
            # Macro
            #
            (r"\$[\w]+", macroToken),
            #
            # Compund Queries
            #
            (r"@[\w]+", compoundQueryToken),
            #
            # Flag
            #
            (r"-[\w]+", identifierToken),
            #
            # Identifier
            #
            (r"[\w]+", identifierToken),
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
            (r"\$[\w]+", macroToken, "#pop"),
            #
            # Compund Query
            #
            (r"@[\w]+", compoundQueryToken, "#pop"),
            #
            # Identifier
            #
            (r"[\w,]+", identifierToken, "#pop"),
        ],
        "pipeModuleInvocation": [
            #
            # Whitespace
            #
            (r"[ \t\r\n]+", Whitespace),
            #
            # Macro
            #
            (r"\$[\w]+", macroToken, "#pop"),
            #
            # SearchModule
            #
            (r"\w+", Name.Function, "#pop"),
        ],
        "comment": [
            (r"[^/*]+", Comment.Multiline),
            (r"/\*", Comment.Multiline, "#push"),
            (r"\*\/", Comment.Multiline, "#pop"),
            (r"[/*]", Comment.Multiline),
        ],
        "string": [
            (r"\$[\w]+", macroToken),
            (r'[^"]', String),
            (r'"', String.Delimiter, "#pop"),
        ],
    }
