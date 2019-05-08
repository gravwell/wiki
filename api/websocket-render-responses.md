# Response formats

Although all modules respond to the same commands, the format in which they return entries differs due to the differing nature of the data types involved. The responses described here are common to RESP_GET_ENTRIES, RESP_STREAMING, and RESP_TS_RANGE; we use these request/response IDs indiscriminately through the examples in this section.

## Text & raw module responses

The 'text' and 'raw' render modules return their entries as an array in a field labeled "Entries":

```
{
	"ID": 16,
	"EntryCount": 1575,
	"AdditionalEntries": false,
	"Finished": true,
	"Entries": [
		{
			"TS":"2018-04-02T16:16:39-06:00",
			"SRC":"",
			"Tag":0,"Data"<elided>",
			"Enumerated": [
				{"Name":"count","Value":{"Type":9,"Data":1}},
				{"Name":"SrcIP","Value":{"Type":13,"Data":"10.177.98.189"}}
			]
		},
<entries elided>
		{
			"TS":"2018-04-02T16:16:39-06:00",
			"SRC":"",
			"Tag":0,
			"Data":"<elided>",
			"Enumerated": [
				{"Name":"count","Value":{"Type":9,"Data":1}},
				{"Name":"SrcIP","Value":{"Type":13,"Data":"35.174.22.108"}}
			]
		}
	]
}
```

## Table module responses

The table module returns the entries in a field called "Entries", containing a structure defining Rows and Columns:

```
{
	"ID": 16,
	"EntryCount": 1575,
	"AdditionalEntries": false,
	"Finished": true,
	"Entries": {
		"Rows": [
			{
				"TS": "2018-04-02T10:30:29-06:00",
				"Row": [
					"10.144.162.236",
					"9410"
				]
			},
<67 similar entries elided>
			{
				"TS": "2018-04-02T10:28:51-06:00",
				"Row": [
					"192.168.1.1",
					"2"
				]
			}
		],
		"Columns": [
			"SrcIP",
			"count"
		]
	}
}
```

## Gauge module responses

The gauge module returns entries as an array of structures containing the gauge's name, the magnitude, and (optionally) the minumum and maximum values defined for this gauge:

```
{
	"ID": 16,
	"EntryCount": 1,
	"AdditionalEntries": true,
	"Finished": true,
	"Entries": [
		{
			"Name": "mean",
			"Magnitude": "31691.213",
			"Min": "0",
			"Max": "64000"
		}
	]
}
```

## Point-to-Point module responses

The point2point module returns an array of entries containing DstLocation, SrcLocation, and Magnitude fields. Optionally, the entries may also contain a 'Values' array, containing additional enumerated values specified as arguments to the renderer. The names for these enumerated values are given in the 'ValueNames' array.

This query:

```
tag=pcap packet tcp.Port ipv4.SrcIP ipv4.DstIP ipv4.Length | geoip SrcIP.Location as srcloc DstIP.Location as dstloc | sum Length by srcloc dstloc | point2point -srcloc srcloc -dstloc dstloc -mag sum SrcIP DstIP
```

should produce a result like this:

```
{
    "AdditionalEntries": true,
    "Entries": [
        {
            "DstLocation": "33.381516 -108.391164",
            "Magnitude": 420471,
            "SrcLocation": "34.054400 -118.244000",
            "Values": [
                "151.11.24.133",
                "192.168.2.60"
            ]
        },
        {
            "DstLocation": "33.381516 -108.391164",
            "Magnitude": 373204,
            "SrcLocation": "52.382400 5.899500",
            "Values": [
                "185.19.10.154",
                "192.168.2.60"
            ]
        },
        {
            "DstLocation": "33.381516 -108.391164",
            "Magnitude": 246593,
            "SrcLocation": "39.048100 -76.472800",
            "Values": [
                "53.1.11.28",
                "192.168.2.60"
            ]
        },
[...]
        {
            "DstLocation": "32.769700 -122.393300",
            "Magnitude": 8662,
            "SrcLocation": "33.381516 -108.391164",
            "Values": [
                "192.168.2.60",
                "192.33.23.124"
            ]
        }
    ],
    "EntryCount": 16,
    "Finished": true,
    "ID": 18,
    "ValueNames": [
        "SrcIP",
        "DstIP"
    ]
}
```

Note that the "Values" arrays on each entry correspond to the titles in the "ValueNames" array, e.g. the first entry has a "SrcIP" of 151.11.24.133.

## Chart module responses

The chart module returns entries in a field called "Entries", containing a structure which defines "Names" and "Values". The "Names" component is an array of names for the lines being plotted; in the case of this example, it contains IP addresses. The "Values" component contains a timestamp and a "Data" array; the elements in the Data array are the values corresponding to the names in the "Names" array at the given timestamp.

