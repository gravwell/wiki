# Abs

The abs module takes the absolute value of an enumerated value. With one argument, the enumerated value is overwritten with its absolute value:

```gravwell
tag=default json offset | abs offset
```

Add `as` and a new name to write the result to a different enumerated value and leave the original intact:

```gravwell
tag=default json offset | abs offset as offsetAbsolute | chart offsetAbsolute
```

Text values such as `-12` or `-3.5` are parsed as numbers. Entries where the enumerated value is missing or isn't a number are dropped.
