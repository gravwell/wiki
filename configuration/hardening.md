# Hardening a Gravwell Installation

Hardening a Gravwell installation is a pretty straight forward process.  We take pride in shipping a well contained and well isolated product that adheres to the [principle of least privilege](https://en.wikipedia.org/wiki/Principle_of_least_privilege).

There are a few areas that may warrant some attention upon initial installation, namely TLS certificates.  We ship with a set of defaults that should satisfy a most users but there are a few settings that you may went to tweak.

We also highly reccomend keeping up-to-date with the latest Gravwell releases and occasionally checking in on the [changelog](/docs/#!changelog/list.md).  If we encounter a security issue we will document it there.  We will also notify customers of critical security issues via the prescribed point of contact.

## Quickstart

Securing Gravwell is not much different than securing any other accessible application.  Change the default passwords, setup proper encryption via TLS certificates, and make sure permissiong and authentication tokens are strong.

If you are in a hurry and just want to hit the high points, do this:

1. Install a valid TLS Certificate and enable HTTPS [More Info](/docs/#!configuration/certificates.md)
2. Change the admin users password
3. Change the username for the admin password
4. Ensure you use good secrets for ingesters and enable Ciphertext connections [More Info](/docs/#!ingesters/ingesters.md)

# Gravwell Users and Groups

Gravwell users and groups loosely follow the unix design patterns.  At a high level, Gravwell access controls boil down to the following rules:

1. A user can be a member of multiple groups
2. All searches, resources, and scripts are owned by a single user
3. Searches, resources, scripts, and dashboards can be shared via group membership
4. Access via a group membership does not grant write access (only owners can write)
5. Admin users are not restricted in anyway (think `root`)
6. The Admin user with UID 1 cannot be deleted or locked out of the system (again, think `root`)

## Default Admin User

Default Gravwell installations have a single user named `admin` with the user `changeme`.  The default admin user has the special UID of 1, while you can change the name of the default admin user, you cannot change its UID.  Gravwell treats UID 1 in the same way that Unix treats UID 0.  It is special and cannot be deleted, locked, or otherwise disabled.  You should protect this account.

## Account Lockout

## Persistent Sessions

# Installation Components

The Default installation of Gravwell is installed in `/opt/gravwell`.  The installers create the user and group `gravwell`:`gravwell`.  Neither the user nor the group is installed with login priveleges.  All components execute under the `gravwell` user, and almost all execute under the `gravwell` group.

The notable exception is the `File Follower` ingester which executes under the `admin` group so that it can tail log files in `/var/log`.  If you do NOT want any gravwell component executing with elevated priveleges we reccomend not using the [File Follower](/docs/#!ingesters/ingesters.md#File_Follower) and instead configure syslog to send data to the [Simple Relay](/docs/#!ingesters/ingesters.md#Simple_Relay) ingester via TCP syslog.  You may also alter the File Follower systemd unit file to execute using the `gravwell` group if you do not need to follow any controlled system log files.  Checkout the system unit file section below for more info.

Gravwell installers come in two forms: respository installation packages (either Debian based `.deb` or RedHat based `.rpm`) and shell based self-extracting installers.  The repository installation packages are all signed using the published Gravwell [respository key](https://update.gravwell.io/debian/update.gravwell.io.gpg.key).  The self-extracting shell installers are always accompanied by MD5 hashes, always validate the MD5 hashes and/or repository signatures before installing any package (Gravwell or otherwise).

## Installation Configuration Files

Gravwell configuration files are stored in `/opt/gravwell/etc` and are used to control how webservers, serach agents, indexers, and ingesters behave.  The Gravwell configuration files usually contain shared secret tokens that are used for authentication.  The shared secrets may allow for significant control over a Gravwell component.  For instance, if the `Ingest-Secret` is compromised attackers could send supurfluous entries into the index, consuming storage.

## Systemd Unit Files

Gravwell relies on the [systemd](https://www.freedesktop.org/wiki/Software/systemd/) init manager to get Gravwell up and running, as well as manage crash reports.  The installers register and install [SystemD unit files](https://www.freedesktop.org/software/systemd/man/systemd.unit.html) into `/etc/systemd/system`.  These unit files are responsible for starting the Gravwell processes and applying cgroup restrictions to ensure that Gravwell processes behave.

Most users do not need to change the systemd unit files, but if you need to allow an ingester to touch specific resources or want to run as an alternate user or group, you may want to tweak the `User` or `Group` parameters under the `[Service]` section.

### Systemd Resource Restrictions

Gravwell is built for speed and scale.  Many of our customers are processing hundreds of gigabytes per day on very large systems, we thrive on high core counts, big memory, and fast disks.  You got cores?  We'll use em!  However, Community Edition users or customers that need to coreside Gravwell with other systems may need to ensure that Gravwell doesn't stretch out and start trying to flex all the horse power available to it.  Systemd unit files provide for the ability to control Linux Cgroups, this means that you can restrict how much memory, CPU, and even system file descriptors Gravwell gets access to.

Hardening a system FOR Gravwell may also entail hardening a system FROM Gravwell.  Systemd allows for setting up a bit of a wall so that we only use a subset of total system resources, even when Gravwell is working hard.  Below is an example systemd unit file that restricts the number of system threads, CPU usage, and resident memory.

We are restricting this indexer to 4 threads, 8GB of resident memory, and applying a nice value to the process to reduce its priority.  This means the indexer will play very nice with other processes on the system.

```
[Install]
WantedBy=multi-user.target

[Unit]
Description=Gravwell Indexer Service
After=network-online.target
OnFailure=gravwell_crash_report@%n.service

[Service]
Type=simple
ExecStart=/opt/gravwell/bin/gravwell_indexer -stderr %n
ExecStopPost=/opt/gravwell/bin/gravwell_crash_report -exit-check %n
WorkingDirectory=/opt/gravwell
Restart=always
User=gravwell
Group=gravwell
StandardOutput=null
StandardError=journal
LimitNOFILE=infinity
TimeoutStopSec=120
KillMode=process
KillSignal=SIGINT
LimitNPROC=4
LimitNICE=15
MemoryAccounting=true
MemoryHigh=7168M
MemoryMax=8192M
```

Notice that we don't limit the number of open file descriptors using the `LimitNOFILE` parameter, Gravwell is careful with file descriptors and limiting the number that can be open may cause errors when searching large swaths of time or when many ingesters are connected.  For a full list of all system tweaks available, see the freedesktop.org [exec](https://www.freedesktop.org/software/systemd/man/systemd.exec.html) manual.

## Gravwell Application Executables

All Gravwell service executables (excluding the Windows ingesters) are installed in `/opt/gravwell/bin` and owned by the user `gravwell` and the group `gravwell`.  The permission bits on the executables do not allow for reading, writing, or execution by `other`.  This means that only the root user and/or the `gravwell`:`gravwell` user and group can execute or read the applications.

### Application Capabilities

Some components are installed with special capabilities that allow them to access specific resources without executing as a special user or group.  The following table outlines each component that is installed with additional capabilities and why.

## Gravwell Stored Data

# Webserver to Indexer Communications

# Distributed Webservers

# Ingesters

# Search Agent

# Search Scripting and Automation (SOAR)

# Query Controls
