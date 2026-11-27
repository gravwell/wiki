# XML

The xml module extracts elements and attributes from XML data into enumerated values. XML is untyped, so all results are strings.

Consider the following XML snippet, consisting of three nested elements `A`, `B`, and `C`; the inner-most element `C` has an associated attribute named `MyAttr`.

```
<A><B><C MyAttr="foo">bar</C></B></A>
```

Paths start at the root element. To extract the **value** of `C`:

```
xml A.B.C
```

This places `bar` in an enumerated value named `C`.

To extract the value of the `MyAttr` attribute instead, put the attribute name in square brackets:

```
xml A.B.C[MyAttr]
```

This creates an enumerated value named `MyAttr` containing `foo`.

If the final element contains other elements, the module returns its inner XML as a string:

```
xml A.B
```

returns `<C MyAttr="foo">bar</C>`.

Use `as` to give the enumerated value a different name. This helps when an element name doesn't work well in other modules; for instance, the eval module reads `My-Element` as `My` minus `Element`:

```
xml A.B.My-Element as MyElement
```

If an element or attribute name contains `.`, `[`, or `]`, wrap it in double quotes:

```
xml A."My.Element" as MyElement
```

## Filtering

The xml module supports one filter operator, `==`, and it behaves differently from filters in other modules: a failed match means nothing is extracted, but the entry is **not** dropped unless `-s` is set.

Consider this search:

```
xml A.B=="foo"
```

This entry matches, and the module extracts an enumerated value named `B`:

```
<A><B>foo</B></A>
```

This entry doesn't match, so `B` is not extracted. The entry still continues down the pipeline; use `xml -s A.B=="foo"` to drop it:

```
<A><B>bar</B></A>
```

XML often has several sibling elements with the same name. For example:

```
<System>
	<Data Name="OSVersion">Windows 10</Data>
	<Data Name="Username">gravwell</Data>
</System>
```

`xml System.Data` only extracts the first one, `Windows 10`. To pick a specific element, compare against one of its **attributes**. When the attribute matches, the module extracts the **element's** contents, not the attribute value:

```
xml System.Data[Name]=="Username" as Username
```

This creates an enumerated value named `Username` containing `gravwell`. Without `as`, the enumerated value is named after the attribute (`Name`), even though it holds the element's contents.


## Supported Options

* `-e <arg>`: Operate on an enumerated value instead of the entire entry. Entries without that enumerated value are dropped.
* `-s`: Strict mode. Drop the entry unless *all* extractions succeed, including any `==` filters.
