# Data Exploration

Gravwell provides several tools to help make sense of data with little or no knowledge of the data itself or of the Gravwell query language. We call this *Data Exploration*.

## Word Filtering

The most basic form of data exploration is word filtering, in which the user clicks a word within the raw entry and either includes it or excludes it from the search. For example, if we run the query `tag=gravwell`, we'll see a bunch of textual results:

![](words1.png)

If we move the mouse cursor over the results, individual words are highlighted. We can click on a highlighted word to bring up a context menu which allows us to *include* or *exclude* the selected word from the search.

![](words2.png)

Note: In the screenshot above, the string `search.go` is also highlighted by an [actionable](#!gui/actionables/actionables.md); the menu item "search.go" contains a sub-menu with the actionable's actions in it.

If we click "Include search.go:189", the query will be automatically modified to `tag=gravwell words search.go:189`. We can then re-run the query to see the new results:

![](words3.png)

We can do this multiple times, including and excluding words, and the query will continue to update:

![](words4.png)

Note that if we have written a query already, we can still click on words and have them properly inserted into the query. For instance, if we manually run the query `tag=gravwell syslog Appname!=webserver` and then click on the word "flow" in the results, the query will be automatically rewritten to `tag=gravwell words flow | syslog Appname!=webserver`.

### Word Filtering Caveats

At this time, word filtering only works on queries which use the [text renderer](#!search/text/text.md). We hope to introduce clickable options for the table renderer in the future.

## Field Extraction

The Query Studio also has the ability to parse many data formats and split out individual fields. For example, a user may run the search `tag=gravwell` just to see what the raw entries look like. Clicking the Details Pane icon, highlighted below, tells Query Studio to attempt to parse the data and split out individual fields:

![](details-icon.png)

If this is the first time the user is looking at this tag, Gravwell must determine the most appropriate format for parsing the data. A window will appear presenting several options:

![](field-extraction.png)

In this case, Gravwell believes the data to be syslog-formatted and shows a preview of the first result parsed as syslog. The user examines this result and, deeming it acceptable, clicks "Select" and saves the selected extraction.

Now the UI shows the Details Pane at the bottom of the window, with individual fields broken out for easier reading. Note the purple bar on the left side of the raw entry; that indicates which entry is currently being displayed in the Details Pane:

![](details-pane.png)

Note how each entry, when pointed to with the mouse cursor, displays the Details Pane icon in a floating button group. Clicking that icon will switch the Details Pane over to show the fields of *that* entry.

Within the extracted fields of the Details Pane, you can click the filter icon and add filters on the given field. In the screenshot below, we are adding a filter to *exclude* any messages where the Appname is "webserver":

![](filter.png)

Selecting the menu option will automatically update the query with a filter on the specified field. Re-run the query to get the updated results:

![](filtered-query.png)

You can then continue in this way, selecting fields from the Details Pane (or clicking on words in the raw entries, see the "Word Filtering" section!) to drill down into precisely the data you care about.

## The Data Explorer UI

The two previous sections describe how to do data exploration within the Query Studio, in which the user types some portions of the query by hand and has other portions automatically inserted based on clicking in the UI. Another option is the Data Explorer UI, a dedicated interface for mouse-driven data exploration.

Selecting the "Data Explorer" item in the main menu brings up a list of tags currently on the system:

![](de1.png)

Clicking on any tag will start exploration on that tag. The Data Explorer attempts to do field extraction, as described in the section above, which means it must determine an appropriate data format for the tag on the first use. In the screenshot below, we see that it has proposed the "syslog" extractor and the user has selected it:

![](de2.png)

Unlike the Query Studio, when working with the Data Explorer it *only* shows the extracted elements, never the raw entry itself. In the screenshot below, we see that the user has added filters on the Message and Appname fields. Note how both filters are shown in the "Filters" pane, and how the operation and value can be changed on the fly.

![](de3.png)

As filters are added, a query string is generated in the "Query" pane. Clicking "Start a Search" will launch that query in the regular search interface; in this way, you can use the Data Explorer to do initial filtering, poking around the data to find the set you're interested in, before pivoting over to the search interface for more detailed interactions.