```
{
    "EntryCount": 5,
    "Finished": true,
    "ID": 18
    "AdditionalEntries": false,
    "Entries": {
        "Names": [
            "10.177.98.189",
            "192.168.1.101",
            "192.168.1.50",
            "10.174.22.108",
            "10.203.116.250",
            "10.21.206.134"
        ],
        "Values": [
            {
                "Data": [
                    0,
                    0,
                    0,
                    0,
                    0,
                    0
                ],
                "TS": "2018-04-02T15:22:13.815-06:00"
            },
<elided>
            {
                "Data": [
                    1,
                    7,
                    2,
                    1,
                    1,
                    1
                ],
                "TS": "2018-04-02T16:16:13.815-06:00"
            },
<elided>
            {
                "Data": [
                    0,
                    0,
                    0,
                    0,
                    0,
                    0
                ],
                "TS": "2018-04-02T16:21:28.815-06:00"
            }
        ]
    }
}
```

## Force Directed Graph reponses

The fdg module's responses are the most complex. There are three sections to the data returned: groups, links, and nodes.

The nodes and links represent the graph. Each node has a name and a group to which it belongs. Links are directional, so they specify a source node and a target node as indexes into the array of nodes, plus a 'value', which represents how "weighty" the link is (in this example, the value is always 1, but it will often be larger).

Groups are defined in the search and are used to color nodes in the fdg display. In this example there are three groups: "operations", "IT", and an unnamed group. Each node belongs to one of the groups, referenced by their index in the groups array.

```
{
    "AdditionalEntries": false,
    "Entries": {
        "groups": [
            "",
            "operations",
            "IT"
        ],
        "links": [
            {
                "source": 0,
                "target": 1,
                "value": 1
            },
            {
                "source": 2,
                "target": 1,
                "value": 1
            },
            {
                "source": 3,
                "target": 1,
                "value": 1
            },
            {
                "source": 2,
                "target": 4,
                "value": 1
            },
            {
                "source": 4,
                "target": 5,
                "value": 1
            },
            {
                "source": 0,
                "target": 6,
                "value": 1
            },
            {
                "source": 2,
                "target": 7,
                "value": 1
            },
            {
                "source": 2,
                "target": 8,
                "value": 1
            },
            {
                "source": 9,
                "target": 8,
                "value": 1
            },
            {
                "source": 10,
                "target": 5,
                "value": 1
            },
            {
                "source": 11,
                "target": 8,
                "value": 1
            },
            {
                "source": 2,
                "target": 12,
                "value": 1
            },
            {
                "source": 2,
                "target": 13,
                "value": 1
            },
            {
                "source": 14,
                "target": 12,
                "value": 1
            },
            {
                "source": 15,
                "target": 12,
                "value": 1
            },
            {
                "source": 16,
                "target": 12,
                "value": 1
            },
            {
                "source": 17,
                "target": 13,
                "value": 1
            },
            {
                "source": 18,
                "target": 13,
                "value": 1
            },
            {
                "source": 13,
                "target": 19,
                "value": 1
            }
        ],
        "nodes": [
            {
                "group": 0,
                "name": "bbd307455de9"
            },
            {
                "group": 1,
                "name": "operations-5 OPERATIONS-5$"
            },
            {
                "group": 0,
                "name": "9b10deadbeef"
            },
            {
                "group": 0,
                "name": "db48a5920a82"
            },
            {
                "group": 2,
                "name": "desktop-2 DESKTOP-2$"
            },
            {
                "group": 2,
                "name": "e758bb7d2630"
            },
            {
                "group": 1,
                "name": "operations-2 OPERATIONS-2$"
            },
            {
                "group": 2,
                "name": "desktop-3 DESKTOP-3$"
            },
            {
                "group": 2,
                "name": "desktop-4 DESKTOP-4$"
            },
            {
                "group": 0,
                "name": "4f194d5cf71a"
            },
            {
                "group": 0,
                "name": "desktop-1 DESKTOP-1$"
            },
            {
                "group": 0,
                "name": "6"
            },
            {
                "group": 1,
                "name": "operation-desktop OPERATION-DESKT$"
            },
            {
                "group": 2,
                "name": "DESKTOP-67T38GD DESKTOP-67T38GD$"
            },
            {
                "group": 0,
                "name": "2fd6276575c7"
            },
            {
                "group": 0,
                "name": "dfc56224743c"
            },
            {
                "group": 0,
                "name": "cb7b71a72272"
            },
            {
                "group": 0,
                "name": "2f01fbc81c46"
            },
            {
                "group": 0,
                "name": "379bd32ecec6"
            },
            {
                "group": 2,
                "name": "foobar"
            }
        ]
    },
    "EntryCount": 19,
    "Finished": true,
    "ID": 18
}
```