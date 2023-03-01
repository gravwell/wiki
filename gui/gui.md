# The Gravwell GUI

Most users will interact with Gravwell through the web GUI. This page describes a few high-level menus and interface concepts; the pages linked below go into detail on other topics.

```{toctree}
---
maxdepth: 1
---
The Search Interface <queries/queries>
Dashboards <dashboards/dashboards>
Persistent Searches <persistent/persistent>
Labels and Filtering <labels/labels>
The Query Library <querylibrary/querylibrary>
Resources </resources/resources>
Auto-Extractors </configuration/autoextractors>
Kits </kits/kits>
Templates <templates/templates>
Macros </search/macros>
User Files <files/files>
Playbooks <playbooks/playbooks>
Actionables <actionables/actionables>
Secrets <secrets/secrets>
Email Configuration</configuration/email>
Advanced GUI User Preferences </configuration/gui>
```

## GUI Introduction

After logging in, you will by default be directed to the search page, shown below.

![](searchpage.png)

The icons along the top (labeled Main Menu, Help, Notifications, and User Profile) are visible at all times within the Gravwell GUI.

## The Main Menu

Clicking the "hamburger" menu in the upper left will open the Main Menu:

![](menu.png)

This menu is used to access all the primary functionalities of Gravwell, including dashboards, the query library, and playbooks. Note that several items within the menu are actually sub-menus, which can be expanded to show additional options:

![](menu-expanded.png)

Items within these sub-menus will typically be used less frequently that the top-level items.

```{note}
These screenshots include an "Administrator" sub-menu, which contains admin-only management tools and is only visible to users flagged as administrators.
```

## Notifications

Clicking the bell-shaped Notifications icon in the upper right of the screen brings up the Notifications display:

![](notifications.png)

Clicking the "snooze" button on a notification will remove that notification from counter shown on the icon; this can be useful to prevent distractions.

Depending on the type of notification, clicking the "delete" icon may clear the notification entirely. Some notifications are persistent and cannot be deleted; some are system-wide and can only be deleted by the administrator, and some are targeted at the current user and can be deleted by that user. Note that there is no harm in clicking "delete" on a notification the user isn't allowed to delete.

Critical notifications, such as an offline indexer, will change the notification bell into a warning sign:

![](notif-warn.png)

## User Preferences

Selecting the User Profile icon in the upper right of the screen brings up a small drop-down menu:

![](user-dropdown.png)

### Account

Selecting "Account" will open your preferences page, shown below. Here, you can change your email address, display name, or password; be sure to click "Update Account" after making changes! The "Log out all sessions" button at the bottom of the screen will kick *all* active sessions for your account, across all client machines.

![](account-prefs.png)

### Interface & Appearance 

The second tab of the Preferences page, "Interface & Appearance", has options for customizing the Gravwell user interface. The "Interface theme" dropdown is of particular interest, as it selects a GUI-wide color scheme (including the ever-popular dark modes). 

The "Chart theme" dropdown selects different color palettes which will be used when drawing charts. The editor theme & font size options control the appearance of Gravwell's built-in text editor, which is used to create automation scripts and in a few other places.

![](interface-prefs.png)

### Preferences

The third tab, "Preferences", allows you to change some default behaviors of Gravwell.

![](general-prefs.png)

The "Home Page" dropdown menu selects which page will be displayed after logging in or clicking the Gravwell icon next to the main menu. By default, the new search page is shown, but you can chose to be shown a list of dashboards, kits, or playbooks instead.

The "Search Group Visibility" option allows you to share the results of all searches with a given group; this can be a convenient way to collaborate. In the screenshot, the user has selected the group named "foo"; all members of that group will have access to the searches this user runs in the future.

The "Advanced Preferences" section can be ignored by most users. Selecting "Developer mode" enables manual editing of JSON preferences (see [this page](/configuration/gui) for more information), while toggling "Experimental Features" will enable the Experimental Features section in the main menu.

### Email Server

The final tab, "Email Server", is extremely important for users who intend to use Gravwell to send automated email alerts.  Complete documentation is available on the [Email Configuration](/configuration/email) page.

## Systems & Health Menu

The Systems & Health sub-menu contains pages which describe the current state of the Gravwell cluster.

![](systems-and-health.png)

### Storage, Indexers, & Wells

This page shows information about the data stored in the indexers of the Gravwell system.

![](storage.png)

The Storage section shows a summary of how much data is in the system, with separate stats for hot and cold storage. The "dropper" graphic indicates how fast new entries are entering the system.

At the bottom of the page, the Search Agent section shows information about the search agent component and when it last "checked in".

The Indexer Status section shows how much data is on each indexer and how quickly each is ingesting new data. If you see that one indexer has much less data than the others, you may need to investigate your ingester configs to make sure they are configured to use *all* indexers. Clicking on an indexer in this section, or clicking on it in the left-hand menu, will open a page which displays more detailed information specific to that indexer:

![](ingester-stats.png)

### Ingesters & Federators

This page shows information about ingesters. The ingester list is searchable and sortable. Ingesters which have connected via Federators will appear in this page, as will the Federators themselves; be aware that entry/byte counts for Federators are the sum of counts from all ingesters connected to them.

![](ingesters-page.png)

### Hardware

The Hardware page shows information about the individual computers which make up the Gravwell cluster. At the top of the page is information about cluster-wide CPU and memory usage, ingest rates, etc.; below are individual "cards" for each indexer (be1, be2, be3):

![](hardware.png)

Each card has several different display options, selected via the links in the upper-right corner of each card. "Health" shows uptime, CPU and memory usage, and network/disk read & write stats. "Ingestion" shows the rate at which new entries are being ingested into that particular indexer. "Specifications" shows system specs for the hardware. "Disks" shows information about the storage on the system, but in general that information is more conveniently viewed on the Disks page.

### Disks

The Disks page contains information about disk storage on the cluster. Only disks which contain Gravwell data will be displayed, to avoid clutter. In the screenshot below, the root disk of indexer `be1` has been expanded to show the wells contained in Gravwell's storage area on that disk.

![](disks.png)


### Topology

The Topology page shows how indexers and ingesters are connected.

![](topology.png)

Note how both indexer1 and indexer2 connect to the same set of wells. This means that the same wells are defined on each indexer. Note also that the "flow" ingester is connected directly to the ingesters, while the others connect via a Federator.
