The following is an example default configuration for a Gravwell indexer.

```
[global]
Control-Port=9404
Ingest-Port=4023
TLS-Ingest-Port=4024
Log-Level=INFO
Ingest-Auth=IngestSecrets #This should be changed
Control-Auth=ControlSecrets #This should be changed
Search-Agent-Auth=SearchAgentSecrets #This should be changed
Search-Pipeline-Buffer-Size=1024

Pipe-Ingest-Path=/opt/gravwell/comms/pipe
Log-Location=/opt/gravwell/log
Certificate-File=/opt/gravwell/etc/cert.pem
Key-File=/opt/gravwell/etc/key.pem
Search-Scratch=/opt/gravwell/scratch
License-Location=/opt/gravwell/etc/license

[Default-Well]
        Location=/tmp/storage/default/
        #Tags= should not appear in the Default-Well. Any tag not defined elsewhere will go into the default well

[Storage-Well "syslog"]
        Location=/opt/gravwell/storage/syslog
        Tags=syslog
        Tags=kernel
        Tags=dmesg

[Storage-Well "windows"]
        Location=/opt/gravwell/storage/windows
        Tags=windows
        Tags=winevent

[Storage-Well "weblogs"]
        Location=/opt/gravwell/storage/weblogs
        Tags=apache
        Tags=nginx
        Tags=www

[Storage-Well "raw"]
        Location=/opt/gravwell/storage/raw
        Tags=pcap
        Tags=video
        Tags=audio
```
