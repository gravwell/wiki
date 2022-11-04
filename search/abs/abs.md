# Abs

The abs module takes the absolute value of the specified enumerated value. If only one argument is given, the enumerated value is overwritten with its absolute value:

```gravwell
tag=default json offset | abs offset
```

It is also possible to specify a 'destination' enumerated value name, which will leave the original intact:

```gravwell
tag=default json offset | abs offset as offsetAbsolute | chart offsetAbsolute
```
