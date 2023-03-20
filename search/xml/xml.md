## XML

The xml module can extract components from XML data into enumerated values for later use. Since XML is untyped, all results are returned as strings.

Consider the following XML snippet, consisting of three nested elements `A`, `B`, and `C`; the inner-most element `C` has an associated attribute named `MyAttr`.

```
<A><B><C MyAttr="foo">bar</C></B></A>
```

To extract the **value** of `C`, we would use the following query:

```
xml A.B.C
```

The query places the result, "bar", in an enumerated value named "C".

If, on the other hand, we wanted to extract the value of the MyAttr attribute, we would use the following syntax:

```
xml A.B.C[MyAttr]
```

This creates an enumerated value named "MyAttr" containing the value "foo".

If the final element specified contains more XML elements, the module will return the inner XML as a string:

```
xml A.B
```

returns ```<C MyAttr="foo">bar</C>```.

XML elements are not always named in a way suitable for use with other modules; for instance, an element named "My-Element" will confuse the eval module, which will think you wish to subtract a variable `Element` from another variable `My`. The xml module allows you to select a different name for the resulting enumerated value:

```
xml A.B.My-Element as MyElement
```

Similarly, if an element or attribute contains a ".", "[", or "]" character, you can simply enclose the element in quotes to let the module know that the special character is part of the element name:

```
xml A."My.Element" as MyElement
```

### Filtering

The xml module allows for simple pre-filtering of data, to avoid invoking addition modules (like eval) when possible. Because of the way XML works, there are a few peculiarities in how the xml module handles filtering, so it is important to read this section carefully.

The module can test if an element value is equal to a literal value. If so, the requested element is extracted and the search entry is passed along the pipeline. Consider this search:

```
xml A.B=="foo"
```

The following XML entry will pass the test and continue down the pipeline with a new enumerated element named "B":

```
<A><B>foo</B></A>
```

This entry will not:

```
<A><B>bar</B></A>
```

This is all more or less intuitive. However, XML also allows multiple child elements at the same level, making it very difficult to extract the desired data. For example, one might see something like this:

```
<System>
	<Data Name="OSVersion">Windows 10</Data>
	<Data Name="Username">gravwell</Data>
</System>
```

Simply saying `xml System.Data` will only extract the first result, "Windows 10". In order to extract the username, you can compare against an **attribute**; when the xml module finds a match, however, it extract the current **element** data:

```
xml System.Data[Name]=="Username"
```

This query results in a new enumerated value named `Data` containing the string "gravwell".


### Supported Options

* `-e <arg>`: The “-e” option operates on an enumerated value instead of on the entire record.