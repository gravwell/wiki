# Gravwell Metrics and Crash Reporting

Gravwell users care about what's going on in their networks--it's a big reason people check us out in the first place. We want to make sure all users are aware of and comfortable with the automated crash reporting & metrics systems built into Gravwell. This document will cover both systems, with complete examples of exactly what we send back to the Gravwell servers.

## Crash Reporting

When a Gravwell component crashes, an automated crash report is sent to Gravwell. This consists of the console output from the component in question, which typically includes some brief information about the license (in order to determine whose system just crashed) and a stack trace. **Every** Gravwell component--the webserver, the indexer, the ingesters, the search agent--is set up to send crash reports.
 
Note: Crash reports are always sent via TLS-verified HTTPS to update.gravwell.io. If we are unable to fully validate the remote certificate, the report does *not* go out.

Here's an example of a crash report from a Gravwell employee's test system:

```
Component:      webserver
Version:        3.3.5
Host:           X.X.X.X
Domain: c-X-X-X-X.hsd1.nm.comcast.net
Full Log Location:      /opt/gravwellCustomers/uploads/crash/webserver_3.3.5_X.X.X.X_2020-01-31T14:39:42


Log Snippet:
Version         3.3.10
API Version     0.1
Build Date      2020-Apr-30
Build ID        745dc6ca
Cmdline         /opt/gravwell/bin/gravwell_webserver -stderr gravwell_webserver.service
Executing user  gravwell
Parent PID      1
Parent cmdline  /sbin/init
Parent user     root
Total memory    4147781632
Memory used     5.781707651865122%
Process heap allocation 2005776
Process sys reserved    72112017
CPU count       4
Host hash       4224be94ae35247ed32013d9021f64bc40986c9fbbafac97787ab58b400f1666
Virtualization role     guest
Virtualization system   kvm
max_map_count   65530
RLIMIT_AS (address space)       18446744073709551615 / 18446744073709551615
RLIMIT_DATA (data seg)  18446744073709551615 / 18446744073709551615
RLIMIT_FSIZE (file size)        18446744073709551615 / 18446744073709551615
RLIMIT_NOFILE (file count)      1048576 / 1048576
RLIMIT_STACK (stack size)       8388608 / 18446744073709551615
SKU             2UX
Customer NUM    00000000
Customer GUID   ffffffff-ffff-ffff-ffff-ffffffffffff

panic: send on closed channel

goroutine 90 [running]:
gravwell/pkg/search.(*SearchManager).clientSystemStatsRoutine(0xc01414edc0,
0xc00017fd20, 0x16, 0xc000c300c0, 0xc000c1dc80)
        gravwell@/pkg/search/manager_handlers.go:284 +0x106
created by gravwell/pkg/search.(*SearchManager).GetSystemStats.func1
        gravwell@/pkg/search/manager_handlers.go:301 +0x65
```

The message starts with the particular component that crashed, in this case the webserver. It then lists the Gravwell version, the IP and hostname of the system that crashed, and a path where Gravwell staff can find a full copy of the crash log--if a backtrace is particularly long, we receive an email with only the first 100 lines or so.

The remainder of the message is the console output from the crashed program. The crash reporter pulls this directly from the component's output file in `/dev/shm`; you can see what your system might send by looking at e.g. `/dev/shm/gravwell_webserver`. You can also view any past crash reports in `/opt/gravwell/log/crash`, but be aware that if you *disable* the crash reporter, crash logs will no longer be stored in that directory.

The first few lines (Version, API Version, Build Date, and Build ID) help us determine exactly what version of Gravwell was running. "Cmdline", "Executing user", "Parent PID", "Parent cmdline", and "Parent user" all help us figure out how the Gravwell process is being run and identify potential issues there--in this example, because the parent process as PID 1 and named "manager", we can infer that Gravwell is being run in a Docker container. Sometimes issues can crop up in one environment (e.g. in Docker being launched by "manager") but not in others (Ubuntu, launching via systemd), and this helps us chase that down.

