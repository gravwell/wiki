# Gravwell 4.0 Big Bang Beta

Greetings and salutations!

You are reading this message because you are in the early access group for Gravwell 4.0 Beta, the Big Bang release! We sincerely thank you for your participation and look forward to feedback and bug reports.

## Current "State of the Beta"

### Desired Testing

Testing desires for this sprint (in order of priority)

* Browsing, installing, and managing kits
* Search template creation and use to create "investigative" dashboards
* Playbooks - New features are coming to enrich playbooks a bit more, but testing of creating and using them as markdown READMEs is the desire for this phase
* Macros
* Templates


We are actively iterating on these features and they may change entirely so maybe don't put a ton of work in building stuff around them.

* Building Kits
 * We intend to make it possible for the community to develop and publish kits but for now we are focused on testing Gravwell published and signed kits.
* Actionables
 * We are playing with how actionables work and including the idea of "active" and "passive" triggers that happen when things are regex-able vs only available when highlighted.
 * We are working on cleaning up how actionables are catagorized to make it clearer what type of data you are pivoting into.
* Search Studio
 * Search Studio is a new experimental feature that allows for rapid testing and developement of queries.
 * You must be in developer mode to see the Search Studio.
* Search templates
 * We are working on an optional "prefix" and "postfix" in order to support optional variables.


### Known Issues

Here are the issues we are aware of, so you don't necessarily need to test or report them.

* Long-running chrome tabs can become unresponsive.
 * We believe we've got an angle on this but be advised that you may need to close the Gravwell UI and open in a new tab or otherwise refresh your page in order to resolve this issue. If you can consistently reproduce this, please jot down your path and pop that over to `beta@gravwell.io`.
* Outdated cached data.
 * Updates to objects like playbooks or installed kits aren't always making it to actual rendering in the UI despite having successfully occurred.  Refreshing the page should update appropriately.
* Search history may be outdated.
 * when multiple sessions exist with the same user, there can be some delays in seeing search history.


## Installation and Upgrade

We're very excited to say this build is now available for your use and testing. We have created a new ubuntu repository and Docker images. Switching from Stable to Beta is done by modifying your apt source respository (or our quick start instructions if installing from scratch).

### Upgrading:
Edit your `/etc/apt/sources.list.d/gravwell.list` file and substitute `https://update.gravwell.io/debian/` for `https://update.gravwell.io/debianbeta/`. Then `apt update` and `apt upgrade` and you should be on the new release.

### Installing from scratch:

```
curl https://update.gravwell.io/debian/update.gravwell.io.gpg.key | sudo apt-key add -
echo 'deb [ arch=amd64 ] https://update.gravwell.io/debianbeta/ community main' | sudo tee /etc/apt/sources.list.d/gravwell.list
sudo apt-get install apt-transport-https
sudo apt-get update
sudo apt-get install gravwell
```

### Docker

The Docker images is available at [gravwell/beta](https://hub.docker.com/r/gravwell/beta). You can substitue `gravwell/gravwell` with `gravwell/beta` in any of the docker documentation and it should "just work."


## Major features included in Gravwell Big Bang

There are many new features in both the GUI and under the hood, but we would like to focus the beta on the features that may need some work.

### Kits

This is it. The Big Bang, the start, a beginning of time. The supermassive featureset that starts the future of Gravwell.

Kits are pre-packaged use case bundles made of searches, dashboards, resources, and more. Read on to see what kind of awesome stuff can be included.

#### Available Kits

Currently, there are two kits available in the Kit repository. We have a Network Enrichment kit, which provides a bunch of great resources for resolving things like port->service and IP geolocation. We also have a Netflow v5 kit, which comes with overview dashboards, investigative dashboards, query library searches, templates, and a playbook to get you started analyzing netflow data.

NOTE: The current netflow kit REQUIRES that netflow data be coming in under the `netflow` tag. In the future, you will have the option of specifying the tag when you install a kit (since tags are arbitrary). For now, make sure that your netflow data is coming in under the `netflow` tag for your beta Gravwell instance.

See the [netflow ingester docs](https://docs.gravwell.io/docs/#!ingesters/ingesters.md#Netflow_Ingester) for more information.


### Query Library

This actually isn't new to the Big Bang release but we feel it deserves another mention due to Kits bundling up saved queries. The query library enables users to save commonly used queries and share with other members of their group. You can now filter based on labels or search using basic text search.

### Playbooks

Playbooks are an opportunity to include READMEs, documentation, directions for activities like threat hunting, context and knowledge about specific internal systems, application log explanations -- YOU NAME IT.

Built using markdown syntax, playbooks can include searches directly for a reader to launch from the playbook. We are actively developing this feature to include more syntactical goodies and are excited about the potential.

### SOAR Scripts

Also not a new feature, but we have expanded the debugging capability and enhanced the UX developing and managing SOAR scripts.

### Actionables

Actionables allow you to set triggers which will make matching text "clickable" when it shows up in the text-based renderers of Gravwell. This can allow you to click an IP address for further investigation, open an AWS security group in the AWS console, or lookup what Shodan knows about a given host. Actionables are also usable when interesting text is highlighted.


#### A note on triggers

The triggers are based on regex parsing of data as it shows up in Gravwell results. Practically speaking, the number of things you can actually regex out of data is fairly small. IP addresses are great, things like AWS instance IDs work ok, but something like a network port does not since it is just an integer. If an actionable trigger matches, that text will be styled as a link and a user who clicks on it will receive the contextual menu of actions available for that item.

Secondarily, a user can highlight/select some text and the contextual menu will appear with available actionables. Thus, they can still be created for things like network ports, but until the user selects the text, they won't attempt to make anything "clickable".

### Macros

Macos now exist for use within search queries. They do what you typically expect a macro to do and enable you to shrink up common activity into a single macro. For example, if I am looking at DNS entries and frequently removing Microsoft domains from my results, I might want to include a macro to make that easier.

I may have a query like `tag=dns ax | grep -v microsoft.com bing.com office365.com live.com windows.net | table`

If I set a $REMOVEMS macro to `grep -v microsoft.com bing.com office365.com live.com windows.net`

My query becomes `tag=dns ax | $REMOVEMS | table`

### Templates

Search templates are similar to the query library but include the ability to define a variable which will be filled in at search time. These are great for conducting investigations or as part of a playbook.

## Thank You

We are very excited about the capabilities that Gravwell Big Bang brings. Thank you for your interest and participation in the beta program. We couldn't do it without you!

Please send us feedback, bug reports, and especially show us cool stuff that you build with the new tools!
