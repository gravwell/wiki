# Text

The text renderer is designed to show human readable entries in a text format. Any non-printable characters will be converted to the ‘.’ character. Text also fully supports Unicode and can render non-ASCII characters. Text is the default renderer and is applied if no renderer is specified.
Text also has a default limit of approximately 1000 characters per entry, to prevent accidentally displaying multiple megabytes of raw data. To increase the maximum length of output add the `limit <n>` argument, where `n` is the number of characters to display.

Example: `text limit 4096`