We also include information about the amount of memory on the system and the rlimits set because this can help us track down certain classes of crashes--for instance, an out-of-memory error on a system with 512MB of RAM wouldn't be particularly surprising! Note that the "Host hash" field is a unique identifier for the host running the process, but because it is a hash we can only use it to say "This is the same machine as that other crash report"; no other information is included.

The "SKU", "Customer NUM", and "Customer GUID" fields describe the license in use. The SKU describes the capabilities allowed by the user's license; in this case, the Gravwell employee is using an unlimited ("UX") license. The customer number and customer GUID fields allow us to refer to our customer database and see who is having the problem.

Below all this information is the backtrace from the Gravwell process. In this case, we see that a bug in an alpha build caused a crash in the routine that the webserver uses to check CPU/memory information of the indexers for use on the GUI's hardware stats page. We want to make it very clear that these stacktraces will never contain user data, only line numbers from our source code: just enough so we can figure out where in the program it's crashing.

### Disabling Crash Reporting

If for any reason you decide you don't want to send crash reports, you have multiple options for disabling the report system.

* If using the standalone shell installer, you can disable it at install time with the `--no-crash-report` flag.
* If you installed Gravwell from the Debian repositories, you can disable it with `systemctl disable gravwell_crash_report`.
* If you're using the Gravwell Docker image, you can disable the crash reporter by passing `-e DISABLE_ERROR_HANDLING=true` in the Docker command.

However, we'd really appreciate if you'd leave crash reporting enabled. Thanks to these crash reports, we can often identify and fix problems that users may not even notice! It's one of our best feedback mechanisms to improve our software.

If you would like us to remove all past crash reports that your system has sent, please email support@gravwell.io and we will permanently delete them from our system.

## Metrics Reporting

The Gravwell webserver component (*only* the webserver) will occasionally send an HTTPS POST request to the Gravwell corporate servers with generic usage statistics. This information helps us figure out which features get the most use and which can use more work. We can generate statistics about how much RAM is being consumed by Gravwell--do we need to optimize garbage collection, or be more conservative in our default configuration? It also allows us to make sure paid licenses aren't being deployed improperly.

Our most important goal in gathering these metrics is to protect the anonymity of your data. These metrics reports will **never** include the actual contents of any data stored in Gravwell, nor will they ever send actual search queries or even a list of tags on the system.

Note: Metrics reports are always sent via TLS-verified HTTPS to update.gravwell.io. If we are unable to fully validate the remote certificate, the report does *not* go out.

We use this same system to notify users of new Gravwell releases: when the metrics report is sent, the server will respond with the latest version of Gravwell. This lets us display a notification in the Gravwell UI when a new version is available (these notifications can be disabled with the `Disable-Update-Notification` parameter in gravwell.conf).

Here's an example that was sent by a Gravwell employee's home system:

