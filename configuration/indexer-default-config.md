The following is an example default configuration for a Gravwell indexer.

```
[global]
Web-Port=443
Control-Port=9404
Ingest-Port=4023
TLS-Ingest-Port=4024
Log-Level=INFO
Ingest-Auth=IngestSecrets #This should be changed
Control-Auth=ControlSecrets #This should be changed
Search-Agent-Auth=SearchAgentSecrets #This should be changed
Remote-Indexers=net:127.0.0.1:9404
Persist-Web-Logins=True
Session-Timeout-Minutes=1440
Login-Fail-Lock-Count=4
Login-Fail-Lock-Duration=5
Search-Pipeline-Buffer-Size=1024
Web-Port=443

Pipe-Ingest-Path=/opt/gravwell/comms/pipe
Log-Location=/opt/gravwell/log
Web-Log-Location=/opt/gravwell/log/web
Certificate-File=/opt/gravwell/etc/cert.pem
Key-File=/opt/gravwell/etc/key.pem
Render-Store=/opt/gravwell/render
Saved-Store=/opt/gravwell/saved
Search-Scratch=/opt/gravwell/scratch
Web-Files-Path=/opt/gravwell/www
License-Location=/opt/gravwell/etc/license
User-DB-Path=/opt/gravwell/etc/users.db
Web-Store-Path=/opt/gravwell/etc/webstore.db

[Default-Well]
        Location=/opt/gravwell/storage/default/
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
