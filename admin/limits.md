# Controlling System Resource Usage

This page describes the limits Gravwell places on user's consumption of system resources. These limits are important because incautious users can otherwise make the system unusable for others, by filling up the disk, causing an out-of-memory (OOM) situation, or just by making everyone else's queries extremely slow.

If you've been linked to this page from within the Gravwell UI, one of your queries likely hit one of the limits described here. Besides descriptions of the limit and why it's important, we also offer suggestions for how to avoid hitting these limits.

## Webserver Renderer Storage

Render modules such as chart and table store their results on-disk on the webserver. Depending on the search, the results may consume significant amounts of disk space. For instance, the query `tag=*` tells Gravwell to fetch every entry in the system from the specified timeframe; in a large production environment, this could be terabytes or even petabytes of data! Therefore, the webserver will stop storing data after a certain threshold is reached, to avoid filling up the disk.

If you're seeing messages about render storage limits when running your searches, consider the following options:

* If you're just trying to get an idea of what raw entries look like, consider using the Preview timeframe. Running `tag=foo` across the Preview timeframe tells Gravwell to fetch a small number of the most recent entries, just enough to get an idea of the data.
* If you want to see a few entries from a particular timeframe (and thus can't use Preview), try the limit module: rather than running `tag=foo`, try `tag=foo limit 100`.
* Often, users want to see how *much* data exists in a timeframe, so they run `tag=foo` and look at the overview chart. It's much more efficient to instead run `tag=foo chart` (by default, the `chart` renderer will chart the entry count).
  - Running `tag=foo ax | chart`, if an autoextractor exists for the tag, will also populate the Fields tab of the search results with useful stats about the most common values in each field of the entries.

### Controlling Renderer Storage

To prevent users from filling up the disk with over-large queries, use the `Render-Store-Limit` parameter in the webserver's gravwell.conf file. Setting `Render-Store-Limit=64`, for instance, would set a limit of 64 MB of on-disk storage per query. The default is 1024MB; this gives users a great deal of space to work with while hopefully preventing free space issues on modern disks.

### Downloading Partial Results In Renderer Storage Limited Scenarios

Attempting to download results for a search that has reached the render storage limit will display a warning about downloading partial results. If you're seeing this message, it's important to understand that the downloaded data will only be valid up to the point where the render storage limit was reached, regardless of what the overview chart indicates.

## Gravwell Resource Size

User-created [resources](/resources/resources) can take up a lot of space on disk, on both the webserver and the indexers. In addition, one of the most common uses of resources is to provide lookup tables; in order to use a resource as a lookup table, the *entire* resource must be loaded into memory, meaning that running many simultaneous queries with extremely large lookup tables can put the system at risk of an out-of-memory state. The `Resource-Max-Size` parameter in gravwell.conf specifies a limit, in bytes, for resource size. Thus, setting `Resource-Max-Size=20971520` will limit resources to no more than 20 megabytes.

If you're seeing error messages indicating that your resource exceeds the maximum size, consider the following options:

* Consider if you really need all the columns in your lookup table, especially if you're generating the table from a query. If you only need to match MAC addresses to hostnames in your queries, don't say `tag=inventory ax | unique mac | table -save inv_lookup`, because that will build a table with columns for *every* field on the data; instead, extract only what you need: `tag=inventory ax mac hostname | unique mac | table -save inv_lookup`.
* If you're building resources to do lookups on IP addresses or subnets, check if the specialized formats used by the [ipexist](/search/ipexist/ipexist) or [iplookup](/search/iplookup/iplookup) modules might be more suitable to your needs. These special types of lookup tables can be more space-efficient for some needs.
