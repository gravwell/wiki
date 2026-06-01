# Hex

The hex renderer displays search results in hexadecimal format, converting every byte to its two-digit hex representation (uppercase). Unprintable characters are displayed as their hex values rather than being replaced or converted, making this renderer ideal for examining the exact byte content of entries.

This renderer is particularly useful for:

* Analyzing network packet captures and protocol payloads
* Examining binary data in its raw form
* Investigating encoded content (base64, custom binary protocols, etc.)
* Comparing exact byte sequences across multiple entries
* Documenting file-like payloads in network traffic

## Example 

```gravwell
tag=gravwell hex
```

![](text.png)
