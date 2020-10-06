# Common Problems & Caveats

There are a few issues which can crop up in the course of configuring and using Gravwell. This document attempts to cover some of the most common.

## Clock Source Warning

When running Gravwell in a virtualized environment such as KVM, you may see a notification like this:

```
Detected potentially slow clock source acpi_pm. Consider changing webservers and indexers to one of the following: tsc, kvm-clock
```

Because Gravwell is a heavily time-oriented system, it needs to check the current time frequently. Some clock sources are significantly slower than others, which can lead to a noticeable slowdown in Gravwell queries.

To modify the clock source, [follow the directions here](https://aws.amazon.com/premiumsupport/knowledge-center/manage-ec2-linux-clock-source/).

Note: If you are unable to modify the clock source, this notification is only visible to the default 'admin' user, not to any other users, and can be ignored if necessary.

## Cannot Reach Gravwell Interface

After installing Gravwell, you may find that your web browser cannot reach the webserver, or that the webserver is not connecting to any indexers. This is frequently a result of closed firewall ports. Gravwell uses a number of TCP ports which must be opened for proper operation. Please refer to [the networking considerations page](networking.md) for more information on which ports must be opened.

## Configuring HTTPS and Secure Listeners

By default, Gravwell does not include or generate TLS certificates. If you intend to use Gravwell on the Internet or any other untrusted network, we strongly recommend you install certificates as soon as possible. See [the certificates page](certificates.md) for instructions.

Note: Gravwell requires certificates that are compatible with TLS 1.2 or later.

## Gravwell Processes Won't Start

There are a few reasons that Gravwell components (webserver, indexers, searchagent, ingesters, etc.) may refuse to start.

### Invalid Configuration

An invalid configuration file will usually lead to the failure of the associated component. There are a few places you can look for more information.

* `/dev/shm/` usually contains the stderr output of the process in question, for example the webserver component will output to `/dev/shm/gravwell_webserver`.
* `/opt/gravwell/log` contains log files from some components.
* Depending on the precise nature of the misconfiguration, ingesters may log errors or warnings to the `gravwell` tag.

### Ownership Issues

The Gravwell components are intended to run as user "gravwell". If the root user runs a Gravwell component manually, it may create or modify essential files and mark them as belonging to root. When run under the "gravwell" account later, the processes will not be able to access the files. You can use `chown` to reassign these files to the gravwell user, but take care *not* to modify anything in `/opt/gravwell/bin` as this can conflict with SELinux flags!

Warning: Do not modify permissions or ownership of files in `/opt/gravwell/bin` unless explicitly instructed by Gravwell support.

### SELinux Issues

Gravwell makes an attempt to properly flag files in `/opt/gravwell` for SELinux compatibility, but careless use of the `chown` and `chmod` commands in `/opt/gravwell/bin` can clear these flags. See [the SELinux section of the hardening document](hardening.md) for more information.

## Gravwell Consumes Too Much Memory/CPU

Because Gravwell often has to deal with huge quantities of data, we do not restrict how much memory or CPU time it can consume. If, however, you must run Gravwell on the same system as some other important software, you may wish to restrict its access to resources. In that case, see the "Systemd Unit Files" section of the [system hardening document](hardening.md).