```
{
    "ApiVer": {
        "Major": 0,
        "Minor": 1
    },
    "AutomatedSearchCount": 1,
    "BuildVer": {
        "BuildDate": "2020-04-02T00:00:00Z",
        "BuildID": "e755ee13",
        "GUIBuildID": "87e5e523",
        "Major": 3,
        "Minor": 3,
        "Point": 8
    },
    "CustomerNumber": 000000000,
    "CustomerUUID": "ffffffff-ffff-ffff-ffff-ffffffffffff",
    "DashboardCount": 5,
    "DashboardLoadCount": 13,
    "DistributedFrontends": false,
    "ForeignDashboardLoadCount": 0,
    "ForeignSearchLoadCount": 0,
    "Groups": 2,
    "IndexerCount": 4,
    "IndexerNodeInfo": [
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 47899944,
            "ProcessSysReserved": 282423040,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 66157568,
            "ProcessSysReserved": 282554112,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 58577296,
            "ProcessSysReserved": 351827712,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        },
        {
            "CPUCount": 12,
            "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
            "ProcessHeapAllocation": 58304584,
            "ProcessSysReserved": 282226432,
            "TotalMemory": 67479150592,
            "VirtRole": "guest",
            "VirtSystem": "docker"
        }
    ],
    "IndexerStats": [
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 658757162,
                    "Entries": 2770447
                },
                {
                    "Cold": false,
                    "Data": 12681258,
                    "Entries": 9882
                },
                {
                    "Cold": false,
                    "Data": 325462303,
                    "Entries": 1344586
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 45312907669,
                    "Entries": 119150365
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 50161444,
                    "Entries": 297743
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 669469662,
                    "Entries": 2931573
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 325986097,
                    "Entries": 1348645
                },
                {
                    "Cold": false,
                    "Data": 50301788,
                    "Entries": 298556
                },
                {
                    "Cold": false,
                    "Data": 45316008062,
                    "Entries": 119174395
                },
                {
                    "Cold": false,
                    "Data": 12341038,
                    "Entries": 9559
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 663669955,
                    "Entries": 2782081
                },
                {
                    "Cold": false,
                    "Data": 326449600,
                    "Entries": 1350525
                },
                {
                    "Cold": false,
                    "Data": 50427080,
                    "Entries": 299538
                },
                {
                    "Cold": false,
                    "Data": 12552734,
                    "Entries": 9759
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 45445347364,
                    "Entries": 119473828
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                }
            ]
        },
        {
            "WellStats": [
                {
                    "Cold": false,
                    "Data": 660249138,
                    "Entries": 2794164
                },
                {
                    "Cold": false,
                    "Data": 45332590720,
                    "Entries": 119204014
                },
                {
                    "Cold": false,
                    "Data": 50572152,
                    "Entries": 300251
                },
                {
                    "Cold": false,
                    "Data": 12608944,
                    "Entries": 9730
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 0,
                    "Entries": 0
                },
                {
                    "Cold": false,
                    "Data": 325899751,
                    "Entries": 1347670
                }
            ]
        }
    ],
    "IngesterCount": 8,
    "LicenseHash": "kH3R+R4AdTCnXFYDi3L4nZ==",
    "LicenseTimeLeft": 23550517079782204,
    "ManualSearchCount": 330,
    "ResourceUpdates": 13356,
    "ResourcesCount": 4,
    "SKU": "2UX",
    "ScheduledSearchCount": 4,
    "SearchCount": 331,
    "Source": "X.X.X.X",
    "SystemMemory": 67479150592,
    "SystemProcs": 3,
    "SystemUptime": 1920449,
    "TimeStamp": "2020-04-02T22:11:23Z",
    "TotalData": 185614443921,
    "TotalEntries": 494907311,
    "Uptime": 300,
    "UserLoginCount": 27,
    "Users": 2,
    "WebserverNodeInfo": {
        "CPUCount": 12,
        "HostHash": "90578d2dcc5bea54614528e1b2c5a25c261cdd7c945f763d2387f309bdd38816",
        "ProcessHeapAllocation": 311618224,
        "ProcessSysReserved": 420052881,
        "TotalMemory": 67479150592,
        "VirtRole": "guest",
        "VirtSystem": "docker"
    },
    "WebserverUUID": "17405830-3ac4-4b75-a639-6a265e6718a4",
    "WellCount": 28
}
```

The structure is large, in part because this webserver is connected to 4 indexers which each get their own set of information. Here's a breakdown of the fields in detail:

