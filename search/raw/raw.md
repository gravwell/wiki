# Raw

The raw renderer is functionally similar to the text renderer, but does not attempt to modify or change any non-printable characters. This renderer hands back the raw record, for better or worse. This renderer can be useful when passing data back to other tools which need the raw values, or when you just want to see if your browser can take a stab at turning packets into emojis.

Example: `raw limit 2048`