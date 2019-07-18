# Gauge

The gauge renderer is a condensing renderer used to turn entries into one or more final values suitable for display as "gauges". For instance, you might wish to find the total number of brute-force attempts over the last hour and display it in a dashboard. While this could be accomplished with the table renderer, the gauge renderer makes for a more attractive result with at-a-glance readability.

## Basic Usage

The simplest way to use the gauge renderer is by passing it a single enumerated value argument:

```
tag=json json class | stats mean(class) | gauge mean
```

![](gauge1.png)

Selecting the gear icon allows you to change some options on the gauge. Clicking the 'Half' will change the style of the gauge display:

![](gauge2.png)

Selecting 'Number card' in the chart type dropdown will change the display to the other kind of gauge:

![](gauge3.png)

## Specifying Max and Min Limits

You can specify minimum and maximum values for the gauge by wrapping the magnitude enumerated value and the desired min/max values in parentheses:

```
tag=json json class | stats mean(class) | gauge (mean 1 100000)
```

![](gauge-minmax1.png)

You can also specify the minimum and maximum by enumerated values:

```
tag=json json class | stats mean(class) min(class) max(class) | gauge (mean min max)
```

![](gauge-minmax2.png)

Or use a mix of constants and enumerated values:

```
tag=json json class | stats mean(class) max(class) | gauge (mean 1 max)
```

## Multiple Gauges

You can list multiple enumerated values to place multiple needles on the gauge:

```
tag=json json class | stats mean(class) stddev(class) | gauge mean stddev
```

![](gauge-multi1.png)

You can specify min/max values for each needle separately if desired, but note that default single-gauge renderer will select the lowest min and highest max for display, ignoring the others. For that reason, you may wish to select the "multiple gauges" option in the configuration menu:

```
tag=json json class | stats mean(class) stddev(class) min(class) max(class) | gauge (mean min max) (stddev 1 35000)
```

![](gauge-multi2.png)

The renderer also behaves appropriately in "number card" mode with multiple items:

![](gauge-multi3.png)