* `ApiVer`: An internal Gravwell API versioning number.
* `AutomatedSearchCount`: The number of searches which have been executed "automatically" (by the search agent, or by loading a dashboard).
* `BuildVer`: A structure describing the particular build of Gravwell on this system.
* `CustomerNumber`: The customer number associated with the license on this system.
* `CustomerUUID`: The UUID of the license on this system.
* `DashboardCount`: The number of dashboards that exist.
* `DashboardLoadCount`: The number of types any dashboard has been opened by any user.
* `DistributedFrontends`: Set to true if [distributed webservers](#!distributed/frontend.md) are enabled.
* `ForeignDashboardLoadCount`: The number of times users have viewed dashboards owned by another user (helps us determine if our dashboard sharing options are sufficiently flexible)
* `ForeignSearchLoadCount`: The number of times users have viewed searches owned by another user (helps us determine if our search sharing options are sufficiently flexible)
* `Groups`: The number of user groups on the system.
* `IndexerCount`: The number of indexers to which this webserver is connected.
* `IndexerNodeInfo`: An array of structures, one per indexer, briefly describing the statistics of each indexer:
	- `CPUCount`: The number of CPU cores on the indexer.
	- `HostHash`: A non-reversible hash (see [github.com/denisbrodbeck/machineid](https://github.com/denisbrodbeck/machineid)) that uniquely identifies the host machine running the indexer. Note that in this example, the indexers are all running on a single Docker host, so they all have the same HostHash.
	- `ProcessHeapAllocation`: The amount of heap memory allocated by the indexer process.
	- `ProcessSysReserved`: The total amount of memory the indexer process has reserved from the OS.
	- `TotalMemory`: The size of the system's main memory.
	- `VirtRole`: "host" or "guest", depending on if the indexer is running in a virtual machine or not.
	- `VirtSystem`: The virtualization system, if any, e.g. "xen", "kvm", "vbox", "vmware", "docker", "openvz", "lxc".
* `IndexerStats`: An array of statistics structures for each indexer:
	* `WellStats`: An array of anonymized information about each well on the indexer:
		* `Cold`: Whether or not this is a "cold" well.
		* `Data`: The number of bytes of data in this well.
		* `Entries`: The number of entries in this well.
* `IngesterCount`: The number of unique ingesters attached to the system.
* `LicenseHash`: The MD5 sum of the license in use.
* `LicenseTimeLeft`: The number of seconds remaining in the license.
* `ManualSearchCount`: The number of searches executed manually.
* `ResourceUpdates`: The number of times any resource has been modified.
* `ResourcesCount`: The number of resources on the system.
* `SKU`: The SKU of the license in use.
* `ScheduledSearchCount`: The number of scheduled searches installed on the system.
* `SearchCount`: a deprecated field, the total of `ManualSearchCount` + `AutomatedSearchCount`.
* `Source`: The IP from which this report originated.
* `SystemMemory`: How many bytes of memory are installed on the webserver's host system.
* `SystemProcs`: The number of processes running on the host system.
* `SystemUptime`: Number of seconds the host system has been running.
* `TimeStamp`: The time at which this report was generated.
* `TotalData`: The number of bytes across all wells on all indexers.
* `TotalEntries`: The number of entries across all wells on all indexers.
* `Uptime`: The number of seconds since the webserver process started.
* `UserLoginCount`: The number of times users have logged in.
* `Users`: The number of users on the system.
* `WebserverNodeInfo`: A brief description of the system running the webserver process:
	- `CPUCount`: The number of CPU cores on the webserver.
	- `HostHash`: A non-reversible hash (see [github.com/denisbrodbeck/machineid](https://github.com/denisbrodbeck/machineid)) that uniquely identifies the host machine running the webserver.
	- `ProcessHeapAllocation`: The amount of heap memory allocated by the webserver process.
	- `ProcessSysReserved`: The total amount of memory the webserver process has reserved from the OS.
	- `TotalMemory`: The size of the system's main memory.
	- `VirtRole`: "host" or "guest", depending on if the webserver is running in a virtual machine or not.
	- `VirtSystem`: The virtualization system, if any, e.g. "xen", "kvm", "vbox", "vmware", "docker", "openvz", "lxc".
* `WebserverUUID`: Every Gravwell webserver generates a UUID when installed; this field contains that UUID.
* `WellCount`: The total number of wells across all indexers.

We carefully considered the information we report, taking pains to make it impossible to glean any intelligence regarding the type or content of the data you've got in Gravwell. We are always happy to discuss the reasoning behind any of the information we gather; please email support@gravwell.io with any questions.

### Limiting Metric Reporting

Customers may set `Disable-Stats-Report=true` in gravwell.conf, which will strip down the metrics report to a bare minimum containing the CustomerUUID, CustomerNumber, BuildVer, ApiVer, LicenseTimeLeft, and LicenseHash fields--just enough information to verify that the correct license is installed and the system is running.

We'd really appreciate if you'd leave full stats reports enabled, though. As we said above, these stats reports help us figure out which features are getting used the most, what kind of systems Gravwell is running on, how much RAM it's using--information that, in aggregate, can help us decide where to prioritize future improvements.