# GUI configuration and tweaking

This page documents user preference options enabled by the "user preferences" section of the account management page.

Currently, these options are configured via JSON and available for advanced Gravwell users. If you're not sure that making any of these tweaks would be helpful, then they probably wouldn't be.


## User Preferences

User preferences are for advanced users and configurable via the User Preferences JSON string under the account navigation section.

Currently supported options are outlined below.

### Search

By adding a "search" property, users can control aspects to how searches appear and operate.


#### Granularity controls

An option exists to control the granularity of the data renderers. Future versions will probably make this dynamically adjustable with a slider.

```json
{"search":{"granularity":{"chart":199}}}
```

Any renderer is supported (e.g. text, table) but chart is probably the only renderer in which a user would want to control the granularity.

### Force directed graphs

Customizable FDG options are outlined here.

#### Tick animations

Users can disable the animations while the FDGs are iterating over the gravity calculations.

```json
{"fdg":{"tick":false}}
```

### Maps

Customizable Maps options are outlined here.

#### Center and Zoom

Users can set the default zoom and lat/lng for map center:

```json
{"maps":{"center":{"lat":41,"lng":-99,"zoom":2}}}
```
